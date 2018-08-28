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

package parser

import "fmt"
import "strings"

type Project struct {
	Path  string
	Repos []Repo
}

type Repo struct {
	Url    string
	Branch string
}

func (p Project) String() string {
	repos := make([]string, len(p.Repos))
	for i, r := range p.Repos {
		repos[i] = r.String()
	}

	return fmt.Sprintf("%q: [%s]", p.Path, strings.Join(repos, ", "))
}

func (r Repo) String() string {
	if r.Branch == "" {
		return fmt.Sprintf("%s", r.Url)
	}

	return fmt.Sprintf("%s(%s)", r.Url, r.Branch)
}
