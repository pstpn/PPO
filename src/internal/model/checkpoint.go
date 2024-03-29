package model

import "time"

type CheckpointID int64

func ToCheckpointID(id int64) *CheckpointID {
	checkpointID := CheckpointID(id)
	return &checkpointID
}

type PassageType int64

const (
	Entrance PassageType = iota
	Exit
)

func ToPassageType(passage int64) *PassageType {
	passageType := PassageType(passage)
	return &passageType
}

type Passage struct {
	CheckpointID *CheckpointID
	DocumentID   *DocumentID
	Type         *PassageType
	Time         *time.Time
}
