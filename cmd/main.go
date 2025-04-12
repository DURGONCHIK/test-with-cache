package main

import (
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"testing/config"
	"testing/internal/controller"
	"testing/internal/entities"
	"testing/internal/repository"
	"testing/internal/usecase"
	"testing/pkg/hashutil"
)

func main() {
	// Инициализация кеша
	cfg := config.Load()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	cache := repository.NewRedisCache(redisClient)

	// Инициализация шифраторов
	hashers := map[string]func(string) string{
		"md5":    hashutil.MD5Hash,
		"sha256": hashutil.SHA256Hash,
	}

	// Инициализация сервиса
	encryptService := &entities.EncryptService{
		Hashers: hashers,
		Cache:   cache,
	}

	// UseCase
	encryptUseCase := usecase.NewEncryptUseCase(encryptService)

	// Контроллер
	encryptHandler := controller.NewEncryptHandler(encryptUseCase)

	// Роутинг
	http.HandleFunc("/encrypt", encryptHandler.Encrypt)

	// Сервер
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
