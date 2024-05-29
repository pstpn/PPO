package dto

import (
	"time"
)

type CreateCheckpointRequest struct {
	PhoneNumber string
}

type CreatePassageRequest struct {
	CheckpointID string
	DocumentID   string
	Type         int64
	Time         *time.Time
}

type GetPassageRequest struct {
	PassageID string
}

type GetCheckpointRequest struct {
	CheckpointID string
}

type ListPassagesRequest struct {
	DocumentID string
}

type DeletePassageRequest struct {
	PassageID string
}

type DeleteCheckpointRequest struct {
	CheckpointID string
}
