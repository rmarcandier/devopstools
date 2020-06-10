package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json: "url"`
	Headers map[string]string `json: "headers"`
}

func (r *GetResponse) ToString() string {
	s := fmt.Sprintf("GET Response\nOrigin IP: %s\nRequest URL: %s\n",
		r.Origin, r.URL)
	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s

}

func main() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to read response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	respContent := GetResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Println(respContent.ToString())

}
