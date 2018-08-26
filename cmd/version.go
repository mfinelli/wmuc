package cmd

import "fmt"

import "github.com/spf13/cobra"

const VERSION = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show wmuc version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wmuc: Wind Me Up, Chuck v%s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
