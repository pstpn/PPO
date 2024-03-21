package postgres

import (
	"course/pkg/storage/postgres"
)

type documentStorageImpl struct {
	db *postgres.Postgres
}
