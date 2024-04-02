package postgres

import (
	"course/pkg/storage/postgres"
)

type fieldStorageImpl struct {
	*postgres.Postgres
}
