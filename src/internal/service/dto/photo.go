package dto

type CreatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

type GetPhotoRequest struct {
	DocumentID int64
}

type UpdatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}
type DeletePhotoRequest struct {
	DocumentID int64
}
