package cmd

import (
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

var restorePath string

func init() {
	CreateTable()
	rootCmd.AddCommand(restoreCmd)
	restore := restoreCmd.Flags()
	restore.StringVarP(&restorePath, "file", "f", "", "File/Directory which should be restored")
	cobra.MarkFlagRequired(restore, "restore")
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore accidentally deleted files",
	Long: `This command is used to restore accidentally deleted
		files from hard links created using gaurd sub command`,
	Run: func(cmd *cobra.Command, args []string) {
		recoverFromLink(restorePath)
	},
}

func recoverFromLink(sourcePath string) {
	var fileInfo FileInfoRow = GetFileInfo(sourcePath)
	byteData, err := os.ReadFile(fileInfo.HardLinkPath)
	logFatal(err)
	err = os.WriteFile(fileInfo.FilePath, byteData, fs.FileMode(fileInfo.Permissions))
	logFatal(err)
	file, err := os.Open(fileInfo.FilePath)
	logFatal(err)
	file.Chown(fileInfo.UID, fileInfo.GID)
}
