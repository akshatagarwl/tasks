package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/akshatagarwl/tasks/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TaskHandler struct {
	svc service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) Register(app *fiber.App) {
	app.Post("/task", h.createTask)
	app.Get("/tasks", h.getTasks)
}

type taskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (h *TaskHandler) createTask(c *fiber.Ctx) error {
	var req taskRequest
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

	amTask := AMTask{
		ID:          smTask.ID.String(),
		Title:       smTask.Title,
		Description: smTask.Description,
		Status:      string(smTask.Status),
	}

	return c.Status(http.StatusCreated).JSON(amTask)
}

func (h *TaskHandler) getTasks(c *fiber.Ctx) error {
	idsParam := c.Query("ids")
	statusesParam := c.Query("statuses")

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
	smTasks, err := h.svc.GetTasks(ctx, ids, statuses)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	amTasks := make([]AMTask, 0, len(smTasks))
	for _, t := range smTasks {
		amTasks = append(amTasks, AMTask{
			ID:          t.ID.String(),
			Title:       t.Title,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return c.JSON(amTasks)
}
