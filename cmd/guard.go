package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var sourcePath string

func init() {
	CreateTable()
	CreateDirTable()
	rootCmd.AddCommand(guardCmd)
	source := guardCmd.Flags()
	source.StringVarP(&sourcePath, "source", "s", "", "File/Directory which should be guarded")
	cobra.MarkFlagRequired(source, "source")
}

var guardCmd = &cobra.Command{
	Use:   "guard",
	Short: "Protect the given file/directory from accidental termination",
	Long:  "This creates the hardlinks to given directory or file.",
	Run: func(cmd *cobra.Command, args []string) {
		createLinks()
	},
}

func createLinks() {
	// Get the source path
	sourcePath, err := filepath.Abs(sourcePath)
	logFatal(err)
	filesList := GetFilesInSource(sourcePath)

	for index, filePath := range filesList {
		hardLink, permissions, uid, gid := createHardLink(index, filePath)
		InsertFileInfo(sourcePath, filePath, hardLink, permissions, uid, gid)
	}
}

// Get List of files to be guarded looping
// internal directories
func GetFilesInSource(path string) []string {
	var filesList []string
	if isDirectory(path) {
		storeDirectoryInfo(path)
		dirFiles := getFilesInPath(path)
		for _, file := range dirFiles {
			fileName := filepath.Join(path, file.Name())
			filesList = append(filesList, GetFilesInSource(fileName)...)
		}
	} else {
		filesList = append(filesList, path)
	}
	return filesList
}

func getFilesInPath(path string) []fs.FileInfo {
	dir, err := os.Open(path)
	logFatal(err)
	dirFiles, err := dir.Readdir(0)
	logFatal(err)
	return dirFiles
}

func isDirectory(path string) bool {
	source, err := os.Open(path)
	logFatal(err)
	defer source.Close()
	sourceFile, err := source.Stat()
	logFatal(err)
	return sourceFile.IsDir()
}

// For each file create hardLinks and return
// source, file, hardlink and permissions
func createHardLink(index int, filePath string) (string, int, int, int) {
	sourcePath, err := filepath.Abs(sourcePath)
	logFatal(err)
	aegisPath, err := filepath.Abs(filepath.Join(filepath.Dir(sourcePath), ".aegis"))
	logFatal(err)
	createIfNotExists(aegisPath)
	hardLinkFileName := fmt.Sprintf("recovery_link_%v_%s", index, filepath.Base(filePath))
	hardLinkPath := filepath.Join(aegisPath, hardLinkFileName)
	err = os.Link(filePath, hardLinkPath)
	logFatal(err)
	permissions, uid, gid := GetFileData(filePath)
	return hardLinkPath, permissions, uid, gid
}

func createIfNotExists(directory string) {
	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(directory, 0644)
		logFatal(err)
	}
}

func GetFileData(path string) (int, int, int) {
	fileInfo, err := os.Stat(path)
	logFatal(err)
	permissions := int(fileInfo.Mode())
	uid := int(fileInfo.Sys().(*syscall.Stat_t).Uid)
	gid := int(fileInfo.Sys().(*syscall.Stat_t).Gid)
	return permissions, uid, gid
}

func storeDirectoryInfo(path string) {
	sourcePath, err := filepath.Abs(sourcePath)
	logFatal(err)
	permissions, uid, gid := GetFileData(path)
	InsertDirectoryInfo(sourcePath, path, permissions, uid, gid)
}
