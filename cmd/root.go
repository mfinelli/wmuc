package cmd

import "fmt"
import "os"

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "wmuc",
	Short: "A git repository manager.",
	Long: `Wind Me Up, Chuck

wmuc is a way to manage multiple git repositories in a specific directory
structure.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
