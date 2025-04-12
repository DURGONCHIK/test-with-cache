package entities

type EncryptService struct {
	Hashers map[string]func(string) string
	Cache   Cache
}
type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
