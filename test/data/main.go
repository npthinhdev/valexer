package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	head, err := ioutil.ReadFile("head.txt")
	if err != nil {
		panic(err)
	}
	testcase, err := ioutil.ReadFile("testcase.txt")
	if err != nil {
		panic(err)
	}
	solution, err := ioutil.ReadFile("solution.txt")
	if err != nil {
		panic(err)
	}

	mess := fmt.Sprintf("%s%s\n", testcase, solution)
	re := regexp.MustCompile(`(package.*)|(import.*")|(import(?s).*"\n\))`)
	mess = string(head) + re.ReplaceAllString(mess, "")

	f, err := os.Create("test.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(mess)
	if err != nil {
		panic(err)
	}

}
