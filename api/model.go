package api

import "github.com/akshatagarwl/tasks/service"

type AMTaskResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`
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
