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
	CheckpointID int64
}

type ListPassagesRequest struct {
	DocumentID string
}

type DeletePassageRequest struct {
	PassageID int64
}

type DeleteCheckpointRequest struct {
	CheckpointID int64
}
