package main

import (
	"context"
	"log/slog"

	"github.com/akshatagarwl/tasks/api"
	"github.com/akshatagarwl/tasks/db"
	"github.com/akshatagarwl/tasks/service"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2"
)

type config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
}

func main() {
	ctx := context.Background()

	cfg, err := env.ParseAs[config]()
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
