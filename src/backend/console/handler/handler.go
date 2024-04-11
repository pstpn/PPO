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
	logger          logger.Interface
	authService     service.AuthService
	employeeService service.EmployeeService
	infoCardService service.InfoCardService
	documentService service.DocumentService

	employees []*model.Employee
}

func CreateHandler(l logger.Interface, db *postgres.Postgres) *Handler {
	employeeStorage := storage.NewEmployeeStorage(db)
	infoCardStorage := storage.NewInfoCardStorage(db)
	documentStorage := storage.NewDocumentStorage(db)
	return &Handler{
		logger:          l,
		employees:       make([]*model.Employee, 0),
		authService:     service.NewAuthService(l, employeeStorage),
		employeeService: service.NewEmployeeService(l, employeeStorage),
		infoCardService: service.NewInfoCardService(l, infoCardStorage),
		documentService: service.NewDocumentService(l, documentStorage),
	}
}

func (h *Handler) RegisterEmployeeForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

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

	return form
}

func (h *Handler) LoginEmployeeForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

	loginEmployeeRequest := &dto.LoginEmployeeRequest{}

	form.AddInputField("Phone number", "", 20, nil, func(phoneNumber string) {
		loginEmployeeRequest.PhoneNumber = phoneNumber
	})
	form.AddPasswordField("Password", "", 20, '*', func(password string) {
		loginEmployeeRequest.Password = password
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

	return form
}

func (h *Handler) CreateGuestMenu(form *tview.Form, pages *tview.Pages) *tview.List {
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
		})
}

func (h *Handler) CreateInfoCardForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (employee)")
	})

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

	return form
}

func (h *Handler) ShowInfoCard(form *tview.Form, pages *tview.Pages) *tview.Form {
	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (employee)")
	})

	if infoCardID == 0 {
		form.AddTextView("", "Create info card before show!\n", 100, 10, true, false)
		return form
	}

	infoCard, err := h.infoCardService.GetInfoCard(context.TODO(), &dto.GetInfoCardRequest{InfoCardID: infoCardID})
	if err != nil {
		form.AddTextView("", err.Error(), 100, 10, true, false)
		return form
	}

	employee, err := h.employeeService.GetEmployee(context.TODO(), &dto.GetEmployeeRequest{PhoneNumber: phoneNumber})
	if err != nil {
		form.AddTextView("", err.Error(), 100, 10, true, false)
		return form
	}

	document, err := h.documentService.GetDocument(context.TODO(), &dto.GetDocumentRequest{DocumentID: documentID})
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
		})
}

func (h *Handler) CreateAdminMenu(form *tview.Form, pages *tview.Pages) *tview.List {
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
		})
}

//func (h *Handler) appendToEmployeeList() {
//	employeesList.Clear()
//	for index, employee := range h.employees {
//		employeesList.AddItem(employee.FullName, " ", rune('1'+index), nil)
//	}
//}

//func (h *Handler) printEmployee(employee *model.Employee) {
//	text.Clear()
//	text.SetText(employee.PhoneNumber)
//}
