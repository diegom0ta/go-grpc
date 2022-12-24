package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const booksURL = "https://api.itbook.store/1.0/"

func BooksHandler() *http.Response {
	r, err := http.Get(booksURL)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func GetBytes(r *http.Response) map[byte]interface{} {
	result := make(map[byte]interface{})

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
