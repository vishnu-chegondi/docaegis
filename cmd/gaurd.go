package cmd

import (
	"errors"
	"os"
	"path/filepath"

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
	logFatal(err)

	hardLinkPath := getHardLinkPath(path)
	// fmt.Println(absPath, hardLinkPath)
	err = os.Link(absPath, hardLinkPath)
	logFatal(err)
}

func getHardLinkPath(path string) string {
	pwd := filepath.Dir(path)
	aegisPath := pwd + "/.aegis/"
	if _, err := os.Stat(aegisPath); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(aegisPath, 0644)
		logFatal(err)
	}
	basePath := filepath.Base(path)
	hardLinkPath := aegisPath + basePath
	return hardLinkPath
}
