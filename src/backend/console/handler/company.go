package handler

import (
	"context"

	"github.com/rivo/tview"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

var (
	companiesList = tview.NewList().ShowSecondaryText(false)
	companyText   = tview.NewTextView()
)

type CompanyHandler struct {
	logger  logger.Interface
	service service.CompanyService
	data    []*model.Company
}

func CreateCompanyHandler(l logger.Interface, companyService service.CompanyService, flex *tview.Flex, text *tview.TextView) *CompanyHandler {
	companyHandler := &CompanyHandler{
		logger:  l,
		service: companyService,
		data:    make([]*model.Company, 0),
	}

	companiesList.SetSelectedFunc(func(index int, name string, city string, shortcut rune) {
		companyHandler.print(companyHandler.data[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(companiesList, 0, 1, true).
			AddItem(companyText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	return companyHandler
}

func (c *CompanyHandler) CreateCompanyForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	companyCreateRequest := &dto.CreateCompanyRequest{}

	form.AddInputField("Name", "", 20, nil, func(name string) {
		companyCreateRequest.Name = name
	})
	form.AddInputField("City", "", 20, nil, func(city string) {
		companyCreateRequest.City = city
	})

	form.AddButton("Create", func() {
		company, err := c.service.Create(context.TODO(), companyCreateRequest)
		if err != nil {
			c.logger.Fatal(err)
		}
		c.data = append(c.data, company)
		c.appendToList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func (c *CompanyHandler) appendToList() {
	companiesList.Clear()
	for index, company := range c.data {
		companiesList.AddItem(company.Name+", "+company.City, " ", rune('1'+index), nil)
	}
}

func (c *CompanyHandler) print(company *model.Company) {
	companyText.Clear()
	text := company.Name + ", " + company.City
	companyText.SetText(text)
}
