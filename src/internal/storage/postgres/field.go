package postgres

import (
	"course/pkg/storage/postgres"
)

type fieldStorageImpl struct {
	db *postgres.Postgres
}
