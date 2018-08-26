package main

import "errors"
import "fmt"
import "io/ioutil"
import "os"

import "github.com/mfinelli/wmuc/parser"

func main() {
	input, err := readChuckfile()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results := parser.Parse(input)

	for _, project := range results {
		fmt.Printf("For project: %q\n", project.Path)

		for _, repo := range project.Repos {
			fmt.Printf("\trepo: %s\n", repo.Url)
		}
	}
}

func readChuckfile() (string, error) {
	allowed := []string{"chuckfile", "Chuckfile", "CHUCKFILE"}

	for _, filename := range allowed {
		bytes, err := ioutil.ReadFile(filename)

		if err == nil {
			return string(bytes), nil
		}
	}

	return "", errors.New("unable to load a chuckfile in this directory")
}
