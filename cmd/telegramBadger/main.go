package main

import (
	"log"

	"github.com/rogercoll/go-dictionary"
)

func main() {
	b, err := dictionary.NewBadgerDB("/tmp/badger")
	if err != nil {
		log.Fatal(err)
	}
	RunBot(b, token)
}
