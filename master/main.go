package main

import (
	"log"
	"github.com/houzhongjian/file-system/db"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	InitSystem()
	db.InitDbs()
	ServeHTTP()
}
