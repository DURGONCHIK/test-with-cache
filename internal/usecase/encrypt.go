package usecase

import (
	"errors"
	"strings"
	"testing/internal/entities"
)

type EncryptUseCase interface {
	Encrypt(algo string, input string) (string, error)
}

type encryptUseCase struct {
	service *entities.EncryptService
}

func NewEncryptUseCase(service *entities.EncryptService) EncryptUseCase {
	return &encryptUseCase{service: service}
}

func (e *encryptUseCase) Encrypt(algo string, input string) (string, error) {
	algo = strings.ToLower(algo)
	hasher, ok := e.service.Hashers[algo]
	if !ok {
		return "", errors.New("unsupported algorithm")
	}

	cacheKey := algo + ":" + input
	if cached, err := e.service.Cache.Get(cacheKey); err == nil && cached != "" {
		return cached, nil
	}

	result := hasher(input)
	_ = e.service.Cache.Set(cacheKey, result)
	return result, nil
}
