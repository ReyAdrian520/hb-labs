package db

import (
	"log"
	"os/exec"
	"testing"
	"database/sql"
    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

const(
	driverName = "postgres"
	dataSourceName = "postgresql:///test_ratings?sslmode=disable"
)
var testQueries *Queries

func TestMain(m *testing.M){

	//Finding test_rating dropping if exist
	dropdb := exec.Command("dropdb", "--if-exist", "test_ratings")
	if err := dropdb.Run(); err != nil{
		log.Fatal(err)
	}

	//opening connection in test_ratings
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil{
		log.Fatal(err)
	}


    driver, err := postgres.WithInstance(db, &postgres.Config{})
    mi, err := migrate.NewWithDatabaseInstance(
        "file:///db/migrations",
        "postgres", driver)
    mi.Up()

	testQueries := New(db)
}
