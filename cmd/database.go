package cmd

import (
	"database/sql"
	"fmt"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

type DirInfoRow struct {
	SourcePath  string
	Directory   string
	Permissions int
	UID         int
	GID         int
}

func (f *DirInfoRow) GetFileRowArray() []interface{} {
	var columnArray []interface{}
	columnArray = append(columnArray, &f.SourcePath, &f.Directory, &f.Permissions, &f.UID, &f.GID)
	return columnArray
}

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

// Used for creating the initial Directory Table where data is stored regarding
func CreateDirTable() {
	var query string = "CREATE TABLE IF NOT EXISTS dir_info (source_path TEXT NOT NULL, directory TEXT PRIMARY_KEY,permissions INTEGER DEFAULT 644, uid INTEGER DEFAULT 0, gid INTEGER DEFAULT 0)"
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

func InsertDirectoryInfo(values ...interface{}) {
	var query string = "INSERT INTO dir_info (source_path, directory, permissions, uid, gid) VALUES (?, ?, ?, ?, ?)"
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

func GetFileInfo(sourcePath string) []FileInfoRow {
	var listOfFiles []FileInfoRow
	var query string = "SELECT * from file_info where source_path=?"
	tnx := getTnx()
	rows, err := tnx.Query(query, sourcePath)
	logFatal(err)
	for rows.Next() {
		var ansrows FileInfoRow
		columnArray := ansrows.GetFileRowArray()
		err := rows.Scan(columnArray...)
		logFatal(err)
		listOfFiles = append(listOfFiles, ansrows)
	}
	return listOfFiles
}

func GetDirectoryInfo(sourcePath string) []DirInfoRow {
	var listOfDir []DirInfoRow
	var query string = "SELECT * from dir_info where source_path=?"
	tnx := getTnx()
	rows, err := tnx.Query(query, sourcePath)
	logFatal(err)
	for rows.Next() {
		var ansrows DirInfoRow
		columnArray := ansrows.GetFileRowArray()
		err := rows.Scan(columnArray...)
		logFatal(err)
		listOfDir = append(listOfDir, ansrows)
	}
	return listOfDir
}
