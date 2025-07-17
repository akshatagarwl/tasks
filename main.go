package main

import (
	"context"
	"log/slog"

	"github.com/akshatagarwl/tasks/api"
	"github.com/akshatagarwl/tasks/config"
	"github.com/akshatagarwl/tasks/db"
	"github.com/akshatagarwl/tasks/docs"
	"github.com/akshatagarwl/tasks/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/akshatagarwl/tasks/docs"
)

// @title Task Management API
// @version 1.0
// @description A simple task management API with CRUD operations
// @contact.name API Support
// @contact.email hey@akshat.dev
// @BasePath /
func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to parse config", "error", err)
		return
	}

	repo, err := db.NewTaskRepository(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return
	}
	defer repo.Close(ctx)

	svc := service.NewTaskService(repo)
	handler := api.NewTaskHandler(svc)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://tasks-api.akshat.dev, http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	docs.SwaggerInfo.Host = cfg.Host
	app.Get("/swagger/*", swagger.HandlerDefault)

	handler.Register(app)

	addr := ":" + cfg.ServerPort
	slog.Info("Starting server", "addr", addr)
	if err := app.Listen(addr); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
