package main

import (
	"log"
	"os"

	"github.com/rogercoll/go-dictionary"
)

func main() {
	dbdir, ok := os.LookupEnv("BADGERDB_DIR")
	if !ok {
		log.Fatal("BADGERDB_DIR is not present")
	}
	b, err := dictionary.NewBadgerDB(dbdir)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()
	RunBot(b, token)
}
