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
	var query string = "CREATE TABLE IF NOT EXISTS file_info (file_path TEXT PRIMARY_KEY,hard_link_path TEXT NOT NULL,permissions INTEGER DEFAULT 644, uid INTEGER DEFAULT 0, gid INTEGER DEFAULT 0)"
	tnx := getTnx()
	_, err := tnx.Exec(query)
	logFatal(err)
	err = tnx.Commit()
	logFatal(err)
}
