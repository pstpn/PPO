package postgres

import (
	"course/pkg/storage/postgres"
)

type infoCardStorageImpl struct {
	db *postgres.Postgres
}
