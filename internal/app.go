package app

import (
	"app/internal/controllers" 
	"app/internal/db"
	"app/internal/router"
)

type Deps struct {
	Db       Db
	Services Services
	Server   Server
}

type App struct {
	db       Db
	services Services
	server   Server
}

func NewApp() *App {
	dataService := db.NewDb()

	controllers := controllers.NewControllers(&controllers.Deps{})
	router := router.Routes(controllers)

	newServer := NewServer(":3000", router)

	return &App{
		db:     &dataService,
		server: newServer,
	}
}

func (app *App) Init() error {
	err := app.db.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Start() error {
	err := app.server.Start()
	if err != nil {
		return err
	}

	return nil
}
