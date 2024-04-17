package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"

	"course/config"
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
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)
	l := logger.New(c.Logger.Level, loggerFile)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.Database.Postgres.User,
		c.Database.Postgres.Password,
		c.Database.Postgres.Host,
		c.Database.Postgres.Port,
		c.Database.Postgres.Database,
	))
	if err != nil {
		log.Fatal(err)
	}

	h := handler.CreateHandler(l, db)

	pages.AddPage("Menu (guest)", h.CreateGuestMenu(form, pages, app), true, true).
		AddPage("Register", form, true, true).
		AddPage("Login", form, true, true)
	pages.AddPage("Menu (employee)", h.CreateEmployeeMenu(form, pages), true, true).
		AddPage("Create info card", form, true, true).
		AddPage("Show info card", form, true, true)
	pages.AddPage("Menu (admin)", h.CreateAdminMenu(form, pages, list), true, true).
		AddPage("Show info cards", form, true, true).
		AddPage("Cards", list, true, true).
		AddPage("Validate info card", form, true, true)

	pages.SwitchToPage("Menu (guest)")

	if err = app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		log.Fatal(err)
	}
}
