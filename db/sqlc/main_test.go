package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbDSN    = "postgres://postgres:root1234@localhost:5433/bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbDSN)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
