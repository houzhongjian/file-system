package main

import (
	"github.com/houzhongjian/file-system/db"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ServHTTP()
	db.InitDbs()
}

