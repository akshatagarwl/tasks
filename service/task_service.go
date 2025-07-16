package service

import (
	"context"

	"github.com/akshatagarwl/tasks/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService interface {
	GetTasks(ctx context.Context, ids []uuid.UUID, statuses []SMTaskStatus, page, pageSize int) ([]*SMTask, error)
	CreateTask(ctx context.Context, title string, description *string, status *SMTaskStatus) (*SMTask, error)
	UpdateTask(ctx context.Context, id uuid.UUID, title *string, description *string, status *SMTaskStatus) (*SMTask, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type taskService struct {
	repo *db.TaskRepository
}

func NewTaskService(repo *db.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetTasks(ctx context.Context, ids []uuid.UUID, statuses []SMTaskStatus, page, pageSize int) ([]*SMTask, error) {
	var stringStatuses []string
	for _, st := range statuses {
		stringStatuses = append(stringStatuses, string(st))
	}

	offset := (page - 1) * pageSize

	dmTasks, err := s.repo.Queries.GetTasksFiltered(ctx, db.GetTasksFilteredParams{
		Column1: ids,
		Column2: stringStatuses,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
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
			CreatedAt:   d.CreatedAt.Time,
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

	dmTask, err := s.repo.Queries.CreateTask(ctx, params)
	if err != nil {
		return nil, err
	}

	task := &SMTask{
		ID:          dmTask.ID,
		Title:       dmTask.Title,
		Description: &dmTask.Description.String,
		Status:      SMTaskStatus(dmTask.Status),
		CreatedAt:   dmTask.CreatedAt.Time,
	}
	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id uuid.UUID, title *string, description *string, status *SMTaskStatus) (*SMTask, error) {
	params := db.UpdateTaskParams{
		ID: id,
	}
	if title != nil {
		params.Title = pgtype.Text{String: *title, Valid: true}
	} else {
		params.Title = pgtype.Text{Valid: false}
	}

	if description != nil {
		params.Description = pgtype.Text{String: *description, Valid: true}
	} else {
		params.Description = pgtype.Text{Valid: false}
	}

	if status != nil {
		statusStr := string(*status)
		params.Status = &statusStr
	}

	dmTask, err := s.repo.Queries.UpdateTask(ctx, params)
	if err != nil {
		return nil, err
	}

	var descPtr *string
	if dmTask.Description.Valid {
		descPtr = &dmTask.Description.String
	}

	task := &SMTask{
		ID:          dmTask.ID,
		Title:       dmTask.Title,
		Description: descPtr,
		Status:      SMTaskStatus(dmTask.Status),
		CreatedAt:   dmTask.CreatedAt.Time,
	}

	return task, nil
}

func (s *taskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.repo.Queries.DeleteTask(ctx, id)
}
