package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

}

func Post(endpoint string, postbody string, authorization string) string {

	baseUrl := "test"
	url := baseUrl + endpoint
	var jsonStr = []byte(postbody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer"+authorization)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return string(resp.Status)
}

func GET(endpoint string, authorization string) map[string]interface{} {
	baseUrl := "test"

	url := baseUrl + endpoint

	httpClient := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "bearer"+authorization)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}
