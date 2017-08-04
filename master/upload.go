package main

import (
	"net/http"
	"github.com/houzhongjian/file-system/system"
	"log"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	path, err := system.FileStoredPath()
	if err != nil {
		log.Printf("%+v\n", err)
		http.Error(w, "服务器异常", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(path))
}
