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

package parser_test

import "testing"
import "github.com/mfinelli/wmuc/parser"

func TestProjectString(t *testing.T) {
	tests := []struct {
		p   parser.Project
		exp string
	}{
		{parser.Project{Path: ""}, "\"\": []"},
		{parser.Project{Path: "", Repos: []parser.Repo{
			{Url: "test1"}, {Url: "test2"},
		}}, "\"\": [test1, test2]"},
		{parser.Project{Path: "test", Repos: []parser.Repo{
			{Url: "test3", Branch: "dev"}, {Url: "test4"},
		}}, "\"test\": [test3(dev), test4]"},
	}

	for _, test := range tests {
		if test.p.String() != test.exp {
			t.Errorf("String didn't match %s (got: %s)",
				test.exp, test.p.String())
		}
	}
}

func TestRepoString(t *testing.T) {
	tests := []struct {
		r   parser.Repo
		exp string
	}{
		{parser.Repo{Url: "https://github.com/mfinelli/wmuc.git"},
			"https://github.com/mfinelli/wmuc.git"},
		{parser.Repo{Url: "https://github.com/mfinelli/wmuc.git",
			Branch: "master"},
			"https://github.com/mfinelli/wmuc.git(master)"},
	}

	for _, test := range tests {
		if test.r.String() != test.exp {
			t.Errorf("String didn't match %s (got: %s)",
				test.exp, test.r.String())
		}
	}
}

func TestRepoDirFromPath(t *testing.T) {
	tests := []struct {
		r   parser.Repo
		exp string
	}{
		{parser.Repo{Url: "https://github.com/mfinelli/wmuc.git"},
			"wmuc"},
	}

	for _, test := range tests {
		if test.r.DirFromPath() != test.exp {
			t.Errorf("DirFromPath didn't match %s (got: %s)",
				test.exp, test.r.DirFromPath())
		}
	}
}
