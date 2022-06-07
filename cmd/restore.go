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
		recoverDirectories(restorePath)
		recoverFromLink(restorePath)
	},
}

func recoverDirectories(sourcePath string) {
	var dirsInfo []DirInfoRow = GetDirectoryInfo(sourcePath)
	for _, dirInfo := range dirsInfo {
		err := os.MkdirAll(dirInfo.Directory, fs.FileMode(dirInfo.Permissions))
		logFatal(err)
		err = os.Chown(dirInfo.Directory, dirInfo.UID, dirInfo.GID)
		logFatal(err)
	}
}

func recoverFromLink(sourcePath string) {
	var filesInfo []FileInfoRow = GetFileInfo(sourcePath)
	for _, fileInfo := range filesInfo {
		byteData, err := os.ReadFile(fileInfo.HardLinkPath)
		logFatal(err)
		err = os.WriteFile(fileInfo.FilePath, byteData, fs.FileMode(fileInfo.Permissions))
		logFatal(err)
		file, err := os.Open(fileInfo.FilePath)
		logFatal(err)
		file.Chown(fileInfo.UID, fileInfo.GID)
	}
}
