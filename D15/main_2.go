package main

import (
	"fmt"
	"io"
	"net/http"
)

func main2() {
	url := "https://api-sandbox.collaborator.komerce.id/tariff/api/v1/calculate?shipper_destination_id=31597&receiver_destination_id=368&weight=490&item_value=50000&cod=no"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "G1Fq3F3O2c1c01802576e7cfzzBXHRSE")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
