package mongodb

import (
	"context"

	"course/internal/model"
)

type PhotoStorage interface {
	Save(ctx context.Context, data []byte) (*model.PhotoKey, error)
	Get(ctx context.Context, key *model.PhotoKey) ([]byte, error)
	Update(ctx context.Context, key *model.PhotoKey, data []byte) error
	Delete(ctx context.Context, key *model.PhotoKey) error
}

type photoStorageImpl struct {
}
