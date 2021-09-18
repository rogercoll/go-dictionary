package main

import (
	"io"
	"log"
	"os"

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
	token, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("TELEGRAM_TOKEN is not present")
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
