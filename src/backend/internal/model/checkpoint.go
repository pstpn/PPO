package model

import "time"

type CheckpointID string

func ToCheckpointID(id string) *CheckpointID {
	checkpointID := CheckpointID(id)
	return &checkpointID
}

type Checkpoint struct {
	ID          *CheckpointID
	PhoneNumber string
}

type PassageID string

func ToPassageID(id string) *PassageID {
	passageID := PassageID(id)
	return &passageID
}

type PassageType int64

const (
	Entrance PassageType = iota
	Exit
	UnknownPassageType
)

func ToPassageTypeFromInt(passage int64) *PassageType {
	passageType := PassageType(passage)
	return &passageType
}

func ToPassageTypeFromString(passage string) *PassageType {
	var passageType PassageType
	switch passage {
	case "Вход":
		passageType = Entrance
	case "Выход":
		passageType = Exit
	default:
		passageType = UnknownPassageType
	}

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
