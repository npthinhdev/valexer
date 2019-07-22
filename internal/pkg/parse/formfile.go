package parse

import (
	"io/ioutil"
	"log"
	"net/http"
)

// FormFile is parsing content from form file
func FormFile(r *http.Request) ([]byte, error) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()
	log.Println(handler.Header)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return content, nil
}
