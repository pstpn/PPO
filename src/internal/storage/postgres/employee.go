package postgres

import (
	"course/pkg/storage/postgres"
)

type employeeStorageImpl struct {
	db *postgres.Postgres
}
