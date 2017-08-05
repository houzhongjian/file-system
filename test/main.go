package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"time"
	"fmt"
)

func main() {
	by, err := ioutil.ReadFile("./1.jpg")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1/upload", bytes.NewReader(by))
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	client := http.Client{}
	req.Header.Set("name", fmt.Sprintf("%d.jpg", time.Now().Unix()))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	log.Println("resp:", string(by))
}
