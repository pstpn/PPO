package dto

import (
	"time"
)

type CreatePassageRequest struct {
	CheckpointID int64
	DocumentID   int64
	Type         string
	Time         *time.Time
}

type ListPassagesRequest struct {
	CheckpointID int64
}
