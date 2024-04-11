package main

import (
	"github.com/rivo/tview"
	"github.com/rs/zerolog"

	"course/console/handler"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
)

var (
	pages = tview.NewPages()
	app   = tview.NewApplication()
	form  = tview.NewForm()
	list  = tview.NewList().ShowSecondaryText(true)
)

func main() {
	l := logger.New(zerolog.ErrorFieldName)

	db, err := postgres.New("postgresql://postgres:admin@localhost:5432/course")
	if err != nil {
		l.Fatal(err)
	}

	h := handler.CreateHandler(l, db)

	pages.AddPage("Menu (guest)", h.CreateGuestMenu(form, pages), true, true).
		AddPage("Register", form, true, true).
		AddPage("Login", form, true, true)
	pages.AddPage("Menu (employee)", h.CreateEmployeeMenu(form, pages), true, true).
		AddPage("Create info card", form, true, true).
		AddPage("Show info card", form, true, true)
	pages.AddPage("Menu (admin)", h.CreateAdminMenu(form, pages, list), true, true).
		AddPage("Show info cards", form, true, true).
		AddPage("Cards", list, true, true).
		AddPage("Confirm info card", form, true, true)

	pages.SwitchToPage("Menu (guest)")

	if err = app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		l.Fatal(err)
	}
}
