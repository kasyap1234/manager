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
	db       *repository.Repository
	config   *config.Config
	services *service.Service
	handlers *handler.Handler
}

func New() (*App, error) {
	cfg := config.NewConfig()
	e := echo.New()

	db, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	repositories := repository.NewRepository(db)

	services := service.NewService(repositories)

	handlers := handler.NewHandler(services)

	a := &App{
		e:        e,
		db:       db,
		config:   cfg,
		services: services,
		handlers: handlers,
	}
	return a, nil
}
func (a *App) Run() error {
	return a.e.Start(a.config.ServerConfig.Port)
}
