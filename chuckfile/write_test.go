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
	head := "# wmuc v0.1.0 generated on: Sat, 01 Sep 2018 10:30:00 UTC\n"

	tests := []struct {
		projects []parser.Project
		exp      string
	}{
		{make([]parser.Project, 0), head},
		{[]parser.Project{
			parser.Project{"", []parser.Repo{
				parser.Repo{Url: "repo1"},
				parser.Repo{Url: "repo2"},
			}},
		}, head + "\nrepo \"repo1\"\nrepo \"repo2\"\n"},
		{[]parser.Project{
			parser.Project{"zzz", []parser.Repo{
				parser.Repo{Url: "repo1", Branch: "dev"},
				parser.Repo{Url: "repo2"},
			}},
			parser.Project{"", []parser.Repo{
				parser.Repo{Url: "repo3"},
				parser.Repo{Url: "repo4"},
			}},
			parser.Project{"aaa", []parser.Repo{
				parser.Repo{Url: "repo5"},
				parser.Repo{Url: "repo6", Branch: "dev"},
			}},
		}, head + "\nrepo \"repo3\"\nrepo \"repo4\"\n\n" +
			"project \"aaa\" do\n    repo \"repo5\"\n    " +
			"repo \"repo6\", branch: \"dev\"\nend\n\n" +
			"project \"zzz\" do\n    repo \"repo1\", " +
			"branch: \"dev\"\n    repo \"repo2\"\nend\n"},
	}

	for _, test := range tests {
		str := chuckfile.ProjectArrayToChuckfile(test.projects,
			version, now)
		if str != test.exp {
			t.Errorf("%s returned bad chuckfile: %q (expected: %q)",
				test.projects, str, test.exp)
		}
	}
}
