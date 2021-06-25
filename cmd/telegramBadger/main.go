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
	err = b.Insert([]byte("hello"), []byte("hola"))
	if err != nil {
		log.Fatal(err)
	}
	result, err := b.Get([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value:%v\n", string(result))
	b.ViewAll()
}
