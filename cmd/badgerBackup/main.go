package main

import (
	"log"
	"os"
	"time"

	"github.com/rogercoll/go-dictionary"
)

func main() {
	dbdir, ok := os.LookupEnv("BADGERDB_DIR")
	if !ok {
		log.Fatal("BADGERDB_DIR is not present")
	}
	bckdir, ok := os.LookupEnv("BADGERDB_BCK_DIR")
	if !ok {
		log.Fatal("BADGERDB_BCK_DIR is not present")
	}
	b, err := dictionary.NewBadgerDB(dbdir)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()
	f, err := os.Create(bckdir + "badgerdb-" + time.Now().Format("01-02-2006") + ".bak")
	if err != nil {
		log.Fatal(err)
	}
	err = b.Backup(f)
	if err != nil {
		log.Fatal(err)
	}
}
