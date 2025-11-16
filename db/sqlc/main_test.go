package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const dbDriver = "postgres"
var dbSource = os.Getenv("DB_SOURCE")

func TestMain(m *testing.M) {
	var err error
	
	if dbSource == "" {
		dbSource = "postgresql://postgres:secret@localhost:5433/runsystem_be_test?sslmode=disable"
	}

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}