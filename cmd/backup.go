package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)


func init() {
  rootCmd.AddCommand(backupCmd)
}


var backupCmd = &cobra.Command{
	Use: "backup [OPTIONS]",
	Short: "Creates a zip archive of the local tracked files",
	Long: `Creates a zip archive of the local tracked files
	in ~/.dotsync/backup directory`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Performing Backup...")
	},
}
