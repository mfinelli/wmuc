// wmuc: a git repository manager
// Copyright (C) 2018  Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
