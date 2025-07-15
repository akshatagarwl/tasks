package main

import (
	"context"
	"log/slog"

	"github.com/akshatagarwl/tasks/api"
	"github.com/akshatagarwl/tasks/config"
	"github.com/akshatagarwl/tasks/db"
	"github.com/akshatagarwl/tasks/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to parse config", "error", err)
		return
	}

	repo, err := db.NewTaskRepository(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return
	}
	defer repo.Close(ctx)

	svc := service.NewTaskService(repo)
	handler := api.NewTaskHandler(svc)

	app := fiber.New()
	handler.Register(app)

	addr := ":" + cfg.ServerPort
	slog.Info("Starting server", "addr", addr)
	if err := app.Listen(addr); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
