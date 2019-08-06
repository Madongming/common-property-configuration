package property

type InterStorage interface {
	Create(key string, data []byte) error
	Update(key string, data []byte) error
	Read(key string) ([]byte, error)
	Delete(key string) ([]byte, error)
}
