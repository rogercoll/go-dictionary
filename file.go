package dictionary

import (
	"bufio"
	"bytes"
	"errors"
	"os"
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
	f.lastmap = time.Now()
	return nil
}

func (f *File) inCacheTime() bool {
	return f.lastmap.After(time.Now().Add(-time.Second * f.cacheTime))
}

func NewFile(filePath string, cacheTime time.Duration) (*File, error) {
	return &File{filePath, nil, time.Unix(0, 0), cacheTime}, nil
}

func (s *File) Get(key []byte) ([]byte, error) {
	if !s.inCacheTime() {
		err := s.mapFile()
		if err != nil {
			return nil, err
		}
	}
	return s.mapvalues[string(key)], nil
}

func (s *File) GetAll() ([]Entry, error) {
	if !s.inCacheTime() {
		err := s.mapFile()
		if err != nil {
			return nil, err
		}
	}
	all := make([]Entry, len(s.mapvalues))
	i := 0
	for k, v := range s.mapvalues {
		all[i] = Entry{[]byte(k), v}
		i += 1
	}
	return all, nil
}

func (s *File) Insert(key []byte, value []byte) error {
	f, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	toString := make([]string, 2)
	toString[0] = string(key)
	toString[1] = string(value)
	if _, err = f.WriteString(strings.Join(toString, " ")); err != nil {
		return err
	}
	s.lastmap = time.Unix(0, 0)
	return nil
}

func (s *File) Delete(key []byte) error {
	f, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		keyFile := strings.Split(scanner.Text(), " ")
		if keyFile[0] != string(key) {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				return err
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	err = os.WriteFile(s.filePath, buf.Bytes(), 0666)
	if err != nil {
		return err
	}
	s.lastmap = time.Unix(0, 0)
	return nil
}
