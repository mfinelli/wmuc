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
import "io/ioutil"
import "os"
import "path"
import "time"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"

import "github.com/mfinelli/wmuc/chuckfile"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a chuckfile from the repositories on disk",
	Run: func(cmd *cobra.Command, args []string) {
		if chuckfile.ChuckfileExists() && !viper.GetBool("initForce") {
			fmt.Println("A chuckfile already exists!")
			os.Exit(1)
		} else {
			output := chuckfile.ProjectArrayToChuckfile(
				chuckfile.WalkDirectory(), VERSION, time.Now())

			chuckfile.RemoveExistingChuckfile()

			err := ioutil.WriteFile(path.Join(".", "chuckfile"),
				[]byte(output), 0644)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false,
		"overwrite existing chuckfile")
	viper.BindPFlag("initForce", initCmd.Flags().Lookup("force"))
}
