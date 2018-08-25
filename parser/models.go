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
