package model

import "time"

type CheckpointID int64

func ToCheckpointID(id int64) *CheckpointID {
	checkpointID := CheckpointID(id)
	return &checkpointID
}

type PassageID int64

func ToPassageID(id int64) *PassageID {
	passageID := PassageID(id)
	return &passageID
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

func (p *PassageType) String() string {
	switch *p {
	case Entrance:
		return "Вход"
	case Exit:
		return "Выход"
	default:
		return "Неизвестно"
	}
}

type Passage struct {
	ID           *PassageID
	CheckpointID *CheckpointID
	DocumentID   *DocumentID
	Type         *PassageType
	Time         *time.Time
}
