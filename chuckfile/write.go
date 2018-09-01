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

package chuckfile

import "fmt"
import "time"

import "github.com/mfinelli/wmuc/parser"

func ProjectArrayToChuckfile(projects []parser.Project, version string,
	now time.Time) string {
	output := fmt.Sprintf("# wmuc v%s generated on: %s\n\n", version,
		now.Format(time.RFC1123))

	return output
}
