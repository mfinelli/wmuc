package cmd

import "fmt"
import "os"

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "wmuc",
	Short: "A git repository manager.",
	Long: `Wind Me Up, Chuck

wmuc is a way to manage multiple git repositories in a
specific directory structure.

wmuc  Copyright (C) 2018  Mario Finelli
This program comes with ABSOLUTELY NO WARRANTY; for details run
"wmuc about warranty". This is free software, and you are welcome to
redistribute it under certain conditions; run "wmuc about" for details.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
