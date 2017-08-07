package main

import (
	"net/http"
	"log"
	"github.com/houzhongjian/file-system/system"
)

func ServeHTTP() {
	http.HandleFunc("/upload", HandleUpload)
	http.StripPrefix("/assets/", http.FileServer(http.Dir(system.FILE_SYSTEM_PATH)))
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}
