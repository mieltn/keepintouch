package main

import (
	"log"

	"github.com/mieltn/keepintouch/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
}