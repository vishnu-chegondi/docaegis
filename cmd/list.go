package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all gaurded folders",
	Long:  `Get the gaurded list of directories and files from sqllite database`,
	Run: func(cmd *cobra.Command, args []string) {
		listGaurdedFiles()
	},
}

// Function to connect to database and list all the
// gaurded files which are gaurded
func listGaurdedFiles() {
	GetAllFilesGaurded()
}
