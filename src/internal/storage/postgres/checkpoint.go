package postgres

import (
	"course/pkg/storage/postgres"
)

type checkpointStorageImpl struct {
	db *postgres.Postgres
}
