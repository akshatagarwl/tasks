package api

import (
	"time"

	"github.com/akshatagarwl/tasks/service"
)

type AMTaskResponse struct {
	ID             string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title          string    `json:"title" example:"Complete project documentation"`
	Description    *string   `json:"description,omitempty" example:"Write comprehensive API documentation"`
	Status         string    `json:"status" example:"TODO"`
	CreatedAt      time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	LastModifiedAt time.Time `json:"last_modified_at" example:"2023-01-01T00:00:00Z"`
}

type AMUpdateTaskRequest struct {
	Title       *string               `json:"title,omitempty" example:"Updated task title"`
	Description *string               `json:"description,omitempty" example:"Updated task description"`
	Status      *service.SMTaskStatus `json:"status,omitempty" example:"IN_PROGRESS"`
}

type AMCreateTaskRequest struct {
	Title       string  `json:"title" validate:"required" example:"New task title"`
	Description *string `json:"description,omitempty" example:"Task description"`
	Status      *string `json:"status,omitempty" example:"TODO"`
}

type AMPaginationMeta struct {
	Page       int   `json:"page" example:"1"`
	PageSize   int   `json:"page_size" example:"10"`
	TotalCount int64 `json:"total_count" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
}

type AMTasksResponse struct {
	Tasks []AMTaskResponse `json:"tasks"`
	Meta  AMPaginationMeta `json:"meta"`
}

type AMErrorResponse struct {
	Message string `json:"message" example:"Error message"`
}
