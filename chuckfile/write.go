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
import "sort"
import "strings"
import "time"

import "github.com/mfinelli/wmuc/parser"

func ProjectArrayToChuckfile(projects []parser.Project, version string,
	now time.Time) string {
	output := fmt.Sprintf("# wmuc v%s generated on: %s\n\n", version,
		now.Format(time.RFC1123))

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Path < projects[j].Path
	})

	for _, project := range projects {
		output += fmt.Sprintf("%s\n", formatProject(project))
	}

	return output[:len(output)-1]
}

func formatProject(project parser.Project) string {
	formatted := ""

	indent := "    "
	if project.Path == "" {
		indent = ""
	}

	repos := make([]string, len(project.Repos))
	for i, repo := range project.Repos {
		out := fmt.Sprintf("%srepo %q", indent, repo.Url)
		if repo.Branch != "" {
			out += fmt.Sprintf(", branch: %q", repo.Branch)
		}
		repos[i] = fmt.Sprintf("%s\n", out)
	}

	sort.Strings(repos)

	if project.Path != "" {
		formatted += fmt.Sprintf("project %q do\n", project.Path)
	}

	formatted += strings.Join(repos, "")

	if project.Path != "" {
		formatted += "end\n"
	}

	return formatted
}
