package dictionary

type Entry struct {
	Key   []byte
	Value []byte
}

type Dictionary interface {
	Get([]byte) ([]byte, error)
	//slices cannot be used as key maps
	GetAll() ([]Entry, error)
	Insert([]byte, []byte) error
	Delete([]byte) error
	Close()
}
