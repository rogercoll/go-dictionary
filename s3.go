package dictionary

// key-value file on S3 bucket
type S3File struct {
	bucket string
	object string
}

func (s *S3File) Get(key []byte) ([]byte, error) {

	return nil, nil
}

func (s *S3File) GetAll() ([]Entry, error) {
	return nil, nil
}

func (s *S3File) Insert(key []byte, value []byte) error {
	return nil
}

func (s *S3File) Delete(key []byte) error {
	return nil
}
