package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	conn    *pgx.Conn
	Queries *Queries
}

func NewTaskRepository(ctx context.Context, host string, port int, user, password, dbname, sslmode string) (*TaskRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	repo := &TaskRepository{
		conn:    conn,
		Queries: New(conn),
	}
	return repo, nil
}

func (r *TaskRepository) Close(ctx context.Context) error {
	if r == nil || r.conn == nil {
		return nil
	}
	return r.conn.Close(ctx)
}
