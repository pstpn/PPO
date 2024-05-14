package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rivo/tview"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
)

var (
	employeeID  int64
	phoneNumber string
	infoCardID  int64
	documentID  int64
	text        = tview.NewTextView()
)

type Handler struct {
	authService     service.AuthService
	employeeService service.EmployeeService
	infoCardService service.InfoCardService
	documentService service.DocumentService
}

func CreateHandler(l logger.Interface, db *postgres.Postgres) *Handler {
	employeeStorage := storage.NewEmployeeStorage(db)
	infoCardStorage := storage.NewInfoCardStorage(db)
	documentStorage := storage.NewDocumentStorage(db)
	return &Handler{
		authService:     service.NewAuthService(l, employeeStorage),
		employeeService: service.NewEmployeeService(l, employeeStorage),
		infoCardService: service.NewInfoCardService(l, infoCardStorage),
		documentService: service.NewDocumentService(l, documentStorage),
	}
}

func (h *Handler) RegisterEmployeeForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	registerEmployeeRequest := &dto.RegisterEmployeeRequest{}

	form.AddInputField("Phone number", "", 20, nil, func(phoneNumber string) {
		registerEmployeeRequest.PhoneNumber = phoneNumber
	})
	form.AddInputField("Full name", "", 20, nil, func(fullName string) {
		registerEmployeeRequest.FullName = fullName
	})
	form.AddInputField("Company ID", "", 20, nil, func(companyID string) {
		id, err := strconv.Atoi(companyID)
		if err != nil {
			id = -1
		}
		registerEmployeeRequest.CompanyID = int64(id)
	})
	form.AddInputField("Post", "", 20, nil, func(post string) {
		registerEmployeeRequest.Post = model.ToPostTypeFromString(post).Int()
	})
	form.AddPasswordField("Password", "", 20, '*', func(password string) {
		registerEmployeeRequest.Password = &model.Password{
			Value:    password,
			IsHashed: false,
		}
	})
	form.AddInputField("Date of birth", "", 20, nil, func(dateOfBirth string) {
		parsedDate, err := time.Parse("02.01.2006", dateOfBirth)
		if err != nil {
			registerEmployeeRequest.DateOfBirth = nil
			return
		}
		registerEmployeeRequest.DateOfBirth = &parsedDate
	})

	form.AddButton("Register", func() {
		employee, err := h.authService.RegisterEmployee(context.TODO(), registerEmployeeRequest)
		if err != nil {
			pages.SwitchToPage("Menu (guest)")
			return
		}
		employeeID = employee.ID.Int()
		phoneNumber = employee.PhoneNumber
		pages.SwitchToPage("Menu (employee)")
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

	return form
}

func (h *Handler) LoginEmployeeForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	loginEmployeeRequest := &dto.LoginEmployeeRequest{}

	form.AddInputField("Phone number", "", 20, nil, func(phoneNumber string) {
		loginEmployeeRequest.PhoneNumber = phoneNumber
	})
	form.AddPasswordField("Password", "", 20, '*', func(password string) {
		loginEmployeeRequest.Password = &model.Password{
			Value:    password,
			IsHashed: false,
		}
	})

	form.AddButton("Login", func() {
		err := h.authService.LoginEmployee(context.TODO(), loginEmployeeRequest)
		if err != nil {
			pages.SwitchToPage("Menu (guest)")
			return
		}
		employee, err := h.employeeService.GetEmployee(context.TODO(), &dto.GetEmployeeRequest{
			PhoneNumber: loginEmployeeRequest.PhoneNumber,
		})
		if err != nil {
			pages.SwitchToPage("Menu (guest)")
			return
		}
		employeeID = employee.ID.Int()
		phoneNumber = employee.PhoneNumber

		if employee.Post.IsAdmin() {
			pages.SwitchToPage("Menu (admin)")
		} else {
			pages.SwitchToPage("Menu (employee)")
		}
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

	return form
}

func (h *Handler) CreateGuestMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Register", "", '1', func() {
			form.Clear(true)
			h.RegisterEmployeeForm(form, pages)
			pages.SwitchToPage("Register")
		}).
		AddItem("Login", "", '2', func() {
			form.Clear(true)
			h.LoginEmployeeForm(form, pages)
			pages.SwitchToPage("Login")
		}).
		AddItem("Exit", "", '3', func() {
			exitFunc.Stop()
		})
}

