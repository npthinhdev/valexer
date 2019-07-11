package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
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

	reFormat(&mess)
	msg := getTesting(&mess)
	fmt.Print(msg)

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(mess)
	if err != nil {
		panic(err)
	}
}

func reFormat(mess *string) {
	apiURL := "https://play.golang.org/fmt"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("imports", "true")
	body := getPost(apiURL, data)

	var msg struct {
		Body string `json:"body"`
	}

	_ = json.Unmarshal(body, &msg)
	*mess = msg.Body
}

func getTesting(mess *string) string {
	apiURL := "https://play.golang.org/compile"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("version", "2")
	data.Set("withVet", "true")
	body := getPost(apiURL, data)

	var msg struct {
		Errors string
		Events []struct {
			Message string `json:"message"`
		}
	}

	_ = json.Unmarshal(body, &msg)
	if msg.Errors != "" {
		return msg.Errors
	}
	return msg.Events[0].Message
}

func getPost(apiURL string, data url.Values) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}
