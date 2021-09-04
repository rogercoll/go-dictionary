package dictionary

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

// IMPORTANT: The best approach would be that this struct would be a generic file and use a sidecar container for S3 file retriving and push
// File format:
// key value
// key2 value2

// key-value file on S3 bucket
type File struct {
	filePath  string
	mapvalues map[string][]byte
	lastmap   time.Time
	//cacheTime in seconds
	cacheTime time.Duration
}

func (f *File) mapFile() error {
	file, err := os.Open(f.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	keyvalues := make(map[string][]byte)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), " ")
		if len(entry) < 2 {
			return errors.New("Invalid entry")
		}
		keyvalues[entry[0]] = []byte(strings.Join(entry[1:], " "))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	f.mapvalues = keyvalues
	return nil
}

func (f *File) inCacheTime() bool {
	return f.lastmap.After(time.Now().Add(-time.Second * f.cacheTime))
}

func NewFile(filePath string, cacheTime time.Duration) (*File, error) {
	return &File{filePath, nil, time.Now(), cacheTime}, nil
}

func (s *File) Get(key []byte) ([]byte, error) {
	if !s.inCacheTime() {
		err := s.mapFile()
		if err != nil {
			return nil, err
		}
	}
	keys := reflect.ValueOf(s.mapvalues).MapKeys()
	return keys[rand.Intn(len(keys))].Bytes(), nil
}

func (s *File) GetAll() ([]Entry, error) {
	return nil, nil
}

func (s *File) Insert(key []byte, value []byte) error {
	return nil
}

func (s *File) Delete(key []byte) error {
	return nil
}
