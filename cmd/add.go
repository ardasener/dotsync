package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add [FILE]",
	Short: "Add files to be tracked by DotSync",
	Long: `Any files added using this method will be tracked by DotSync
	and synced to the GIST when the sync command is used`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Adding...")
	},
}
