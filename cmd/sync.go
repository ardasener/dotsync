package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use: "sync [OPTIONS]",
	Short: "Syncs the files with GIST",
	Long: `Compares the local files with the ones on the GIST,
	old local files are replaced with the newer ones from GIST,
	new local files are uploaded to the GIST`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Syncing...")
	},
}
