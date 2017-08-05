package main

import "net/http"

func ServHTTP() {
	http.HandleFunc("/notice", HandleNotice)
}

func HandleNotice(w http.ResponseWriter, r *http.Request)  {

}
