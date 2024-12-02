package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"advent2024/util"
	"github.com/otiai10/copy"
)

func main() {
	var nextN int
	var err error
	if len(os.Args) > 1 {
		nextN, err = strconv.Atoi(os.Args[1])
		util.Panic(err)
	} else {
		nextN = getExistingMax() + 1
	}

	nextDir := fmt.Sprintf("advent%s", twoDigitInt(nextN))
	err = copy.Copy("adventN", nextDir)
	util.Panic(err)

	codeName := fmt.Sprintf("%s/%s.go", nextDir, nextDir)
	err = os.Rename(nextDir+"/adventN.go", codeName)
	util.Panic(err)

	testName := fmt.Sprintf("%s/%s_test.go", nextDir, nextDir)
	err = os.Rename(nextDir+"/adventN_test.go", testName)
	util.Panic(err)

	replacePackage(codeName, nextDir)
	replacePackage(testName, nextDir)
}

func getExistingMax() int {
	files, err := ioutil.ReadDir(".")
	util.Panic(err)

	var max int
	for _, f := range files {
		if !f.IsDir() || !strings.HasPrefix(f.Name(), "advent") {
			continue
		}
		n := f.Name()[len("advent"):]
		i, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		if i > max {
			max = i
		}
	}
	return max
}

func replacePackage(fileName, packageName string) {
	b, err := os.ReadFile(fileName)
	util.Panic(err)

	lines := strings.Split(string(b), "\n")
	lines[0] = "package " + packageName
	output := strings.Join(lines, "\n")

	err = os.WriteFile(fileName, []byte(output), 0644)
	util.Panic(err)
}

func twoDigitInt(i int) string {
	s := strconv.Itoa(i)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}
