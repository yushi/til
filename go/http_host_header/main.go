package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	req.Host = "hoge"
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(b))
}
