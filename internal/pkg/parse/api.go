package parse

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get is parsing json from api
func Get(apiURL string) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", apiURL, nil)
	r.Header.Add("Accept", "application/json")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

// Post is parsing json from api
func Post(apiURL string, data url.Values) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

// Put is parsing json from api
func Put(apiURL string, data url.Values) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("PUT", apiURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

// Delete is parsing json from api
func Delete(apiURL string) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("DELETE", apiURL, nil)
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}
