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

func TestMain(m *testing.M) {
	// Open sql connection
	testDB, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/budgets?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}
	// New Queries object
	testQueries = New(testDB)
	os.Exit(m.Run())
}
