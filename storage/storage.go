package storage

type Store interface {
	Save(data []byte) error
	Load() ([]byte, error)
	GetFileName() string
}

type Storage struct {
	filename string
}

func (s *Storage) GetFileName() string {
	return s.filename
}