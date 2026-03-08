package app

import (
	"manager/internal/config"
	"manager/internal/db"
	"manager/internal/handler"
	"manager/internal/repository"
	"manager/internal/service"

	"github.com/labstack/echo/v4"
)



type App struct {
	e        *echo.Echo
	repo     *repository.Repository
	config   *config.Config
	services *service.Service
	handlers *handler.Handler
}

func New() (*App, error) {
	cfg := config.NewConfig()
	e := echo.New()
	connStr := ""
	database, err := db.NewDB(connStr)
	if err != nil {
		return nil, err
	}
	repositories := repository.NewRepository(database)

	services := service.NewService(repositories)

	handlers := handler.NewHandler(services)

	a := &App{
		e:        e,
		repo:     repositories,
		config:   cfg,
		services: services,
		handlers: handlers,
	}
	return a, nil
}
func (a *App) Run() error {
	return a.e.Start(a.config.ServerConfig.Port)
}
