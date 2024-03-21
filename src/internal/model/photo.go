package model

type PhotoID int64

func ToPhotoID(id int64) *PhotoID {
	photoID := PhotoID(id)
	return &photoID
}

type PhotoKey string

func ToPhotoKey(key string) *PhotoKey {
	photoKey := PhotoKey(key)
	return &photoKey
}

type PhotoMeta struct {
	PhotoID  *PhotoID
	PhotoKey *PhotoKey
}

type Photo struct {
	Meta *PhotoMeta
	Data []byte
}
