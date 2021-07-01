package backup

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"github.com/rogercoll/go-dictionary"
)

type Backup struct {
	Entries []dictionary.Entry `json:"entries"`
}

func MakeBackup(d dictionary.Dictionary, w io.Writer) (int, error) {
	entries, err := d.GetAll()
	if err != nil {
		return 0, err
	}
	if len(entries) < 1 {
		return 0, errors.New("Error: No stored entries, please add a definition with /add")
	}
	data, err := json.Marshal(Backup{entries})
	if err != nil {
		return 0, err
	}
	return w.Write(data)
}

//path to for new dictionary data
func RestoreBackup(r io.Reader, path string) (dictionary.Dictionary, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var data Backup
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	d, err := dictionary.NewBadgerDB(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range data.Entries {
		err = d.Insert(entry.Key, entry.Value)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}
