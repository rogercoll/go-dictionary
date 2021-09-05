package main

import (
	"log"
	"os"

	"github.com/rogercoll/go-dictionary"
)

func main() {
	dbdir, ok := os.LookupEnv("FILE_DB")
	if !ok {
		log.Fatal("FILE_DB is not present")
	}
	b, err := dictionary.NewFile(dbdir, 10)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()
	err = RunBot(b, token)
	if err != nil {
		log.Fatal(err)
	}
}
