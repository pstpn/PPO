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

type GetPassageRequest struct {
	PassageID int64
}

type ListPassagesRequest struct {
	InfoCardID int64
}

type DeletePassageRequest struct {
	PassageID int64
}