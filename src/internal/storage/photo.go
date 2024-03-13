package storage

import (
	"course/internal/storage/mongodb"
	postgres "course/internal/storage/postgres"
)

type PhotoStorage interface {
	mongodb.PhotoStorage
	postgres.PhotoStorage
}
