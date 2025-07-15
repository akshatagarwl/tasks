package service

import (
	"context"

	"github.com/akshatagarwl/tasks/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService interface {
	GetTasks(ctx context.Context, ids []uuid.UUID, statuses []SMTaskStatus) ([]*SMTask, error)
	CreateTask(ctx context.Context, title string, description *string, status *SMTaskStatus) (*SMTask, error)
}

type taskService struct {
	repo *db.TaskRepository
}

func NewTaskService(repo *db.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetTasks(ctx context.Context, ids []uuid.UUID, statuses []SMTaskStatus) ([]*SMTask, error) {
	var stringStatuses []string
	for _, st := range statuses {
		stringStatuses = append(stringStatuses, string(st))
	}

	dmTasks, err := s.repo.Queries.GetTasksFiltered(ctx, db.GetTasksFilteredParams{
		Column1: ids,
		Column2: stringStatuses,
	})
	if err != nil {
		return nil, err
	}

	tasks := make([]*SMTask, 0, len(dmTasks))
	for _, d := range dmTasks {
		var descPtr *string
		if d.Description.Valid {
			descPtr = &d.Description.String
		}
		t := &SMTask{
			ID:          d.ID,
			Title:       d.Title,
			Description: descPtr,
			Status:      SMTaskStatus(d.Status),
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
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
		Status:      string(st),
	}

	id, err := s.repo.Queries.CreateTask(ctx, params)
	if err != nil {
		return nil, err
	}

	task := &SMTask{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      st,
	}
	return task, nil
}
