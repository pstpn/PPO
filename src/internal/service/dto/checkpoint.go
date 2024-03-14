package dto

import (
	"time"
)

type CreatePassageRequest struct {
	CheckpointID int64
	DocumentID   int64
	Type         int64
	Time         *time.Time
}

type ListPassagesRequest struct {
	CheckpointID int64
}
