package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/akshatagarwl/tasks/db"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
}

func main() {
	ctx := context.Background()

	cfg, err := env.ParseAs[config]()
	if err != nil {
		slog.Error("failed to parse config", "error", err)
		return
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	_, err = queries.CreateTask(ctx, db.CreateTaskParams{
		Title:       "implement service layer",
		Description: pgtype.Text{String: "Once the repository layer has been implemented. Implement the Service Layer.", Valid: true},
		Status:      "PENDING",
	})
	if err != nil {
		slog.Error("failed to create task", "error", err)
		return
	}
}
