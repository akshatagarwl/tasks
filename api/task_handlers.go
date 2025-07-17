package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/akshatagarwl/tasks/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/google/uuid"
)

type TaskHandler struct {
	svc service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) Register(app *fiber.App) {
	app.Use(healthcheck.New())
	app.Get("/task", h.getTasks)
	app.Post("/task", h.createTask)
	app.Put("/task/:id", h.updateTask)
	app.Delete("/task/:id", h.deleteTask)
}

func (h *TaskHandler) createTask(c *fiber.Ctx) error {
	var req AMCreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	if req.Title == "" {
		return fiber.NewError(http.StatusBadRequest, "title is required")
	}

	var statusPtr *service.SMTaskStatus
	if req.Status != nil {
		st := service.SMTaskStatus(*req.Status)
		if !st.IsValid() {
			return fiber.NewError(http.StatusBadRequest, "invalid status value")
		}
		statusPtr = &st
	}

	ctx := context.Background()
	smTask, err := h.svc.CreateTask(ctx, req.Title, req.Description, statusPtr)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	amTask := AMTaskResponse{
		ID:          smTask.ID.String(),
		Title:       smTask.Title,
		Description: smTask.Description,
		Status:      string(smTask.Status),
		CreatedAt:   smTask.CreatedAt,
	}

	return c.Status(http.StatusCreated).JSON(amTask)
}

func (h *TaskHandler) updateTask(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid UUID format")
	}

	var req AMUpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid request body")
	}

	if req.Title == nil && req.Description == nil && req.Status == nil {
		return fiber.NewError(http.StatusBadRequest, "At least one field must be provided for update")
	}

	task, err := h.svc.UpdateTask(c.Context(), id, req.Title, req.Description, req.Status)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(&AMTaskResponse{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt,
	})
}

func (h *TaskHandler) deleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid UUID format")
	}

	if err := h.svc.DeleteTask(c.Context(), id); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *TaskHandler) getTasks(c *fiber.Ctx) error {
	idsParam := c.Query("ids")
	statusesParam := c.Query("statuses")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)

	var ids []uuid.UUID
	if idsParam != "" {
		for _, s := range strings.Split(idsParam, ",") {
			id, err := uuid.Parse(strings.TrimSpace(s))
			if err != nil {
				return fiber.NewError(http.StatusBadRequest, "invalid uuid in ids")
			}
			ids = append(ids, id)
		}
	}

	var statuses []service.SMTaskStatus
	if statusesParam != "" {
		for _, stStr := range strings.Split(statusesParam, ",") {
			st := service.SMTaskStatus(strings.TrimSpace(stStr))
			if !st.IsValid() {
				return fiber.NewError(http.StatusBadRequest, "invalid status value")
			}
			statuses = append(statuses, st)
		}
	}

	ctx := context.Background()
	result, err := h.svc.GetTasksWithCount(ctx, ids, statuses, page, pageSize)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	amTasks := make([]AMTaskResponse, 0, len(result.Tasks))
	for _, t := range result.Tasks {
		amTasks = append(amTasks, AMTaskResponse{
			ID:          t.ID.String(),
			Title:       t.Title,
			Description: t.Description,
			Status:      string(t.Status),
			CreatedAt:   t.CreatedAt,
		})
	}

	totalPages := int((result.TotalCount + int64(pageSize) - 1) / int64(pageSize))

	response := AMTasksResponse{
		Tasks: amTasks,
		Meta: AMPaginationMeta{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: result.TotalCount,
			TotalPages: totalPages,
		},
	}

	return c.JSON(response)
}
