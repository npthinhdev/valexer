package parse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"

	"github.com/npthinhdev/valexer/internal/app/types"
)

// GetTesting get result form play.golang
func GetTesting(id string, solution []byte) string {
	head := `package main
func main() {
	matchAll := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{F: TestSolution}}
	testing.Main(matchAll, tests, nil, nil)
}`
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	body := Get(apiURL)
	exer := types.Exercise{}
	_ = json.Unmarshal(body, &exer)
	testcase := exer.Testcase
	mess := testcase + string(solution)
	re := regexp.MustCompile(`(package.*)|(import.*")|(import(?s).*"\n\))`)
	mess = string(head) + re.ReplaceAllString(mess, "")
	getFormat(&mess)
	result := getCompile(&mess)
	return result
}

func getFormat(mess *string) {
	apiURL := "https://play.golang.org/fmt"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("imports", "true")
	body := Post(apiURL, data)
	var msg struct {
		Body string `json:"body"`
	}
	_ = json.Unmarshal(body, &msg)
	*mess = msg.Body
}

func getCompile(mess *string) string {
	apiURL := "https://play.golang.org/compile"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("version", "2")
	data.Set("withVet", "true")
	body := Post(apiURL, data)
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
