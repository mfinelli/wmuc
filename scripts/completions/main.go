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

import wmuc "github.com/mfinelli/wmuc/cmd"

func main() {
	err := wmuc.RootCmd.GenBashCompletionFile("./bash.completion")
	if err != nil {
		fmt.Println(1)
		os.Exit(1)
	}

	err = wmuc.RootCmd.GenZshCompletionFile("./zsh.completion")
	if err != nil {
		fmt.Println(1)
		os.Exit(1)
	}
}
