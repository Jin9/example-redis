package main

import (
	"fmt"
	_ "wasabi/router"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/sha3"
)

func hashUUID(keyValue string) (string, error) {
	h := sha3.New256()

	_, err := h.Write([]byte(keyValue))

	if err != nil {
		return "", err
	}

	keyBytes := h.Sum(nil)
	return fmt.Sprintf("%x", keyBytes), nil
}

func createNewClient(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	// e := router.Init()
	// e.Logger.Fatal(e.Start(":8880"))
}
