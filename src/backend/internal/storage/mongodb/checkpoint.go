package mongodb

import (
	"context"
	"fmt"

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

func NewCheckpointStorage(db *mongodb.MongoDB) storage.CheckpointStorage {
	return &checkpointStorageImpl{db}
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
		ID:           model.ToPassageID(insertedID.InsertedID.(primitive.ObjectID).Hex()),
		CheckpointID: model.ToCheckpointID(request.CheckpointID),
		DocumentID:   model.ToDocumentID(request.DocumentID),
		Type:         model.ToPassageTypeFromInt(request.Type),
		Time:         request.Time,
	}, nil
}

func (c *checkpointStorageImpl) GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error) {
	passage := model.Passage{ID: model.ToPassageID(request.PassageID)}
	objectID, err := primitive.ObjectIDFromHex(request.PassageID)
	if err != nil {
		return nil, err
	}
	err = c.Bucket.GetChunksCollection().FindOne(
		ctx,
		bson.D{{"_id", objectID}},
	).Decode(&passage)
	if err != nil {
		return nil, err
	}

	return &passage, err
}

func (c *checkpointStorageImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	cur, err := c.Bucket.GetChunksCollection().Find(
		ctx,
		bson.D{{"documentID", request.DocumentID}},
	)
	if err != nil {
		return nil, err
	}

	var passages []*model.Passage
	for cur.Next(ctx) {
		var passage model.Passage
		err = cur.Decode(&passage)
		if err != nil {
			return nil, err
		}
		passages = append(passages, &passage)
	}

	return passages, err
}

func (c *checkpointStorageImpl) DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error {
	objectID, err := primitive.ObjectIDFromHex(request.PassageID)
	if err != nil {
		return err
	}
	result, err := c.Bucket.GetChunksCollection().DeleteOne(
		ctx,
		bson.D{{"_id", objectID}},
	)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("incorrect deleted passages count: expected 1, got %d", result.DeletedCount)
	}
	return nil
}

func (c *checkpointStorageImpl) CreateCheckpoint(ctx context.Context, request *dto.CreateCheckpointRequest) (*model.Checkpoint, error) {
	insertedID, err := c.Bucket.GetChunksCollection().InsertOne(ctx, bson.D{
		{"phoneNumber", request.PhoneNumber},
	})
	if err != nil {
		return nil, err
	}

	return &model.Checkpoint{
		ID:          model.ToCheckpointID(insertedID.InsertedID.(primitive.ObjectID).Hex()),
		PhoneNumber: request.PhoneNumber,
	}, nil
}

func (c *checkpointStorageImpl) GetCheckpoint(ctx context.Context, request *dto.GetCheckpointRequest) (*model.Checkpoint, error) {
	checkpoint := model.Checkpoint{ID: model.ToCheckpointID(request.CheckpointID)}
	objectID, err := primitive.ObjectIDFromHex(request.CheckpointID)
	if err != nil {
		return nil, err
	}
	err = c.Bucket.GetChunksCollection().FindOne(
		ctx,
		bson.D{{"_id", objectID}},
	).Decode(&checkpoint)
	if err != nil {
		return nil, err
	}

	return &checkpoint, err
}

func (c *checkpointStorageImpl) DeleteCheckpoint(ctx context.Context, request *dto.DeleteCheckpointRequest) error {
	objectID, err := primitive.ObjectIDFromHex(request.CheckpointID)
	if err != nil {
		return err
	}
	result, err := c.Bucket.GetChunksCollection().DeleteOne(
		ctx,
		bson.D{{"_id", objectID}},
	)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("incorrect deleted checkpoints count: expected 1, got %d", result.DeletedCount)
	}
	return nil
}
