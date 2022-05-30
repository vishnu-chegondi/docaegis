package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var sourcePath string

func init() {
	CreateTable()
	rootCmd.AddCommand(gaurdCmd)
	source := gaurdCmd.Flags()
	source.StringVarP(&sourcePath, "source", "s", "", "File/Directory which should be gaurded")
	cobra.MarkFlagRequired(source, "source")
}

var gaurdCmd = &cobra.Command{
	Use:   "gaurd",
	Short: "Protect the given file/directory from accidental termination",
	Long:  "This creates the hardlinks to given directory or file.",
	Run: func(cmd *cobra.Command, args []string) {
		createLinks(sourcePath)
	},
}

func createLinks(path string) {
	absPath, err := filepath.Abs(path)
	logFatal(err)
	insertFileDetails(absPath)
	hardLinkPath := getHardLinkPath(path)
	err = os.Link(absPath, hardLinkPath)
	logFatal(err)
}

func getHardLinkPath(path string) string {
	pwd := filepath.Dir(path)
	aegisPath, err := filepath.Abs(pwd + "/.aegis/")
	logFatal(err)
	if _, err := os.Stat(aegisPath); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(aegisPath, 0644)
		logFatal(err)
	}
	basePath := filepath.Base(path)
	hardLinkPath := aegisPath + "/" + basePath
	return hardLinkPath
}

func insertFileDetails(path string) {
	fileInfo, err := os.Stat(path)
	logFatal(err)
	file_path := path
	hard_link := getHardLinkPath(path)
	permissions := int(fileInfo.Mode())
	gid := int(fileInfo.Sys().(*syscall.Stat_t).Gid)
	uid := int(fileInfo.Sys().(*syscall.Stat_t).Uid)
	InsertFileInfo(file_path, hard_link, permissions, uid, gid)
}
