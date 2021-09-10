package cmd

import (
  "github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "dotsync",
	Short: "DotSync: sync your dotfiles using Github GIST",
  Long: `Allows syncing and backing up your dotfile (or any files really)
	 using Github's GIST feature (WIP)`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
