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

package main

import "fmt"
import "os"
import "path/filepath"

import "github.com/spf13/cobra"
import "github.com/spf13/cobra/doc"
import "github.com/spf13/viper"

import wmuc "github.com/mfinelli/wmuc/cmd"

var rootCmd = &cobra.Command{
	Use:   "doc",
	Short: "generate man pages",
	Run: func(cmd *cobra.Command, args []string) {
		header := &doc.GenManHeader{
			Title:   "WMUC",
			Section: "1",
		}

		cwd, _ := os.Getwd()
		err := os.MkdirAll(filepath.Join(cwd,
			viper.GetString("path")), 0755)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = doc.GenManTree(wmuc.RootCmd, header,
			viper.GetString("path"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {
	rootCmd.Flags().StringP("path", "", ".", "man page output path")
	viper.BindPFlag("path", rootCmd.Flags().Lookup("path"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
