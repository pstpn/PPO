package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_checkpointStorageImpl_CreatePassage(t *testing.T) {
	testDB := NewTestStorage()
	defer testDB.Close()
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	request := &dto.CreatePassageRequest{
		CheckpointID: 1,
		DocumentID:   1,
		Type:         0,
		Time:         &tm,
	}

	passage, err := checkpointStorage.CreatePassage(context.TODO(), request)

	require.NoError(t, err)
	require.Equal(t, model.ToCheckpointID(request.CheckpointID), passage.CheckpointID)
	require.Equal(t, model.ToDocumentID(request.DocumentID), passage.DocumentID)
	require.Equal(t, model.ToPassageTypeFromInt(request.Type), passage.Type)
	require.Equal(t, request.Time, passage.Time)

	err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
	require.NoError(t, err)
}