func (h *Handler) CreateInfoCardForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	now := time.Now()

	createInfoCardRequest := &dto.CreateInfoCardRequest{
		EmployeeID:  employeeID,
		IsConfirmed: false,
		CreatedDate: &now,
	}
	createDocumentRequest := &dto.CreateDocumentRequest{}

	form.AddInputField("Document type", "", 20, nil, func(documentType string) {
		createDocumentRequest.DocumentType = model.ToDocumentTypeFromString(documentType).Int()
	})
	form.AddInputField("Serial number", "", 20, nil, func(serialNumber string) {
		createDocumentRequest.SerialNumber = serialNumber
	})

	form.AddButton("Create", func() {
		infoCard, err := h.infoCardService.CreateInfoCard(context.TODO(), createInfoCardRequest)
		if err != nil {
			pages.SwitchToPage("Menu (employee)")
			return
		}
		createDocumentRequest.InfoCardID = infoCard.ID.Int()
		document, err := h.documentService.CreateDocument(context.TODO(), createDocumentRequest)
		if err != nil {
			pages.SwitchToPage("Menu (employee)")
			return
		}
		infoCardID = infoCard.ID.Int()
		fmt.Println(infoCardID)
		documentID = document.ID.Int()
		pages.SwitchToPage("Menu (employee)")
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (employee)")
	})

	return form
}

func (h *Handler) ShowInfoCard(form *tview.Form, pages *tview.Pages) *tview.Form {
	if infoCardID == 0 {
		form.Clear(true)
		form.AddTextView("", "Create info card before show!\n", 100, 10, true, false)
		form.AddButton("Back", func() {
			pages.SwitchToPage("Menu (employee)")
		})
		return form
	}

	infoCard, err := h.infoCardService.GetInfoCard(context.TODO(), &dto.GetInfoCardByIDRequest{InfoCardID: infoCardID})
	if err != nil {
		form.AddTextView("", err.Error(), 100, 10, true, false)
		return form
	}

	employee, err := h.employeeService.GetEmployee(context.TODO(), &dto.GetEmployeeRequest{PhoneNumber: phoneNumber})
	if err != nil {
		form.AddTextView("", err.Error(), 100, 10, true, false)
		return form
	}

	document, err := h.documentService.GetDocument(context.TODO(), &dto.GetDocumentByIDRequest{DocumentID: documentID})
	if err != nil {
		form.AddTextView("", err.Error(), 100, 10, true, false)
		return form
	}

	form.AddTextView("Employee information:", fmt.Sprintf("Phone number: %s\nFull name: %s\nInfo card ID: %d\nIs confirmed: %v\n"+
		"Post: %s\nDocument type: %s\nDocument serial number: %s\nCreated date: %v\n",
		employee.PhoneNumber,
		employee.FullName,
		infoCard.ID.Int(),
		infoCard.IsConfirmed,
		employee.Post.String(),
		document.Type.String(),
		document.SerialNumber,
		infoCard.CreatedDate,
	), 100, 10, true, false)
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (employee)")
	})

	return form
}

func (h *Handler) CreateEmployeeMenu(form *tview.Form, pages *tview.Pages) *tview.List {
	return tview.NewList().
		AddItem("Create info card", "", '1', func() {
			form.Clear(true)
			h.CreateInfoCardForm(form, pages)
			pages.SwitchToPage("Create info card")
		}).
		AddItem("Show info card", "", '2', func() {
			form.Clear(true)
			h.ShowInfoCard(form, pages)
			pages.SwitchToPage("Show info card")
		}).
		AddItem("Exit", "", '3', func() {
			form.Clear(true)
			pages.SwitchToPage("Menu (guest)")
		})
}

