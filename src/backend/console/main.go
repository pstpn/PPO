package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"

	"course/console/handler"
	"course/internal/service"
	company "course/internal/storage/postgres"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
)

var (
	pages = tview.NewPages()
	app   = tview.NewApplication()
	form  = tview.NewForm()
	flex  = tview.NewFlex()
	text  = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(c) to create a new company \n(q) to quit")
)

func main() {
	l := logger.New(zerolog.ErrorFieldName)

	db, err := postgres.New("postgresql://postgres:admin@localhost:5432/course")
	if err != nil {
		l.Fatal(err)
	}

	companyHandler := handler.CreateCompanyHandler(
		l,
		service.NewCompanyService(l, company.NewCompanyStorage(db)),
		flex,
		text,
	)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 'c' {
			form.Clear(true)
			companyHandler.CreateCompanyForm(form, pages)
			pages.SwitchToPage("Create company")
		}

		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Create company", form, true, false)

	if err = app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		l.Fatal(err)
	}
}
