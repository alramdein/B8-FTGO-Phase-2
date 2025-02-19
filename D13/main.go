package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://imdb236.p.rapidapi.com/imdb/tt7631058"

	req, _ := http.NewRequest("GET", url, nil)

	// jsonBody := []byte(`{"client_message": "hello, server!"}`)
	// bodyReader := bytes.NewReader(jsonBody)
	// req, _ := http.NewRequest("POST", url, bodyReader)

	req.Header.Add("x-rapidapi-key", "77a04668bbmshc610752f2f59315p127843jsnbaff27b0f187")
	// req.Header.Add("x-rapidapi-host", "imdb236.p.rapidapi.com")

	reqBody, _ := io.ReadAll(req.Body)
	fmt.Printf("REQ BODY: %v\n", string(reqBody))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
