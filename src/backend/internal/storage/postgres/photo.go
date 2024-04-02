package postgres

import (
	"course/pkg/storage/postgres"
)

type photoStorageImpl struct {
	db *postgres.Postgres
}
