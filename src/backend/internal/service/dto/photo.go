package dto

import "course/internal/model"

type CreatePhotoRequest struct {
	DocumentID string
	Data       []byte
}

type CreatePhotoKeyRequest struct {
	DocumentID *model.DocumentID
	Key        *model.PhotoKey
}

type GetPhotoRequest struct {
	DocumentID string
}

type UpdatePhotoRequest struct {
	DocumentID string
	Data       []byte
}

type UpdatePhotoKeyRequest struct {
	DocumentID *model.DocumentID
	Key        *model.PhotoKey
}

type DeletePhotoRequest struct {
	DocumentID string
}
