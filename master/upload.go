package main

import (
	"net/http"
	"github.com/houzhongjian/file-system/system"
	"log"
	"io/ioutil"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	path, err := system.FileStoredPath()
	if err != nil {
		log.Printf("%+v\n", err)
		http.Error(w, "服务器异常", http.StatusInternalServerError)
		return
	}

	by, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	if !system.CreateFile(path+"1.jpg", by) {
		w.Write([]byte("上传失败"))
		return
	}
	w.Write([]byte("上传成功"))
}
