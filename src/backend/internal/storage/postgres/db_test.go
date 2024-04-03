package postgres

import (
	"course/pkg/storage/postgres"
)

const connURL = "postgresql://postgres:admin@localhost:5432/tests"

func NewTestStorage() *postgres.Postgres {
	conn, err := postgres.New(connURL)
	if err != nil {
		panic(err)
	}
	return conn
}
