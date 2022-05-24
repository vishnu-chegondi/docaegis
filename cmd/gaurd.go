package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	CreateTable()
	rootCmd.AddCommand(gaurdCmd)
}

var gaurdCmd = &cobra.Command{
	Use:   "gaurd",
	Short: "Protect the given file/directory from accidental termination",
	Long:  "This creates the hardlinks to given directory or file.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			createLinks(arg)
		}
	},
}

func createLinks(path string) {
	absPath, err := filepath.Abs(path)
	getFileDetails(absPath)
	logFatal(err)

	hardLinkPath := getHardLinkPath(path)
	// fmt.Println(absPath, hardLinkPath)
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

func getFileDetails(path string) {
	fileInfo, err := os.Stat(path)
	logFatal(err)
	// Name, Mode, Ownership
	fmt.Printf("Group ID: %d\n", int(fileInfo.Sys().(*syscall.Stat_t).Gid))
	fmt.Printf("User ID: %d\n", int(fileInfo.Sys().(*syscall.Stat_t).Uid))
}
