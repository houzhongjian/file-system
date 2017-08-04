package main

import (
	"net/http"
	"log"
)

func ServeHTTP() {
	http.HandleFunc("/upload", HandleUpload)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}
