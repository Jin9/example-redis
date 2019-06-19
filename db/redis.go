package db

import (
	"log"
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
func SetData(key string, value interface{}, dur time.Duration) (err error) {
	client, err := createNewClient()

	if err != nil {
		panic(err)
	}

	err = client.Set(key, value, dur).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

// GetData is used for get data from redis
func GetData(key string) string {
	client, err := createNewClient()

	if err != nil {
		panic(err)
	}

	value, err := client.Get(key).Result()
	if err == redis.Nil {
		log.Panicln(err.Error())
	} else if err != nil {
		log.Panicln(err.Error())
	}

	return value
}
