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

package chuckfile_test

import "testing"
import "time"
import "github.com/mfinelli/wmuc/chuckfile"
import "github.com/mfinelli/wmuc/parser"

func TestProjectArrayToChuckfile(t *testing.T) {
	now := time.Date(2018, time.September, 1, 10, 30, 0, 0, time.UTC)
	version := "0.1.0"

	tests := []struct {
		projects []parser.Project
		exp      string
	}{
		{make([]parser.Project, 0), "# wmuc v0.1.0 generated on: " +
			"Sat, 01 Sep 2018 10:30:00 UTC\n\n"},
	}

	for _, test := range tests {
		if chuckfile.ProjectArrayToChuckfile(test.projects, version,
			now) != test.exp {
			t.Errorf("%s returned bad chuckfile: %q",
				test.projects, test.exp)
		}
	}
}
