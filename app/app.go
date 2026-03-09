// Package app
package app

import (
	"fmt"
	"manager/internal/config"
	"manager/internal/db"
	"manager/internal/handler"
	"manager/internal/parser"
	"manager/internal/repository"
	"manager/internal/service"
	"manager/pkg/llm"
	"os"

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
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		return nil, fmt.Errorf("DB_CONN_STR environment variable is not set")
	}
	
	database, err := db.NewDB(connStr)
	if err != nil {
		return nil, err
	}
	repositories := repository.NewRepository(database)
	llmClient := llm.NewGeminiClient(cfg.AIConfig.APIKey)
	expenseParser := parser.NewSMSParser(llmClient)
	expenseService := service.NewExpenseService(repositories.Expense(), expenseParser)

	services := service.NewService(expenseService)

	handlers := handler.NewHandler(services)
	handlers.RegisterRoutes(e)

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
