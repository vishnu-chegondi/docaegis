package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	CreateTable()
	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore accidentally deleted files",
	Long: `This command is used to restore accidentally deleted
		files from hard links created using gaurd sub command`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			recoverFromLink(arg)
		}
	},
}

func recoverFromLink(path string) {
	absPath, err := filepath.Abs(path)
	logFatal(err)
	// TODO: Recover from database the hardlink path, permissions, owner,group.

	hardLinkPath := getHardLinkPath(path)

	byteData, err := os.ReadFile(hardLinkPath)
	logFatal(err)
	err = os.WriteFile(absPath, byteData, 0644)
	logFatal(err)
}
