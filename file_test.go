package dictionary

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileGetKey(t *testing.T) {
	var dataFileTest = []struct {
		name        string
		fileContent string
		searchKey   string
		out         string
	}{
		{"Firt line key", "hello hola\ngoodbye adeu\n", "hello", "hola"},
		{"Middle line key", "hello hola\ngoodbye adeu\nfrog granota\n", "goodbye", "adeu"},
		{"Last line key", "hello hola\ngoodbye adeu\nfrog granota\n", "frog", "granota"},
	}
	for _, tt := range dataFileTest {
		f, err := os.CreateTemp(os.TempDir(), "dictFileTest")
		if err != nil {
			t.Error(err)
		}
		defer os.Remove(f.Name()) // clean up
		err = ioutil.WriteFile(f.Name(), []byte(tt.fileContent), 0644)
		if err != nil {
			t.Error(err)
		}
		df, err := NewFile(f.Name(), 5)
		if err != nil {
			t.Error(err)
		}
		result, err := df.Get([]byte(tt.searchKey))
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, []byte(tt.out)) {
			t.Errorf("got %q, want %q", result, tt.out)
		}
	}
}
