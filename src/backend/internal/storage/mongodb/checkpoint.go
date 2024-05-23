package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/mongodb"
)

type checkpointStorageImpl struct {
	*mongodb.MongoDB
}

func (c *checkpointStorageImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error) {
	insertedID, err := c.Bucket.GetChunksCollection().InsertOne(ctx, bson.D{
		{"checkpointID", request.CheckpointID},
		{"documentID", request.DocumentID},
		{"type", request.Type},
		{"time", request.Time},
	})
	if err != nil {
		return nil, err
	}

	return &model.Passage{
		ID:           model.ToPassageID(insertedID.InsertedID.(primitive.ObjectID).String()),
		CheckpointID: model.ToCheckpointID(request.CheckpointID),
		DocumentID:   model.ToDocumentID(request.DocumentID),
		Type:         model.ToPassageTypeFromInt(request.Type),
		Time:         request.Time,
	}, nil
}

func (c *checkpointStorageImpl) GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error) {
	//TODO implement me
	panic("implement me")
}

func (c *checkpointStorageImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	//TODO implement me
	panic("implement me")
}

func (c *checkpointStorageImpl) DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c *checkpointStorageImpl) CreateCheckpoint(ctx context.Context, request *dto.CreateCheckpointRequest) (*model.Checkpoint, error) {
	//TODO implement me
	panic("implement me")
}

func (c *checkpointStorageImpl) GetCheckpoint(ctx context.Context, request *dto.GetCheckpointRequest) (*model.Checkpoint, error) {
	//TODO implement me
	panic("implement me")
}

func (c *checkpointStorageImpl) DeleteCheckpoint(ctx context.Context, request *dto.DeleteCheckpointRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewCheckpointStorage(db *mongodb.MongoDB) storage.CheckpointStorage {
	return &checkpointStorageImpl{db}
}
