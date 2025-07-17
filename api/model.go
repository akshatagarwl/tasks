package api

import (
	"time"

	"github.com/akshatagarwl/tasks/service"
)

type AMTaskResponse struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    *string   `json:"description,omitempty"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}

type AMUpdateTaskRequest struct {
	Title       *string               `json:"title,omitempty"`
	Description *string               `json:"description,omitempty"`
	Status      *service.SMTaskStatus `json:"status,omitempty"`
}

type AMCreateTaskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}


type AMPaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}

type AMTasksResponse struct {
	Tasks []AMTaskResponse `json:"tasks"`
	Meta  AMPaginationMeta `json:"meta"`
}
