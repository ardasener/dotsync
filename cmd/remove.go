package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)


func init() {
  rootCmd.AddCommand(removeCmd)
}


var removeCmd = &cobra.Command{
	Use: "remove",
	Short: "Remove file from DotSync interactively",
	Long: `Removed files will not be tracked by DotSync anymore and 
	they will be deleted from the GIST in the next sync command`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Removing...")
	},
}


