package cmd

import "github.com/spf13/cobra"

import "github.com/mfinelli/wmuc/legal"

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "show license and warranty information",
	Run: func(cmd *cobra.Command, args []string) {
		legal.PrintLegalInfo()
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