func (h *Handler) ShowInfoCards(form *tview.Form, pages *tview.Pages, list *tview.List) *tview.Form {
	pagination := &postgres.Pagination{
		PageNumber: 0,
	}

	form.AddInputField("Count", "", 20, nil, func(count string) {
		size, err := strconv.Atoi(count)
		if err != nil || size < 0 {
			return
		}
		pagination.PageSize = uint64(size)
	})
	form.AddInputField("Filter column", "", 20, nil, func(column string) {
		pagination.Filter.Column = column
	})
	form.AddInputField("Filter pattern", "", 20, nil, func(pattern string) {
		pagination.Filter.Pattern = pattern
	})
	form.AddInputField("Sort direction (asc/desc; default: asc)", "", 20, nil, func(dir string) {
		sortDir := postgres.ASC
		if dir == "desc" {
			sortDir = postgres.DESC
		}
		pagination.Sort.Direction = sortDir
	})
	form.AddInputField("Sort column", "", 20, nil, func(column string) {
		pagination.Sort.Columns = []string{column}
	})

	form.AddButton("Show", func() {
		infoCards, err := h.infoCardService.ListInfoCards(context.TODO(), &dto.ListInfoCardsRequest{Pagination: pagination})
		if err != nil {
			form.AddTextView("", err.Error(), 100, 10, true, false)
			return
		}
		h.appendToInfoCardList(list, pages, infoCards)
		pages.SwitchToPage("Cards")
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) ValidateInfoCard(form *tview.Form, pages *tview.Pages) *tview.Form {
	validateInfoCardRequest := &dto.ValidateInfoCardRequest{
		IsConfirmed: false,
	}

	form.AddInputField("Info card ID", "", 20, nil, func(infoCardID string) {
		id, err := strconv.Atoi(infoCardID)
		if err != nil {
			form.AddTextView("", err.Error(), 100, 10, true, false)
			return
		}
		validateInfoCardRequest.InfoCardID = int64(id)
	})
	form.AddCheckbox("Is confirmed", false, func(isConfirmed bool) {
		validateInfoCardRequest.IsConfirmed = isConfirmed
	})

	form.AddButton("Validate", func() {
		err := h.infoCardService.ValidateInfoCard(context.TODO(), validateInfoCardRequest)
		if err != nil {
			return
		}
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) CreateAdminMenu(form *tview.Form, pages *tview.Pages, list *tview.List) *tview.List {
	return tview.NewList().
		AddItem("Show info cards", "", '1', func() {
			form.Clear(true)
			h.ShowInfoCards(form, pages, list)
			pages.SwitchToPage("Show info cards")
		}).
		AddItem("Validate info card", "", '2', func() {
			form.Clear(true)
			h.ValidateInfoCard(form, pages)
			pages.SwitchToPage("Validate info card")
		}).
		AddItem("Exit", "", '3', func() {
			form.Clear(true)
			pages.SwitchToPage("Menu (guest)")
		})
}

func (h *Handler) appendToInfoCardList(list *tview.List, pages *tview.Pages, infoCards []*model.InfoCard) {
	list.Clear()
	for _, infoCard := range infoCards {
		list.AddItem(
			strconv.Itoa(int(infoCard.ID.Int())),
			fmt.Sprintf("Is confirmed: %v; Created employee ID: %d; Created date: %v",
				infoCard.IsConfirmed,
				infoCard.CreatedEmployeeID.Int(),
				infoCard.CreatedDate,
			),
			'*',
			nil)
	}
	list.AddItem(
		"Back",
		"",
		'b',
		func() {
			pages.SwitchToPage("Show info cards")
		},
	)
}

func (h *Handler) printInfoCard(infoCard *model.InfoCard) {
	text.Clear()
	text.SetText(fmt.Sprintf("Is confirmed: %v\nCreated employee ID: %d\nCreated date: %v",
		infoCard.IsConfirmed,
		infoCard.CreatedEmployeeID.Int(),
		infoCard.CreatedDate,
	))
}
