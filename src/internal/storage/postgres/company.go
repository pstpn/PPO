package postgres

import (
	"course/pkg/storage/postgres"
)

type companyStorageImpl struct {
	db *postgres.Postgres
}
