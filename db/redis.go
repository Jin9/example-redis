package db

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

const (
	addr     = "localhost:6379"
	password = ""
	db       = 0
)

func createNewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SetData is used for store data to redis
func SetData(key string, value interface{}, dur time.Duration) error {
	client, err := createNewClient()

	if err != nil {
		return errors.New("Connect to redis fail")
	}

	err = client.Set(key, value, dur).Err()
	if err != nil {
		return errors.New("Cannot set value to memory")
	}

	return nil
}

// GetData is used for get data from redis
func GetData(key string) (string, error) {
	client, err := createNewClient()

	if err != nil {
		return "", errors.New("Connect to redis fail")
	}

	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("Key is not exists")
	} else if err != nil {
		return "", errors.New("Cannot get value from memory")
	}

	return value, nil
}
