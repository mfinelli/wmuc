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
