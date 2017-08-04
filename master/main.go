package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitSystem()
	ServeHTTP()
}
