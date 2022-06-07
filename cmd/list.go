package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all guarded folders",
	Long:  `Get the guarded list of directories and files from sqllite database`,
	Run: func(cmd *cobra.Command, args []string) {
		listguardedFiles()
	},
}

// Function to connect to database and list all the
// guarded files which are guarded
func listguardedFiles() {
	GetAllFilesguarded()
}
