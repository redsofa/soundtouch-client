package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://192.168.2.10:8090/select"
	fmt.Println("URL:>", url)

	var xml = []byte(`
		<ContentItem source="INTERNET_RADIO" sourceAccount="" location="49358">
    		<itemName>Ambiance Reggae</itemName>
		</ContentItem>
	`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xml))
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
