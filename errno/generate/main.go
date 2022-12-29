package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open(path.Join(wd, "types.go"))
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var checkVars bool
	var errorStrings = `[]string{
`
	for fileScanner.Scan() {
		if checkVars == true {
			if fileScanner.Text() == ")" {
				break
			}

			errorStrings = errorStrings + fmt.Sprintf(`	"%s",
`, strings.TrimSpace(strings.TrimSuffix(fileScanner.Text(), " Error = iota")))

		}
		if fileScanner.Text() == "const (" {
			checkVars = true
		}
	}

	toWrite := fmt.Sprintf(`// This file was generated, do not manually edit. 
package errno

var errorStrings = %s}
`, errorStrings)

	err = os.WriteFile(path.Join(wd, "error_strings.go"), []byte(toWrite), 0755)
	if err != nil {
		panic(err)
	}
}
