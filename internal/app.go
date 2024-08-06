package app

import "app/internal/db"

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

	
	newServer := NewServer(":8080", nil)

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
