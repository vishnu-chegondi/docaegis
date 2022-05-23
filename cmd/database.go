package cmd

import (
	"database/sql"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func GetDbLocation() string {
	var dbLocation string
	switch runtime.GOOS {
	case "darwin":
		dbLocation = "/var/lib/docaegis.db"
	default:
		dbLocation = "/var/lib/docaegis.db"
	}
	return dbLocation
}

func GetDb() *sql.DB {
	dbLocation := GetDbLocation()
	db, err := sql.Open("sqlite3", dbLocation)
	logFatal(err)
	return db
}

func getTnx() *sql.Tx {
	db := GetDb()
	tnx, err := db.Begin()
	logFatal(err)
	return tnx
}

// Used for creating the initial Table where data is stored regarding
func CreateTable() {
	var query string = "CREATE TABLE IF NOT EXISTS test_table (column_1 TEXT PRIMARY_KEY,column_2 TEXT NOT NULL,column_3 data_type DEFAULT 0,table_constraints)"

	tnx := getTnx()
	_, err := tnx.Exec(query)
	logFatal(err)
	err = tnx.Commit()
	logFatal(err)
}
