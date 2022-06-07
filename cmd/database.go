package cmd

import (
	"database/sql"
	"fmt"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

type FileInfoRow struct {
	SourcePath   string
	FilePath     string
	HardLinkPath string
	Permissions  int
	UID          int
	GID          int
}

func (f *FileInfoRow) GetFileRowArray() []interface{} {
	var columnArray []interface{}
	columnArray = append(columnArray, &f.SourcePath, &f.FilePath, &f.HardLinkPath, &f.Permissions, &f.UID, &f.GID)
	return columnArray
}

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
	var query string = "CREATE TABLE IF NOT EXISTS file_info (source_path TEXT NOT NULL, file_path TEXT PRIMARY_KEY,hard_link_path TEXT NOT NULL,permissions INTEGER DEFAULT 644, uid INTEGER DEFAULT 0, gid INTEGER DEFAULT 0)"
	tnx := getTnx()
	_, err := tnx.Exec(query)
	logFatal(err)
	err = tnx.Commit()
	logFatal(err)
}

func InsertFileInfo(values ...interface{}) {
	var query string = "INSERT INTO file_info (source_path, file_path, hard_link_path, permissions, uid, gid) VALUES (?,?,?,?,?,?)"
	tnx := getTnx()
	_, err := tnx.Exec(query, values...)
	logFatal(err)
	err = tnx.Commit()
	logFatal(err)
}

func GetAllFilesGaurded() {
	var query string = "SELECT distinct(source_path) from file_info"
	tnx := getTnx()
	rows, err := tnx.Query(query)
	logFatal(err)
	for rows.Next() {
		var filePath string
		err = rows.Scan(&filePath)
		logFatal(err)
		fmt.Println(filePath)
	}
}

func GetFileInfo(sourcePath string) FileInfoRow {
	var query string = "SELECT * from file_info where source_path=?"
	tnx := getTnx()
	row := tnx.QueryRow(query, sourcePath)
	var ansrows FileInfoRow
	columnArray := ansrows.GetFileRowArray()
	err := row.Scan(columnArray...)
	logFatal(err)
	return ansrows
}
