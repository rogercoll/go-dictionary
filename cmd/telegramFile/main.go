package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/rogercoll/go-dictionary"
)

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func main() {
	dbdir, ok := os.LookupEnv("FILE_DB")
	if !ok {
		log.Fatal("FILE_DB is not present")
	}
	bckdir, ok := os.LookupEnv("BACKUP_FILE")
	if !ok {
		log.Fatal("BACKUP_FILE is not present")
	}
	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err := copy(dbdir, bckdir)
				if err != nil {
					log.Println(err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
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
