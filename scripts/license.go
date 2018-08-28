package main

import "crypto/md5"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "path/filepath"
import "sort"
import "strings"

type licenseText struct {
	text string
	projects []string
}

func main() {
	exclusions := []string{
		"github.com/spf13/cobra/cobra/cmd/testdata/LICENSE.golden",
	}

	lfiles := make(map[string][]string)
	licenses := make(map[string]string)
	final := make(map[string]licenseText)

	err := filepath.Walk("vendor",
		func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, exclude := range exclusions {
			if path == filepath.Join("vendor", exclude) {
				return nil
			}
		}

		if strings.HasPrefix(strings.ToUpper(info.Name()), "LICENSE") {
			proj := filepath.Dir(path)[7:]
			lfiles[proj] = append(lfiles[proj], path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for project := range lfiles {
		str := ""

		for _, license := range lfiles[project] {
			bytes, err := ioutil.ReadFile(license)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			str += string(bytes)
		}

		licenses[project] = str
	}

	for project := range licenses {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(licenses[project])))

		lt := findOrCreate(sum, licenses[project], final)
		lt.projects = append(lt.projects, project)
		final[sum] = lt
	}

	sums := []string{}

	for sum := range final {
		sums = append(sums, sum)
	}

	sort.Strings(sums)
	output := "package legal\n\nvar thirdparty = `"

	for _, sum := range sums {
		output += fmt.Sprintf("%s\n", final[sum])
	}

	output += "`\n"

	outputPath := filepath.Join(".", "legal", "third_party.go")
	err = ioutil.WriteFile(outputPath, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func findOrCreate(search, t string, m map[string]licenseText) licenseText {
	if p, ok := m[search]; ok {
		return p
	}

	// not found: initialize, add, and return
	m[search] = licenseText{text: t}
	return m[search]
}

func (lt licenseText) String() string {
	sort.Strings(lt.projects)

	str := ""
	for _, proj := range lt.projects {
		str += fmt.Sprintf("https://%s\n", proj)
	}

	str += fmt.Sprintf("\n%s\n\n", lt.text)
	str += "--------------------------------------------------\n\n"

	return str
}
