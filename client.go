package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8081/api/data/sms"

	// Make a GET request to the server
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the JSON data
	fmt.Println(string(body))
}
