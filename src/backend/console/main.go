package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"

	"course/console/handler"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
)

var (
	employeeMenuList = tview.NewList().ShowSecondaryText(false)
	adminMenuList    = tview.NewList().ShowSecondaryText(false)

	pages = tview.NewPages()
	app   = tview.NewApplication()
	form  = tview.NewForm()
	flex  = tview.NewFlex()
	text  = tview.NewTextView().
		SetTextColor(tcell.ColorGreen)
)

func main() {
	l := logger.New(zerolog.ErrorFieldName)

	db, err := postgres.New("postgresql://postgres:admin@localhost:5432/course")
	if err != nil {
		l.Fatal(err)
	}

	h := handler.CreateHandler(l, db)

	//employeeMenuList.
	//	AddItem("Create info card", "", '1', nil).
	//	AddItem("Show info card", "", '2', nil).
	//	AddItem("Change info card data", "", '3', nil)
	adminMenuList.
		AddItem("Create info card", "", '1', nil).
		AddItem("Show info cards", "", '2', nil).
		AddItem("Confirm info card", "", '3', nil)

	pages.AddPage("Menu (guest)", h.CreateGuestMenu(form, pages), true, true).
		AddPage("Register", form, true, true).
		AddPage("Login", form, true, true)
	pages.AddPage("Menu (employee)", h.CreateEmployeeMenu(form, pages), true, true).
		AddPage("Create info card", form, true, true).
		AddPage("Show info card", form, true, true).
		AddPage("Change info card data", form, true, true)

	pages.SwitchToPage("Menu (guest)")

	if err = app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		l.Fatal(err)
	}
}
