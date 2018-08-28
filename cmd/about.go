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

var warrantyCmd = &cobra.Command{
	Use:   "warranty",
	Short: "show warranty information",
	Run: func(cmd *cobra.Command, args []string) {
		legal.PrintWarranty()
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
	aboutCmd.AddCommand(warrantyCmd)
}
