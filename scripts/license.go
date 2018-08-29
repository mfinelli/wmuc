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
	text     string
	projects []string
}

func main() {
	exclusions := []string{
		"github.com/spf13/cobra/cobra/cmd/testdata/LICENSE.golden",
	}

	lfiles := make(map[string][]string)
	licenses := make(map[string]string)
	final := make(map[string]licenseText)

	// find all the files that begin with LICENSE in the vendor directory
	// and associate the path/to/the/license with the project it belongs
	// to.
	// go makes this easy since we can get a url to the project from the
	// package name
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

			if strings.HasPrefix(strings.ToUpper(info.Name()),
				"LICENSE") {
				// strip the vendor/ prefix for the project
				// url and then add the license file to the
				// list for that project
				proj := filepath.Dir(path)[7:]
				lfiles[proj] = append(lfiles[proj], path)
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// concatenate all the licenses for each individual project
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

	// https://github.com/pelletier/go-buffruneio/issues/6
	licenses["github.com/pelletier/go-buffruneio"] = `MIT License

Copyright (c) 2018 Thomas Pelletier

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`

	// calculate a checksum of the license text so we can group projects
	// that have the same license together
	for project := range licenses {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(licenses[project])))

		lt := findOrCreate(sum, licenses[project], final)
		lt.projects = append(lt.projects, project)
		final[sum] = lt
	}

	// sort the sums so that we get the same output every time we run it
	sums := []string{}
	for sum := range final {
		sums = append(sums, sum)
	}
	sort.Strings(sums)

	// build up our desired go-file:
	// package legal
	// var thirdparty = `ALL THE LICENSES`
	output := "package legal\n\nvar thirdparty = `"
	for _, sum := range sums {
		output += fmt.Sprintf("%s\n", final[sum])
	}
	output += "`\n"

	// write the file...
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
