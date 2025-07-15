package service

import (
	"context"

	"github.com/akshatagarwl/tasks/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService interface {
	CreateTask(ctx context.Context, title string, description *string, status *SMTaskStatus) (*SMTask, error)
}

type taskService struct {
	repo *db.TaskRepository
}

func NewTaskService(repo *db.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, title string, description *string, status *SMTaskStatus) (*SMTask, error) {
	st := SMTaskStatusPending
	if status != nil {
		st = *status
	}

	var desc pgtype.Text
	if description != nil {
		desc = pgtype.Text{String: *description, Valid: true}
	}

	params := db.CreateTaskParams{
		Title:       title,
		Description: desc,
		Status:      db.DMTaskStatus(st),
	}

	pgID, err := s.repo.Queries.CreateTask(ctx, params)
	if err != nil {
		return nil, err
	}

	id := uuid.UUID(pgID.Bytes)

	task := &SMTask{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      st,
	}
	return task, nil
}
