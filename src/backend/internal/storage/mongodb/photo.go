package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"course/internal/model"
	"course/internal/storage"
)

type photoDataStorageImpl struct {
	*mongo.Client
}

func NewPhotoDataStorage(db *mongo.Client) storage.PhotoDataStorage {
	return &photoDataStorageImpl{db}
}

func (p photoDataStorageImpl) Save(ctx context.Context, data []byte) (*model.PhotoKey, error) {
	return nil, nil
}

func (p photoDataStorageImpl) Get(ctx context.Context, key *model.PhotoKey) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p photoDataStorageImpl) Update(ctx context.Context, key *model.PhotoKey, data []byte) error {
	//TODO implement me
	panic("implement me")
}

func (p photoDataStorageImpl) Delete(ctx context.Context, key *model.PhotoKey) error {
	//TODO implement me
	panic("implement me")
}
