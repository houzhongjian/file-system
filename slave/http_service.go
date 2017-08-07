package main

import (
	"net/http"
	"log"
)

func ServHTTP() {
	http.HandleFunc("/slave/notice", HandleSlaveNotice)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}

//通知备份.
func HandleSlaveNotice (w http.ResponseWriter, r *http.Request) {

}

