package cmd

import (
	"errors"
	"fmt"
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
	pwd, err := os.Getwd()
	logFatal(err)

	absPath, err := filepath.Abs(path)
	basePath := filepath.Base(path)
	logFatal(err)
	aegisPath := pwd + "/.aegis/"
	if _, err := os.Stat(aegisPath); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(aegisPath, 0644)
		logFatal(err)
	}
	fmt.Println(absPath, pwd+"/.aegis/"+basePath)
	// err = os.Link(absPath, pwd+"/.aegis/"+basePath)
	// logFatal(err)
}
