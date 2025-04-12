package usecase_test

import (
	"testing"
	"testing/internal/entities"
	"testing/internal/usecase"
)

// Фейковый кэш для теста
type fakeCache struct {
	data map[string]string
}

func (f *fakeCache) Get(key string) (string, error) {
	return f.data[key], nil
}

func (f *fakeCache) Set(key, value string) error {
	f.data[key] = value
	return nil
}

// Фейковый хэшер
type fakeHasher struct{}

func (f *fakeHasher) Hash(input string) string {
	return "hashed:" + input
}

func TestEncryptService_Encrypt(t *testing.T) {
	cache := &fakeCache{data: make(map[string]string)}
	hashers := map[string]func(string) string{
		"md5": (&fakeHasher{}).Hash,
	}
	service := &entities.EncryptService{
		Hashers: hashers,
		Cache:   cache,
	}

	useCase := usecase.NewEncryptUseCase(service)

	input := "test"
	expected := "hashed:" + input

	// Первый вызов — должен захешировать
	result, err := useCase.Encrypt("md5", input)
	if err != nil || result != expected {
		t.Fatalf("unexpected result: %v, error: %v", result, err)
	}

	// Второй вызов — должен достать из кэша
	resultCached, err := useCase.Encrypt("md5", input)
	if err != nil || resultCached != expected {
		t.Fatalf("unexpected cached result: %v, error: %v", resultCached, err)
	}
}
