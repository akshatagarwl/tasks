package service

import "github.com/google/uuid"

type SMTaskStatus string

const (
	SMTaskStatusPending    SMTaskStatus = "PENDING"
	SMTaskStatusInProgress SMTaskStatus = "IN_PROGRESS"
	SMTaskStatusCompleted  SMTaskStatus = "COMPLETED"
)

type SMTask struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Status      SMTaskStatus
}

func (s SMTaskStatus) IsValid() bool {
	switch s {
	case SMTaskStatusPending, SMTaskStatusInProgress, SMTaskStatusCompleted:
		return true
	default:
		return false
	}
}
