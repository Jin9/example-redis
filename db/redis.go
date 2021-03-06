package db

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	addr     = "localhost:6379"
	password = ""
)

func createNewClient(channel int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       channel,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SetData is used for store data to redis
func SetData(key string, value interface{}, dur time.Duration, channel int) error {
	client, err := createNewClient(channel)

	if err != nil {
		log.Println(err.Error())
		return errors.New("Connect to redis fail")
	}

	err = client.Set(key, value, dur).Err()
	if err != nil {
		log.Println(err.Error())
		return errors.New("Cannot set value to memory")
	}

	return nil
}

// GetData is used for get data from redis
func GetData(key string, channel int) (string, error) {
	client, err := createNewClient(channel)

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("Connect to redis fail")
	}

	value, err := client.Get(key).Result()
	if err == redis.Nil {
		log.Println(err.Error())
		return "", errors.New("Key is not exists")
	} else if err != nil {
		log.Println(err.Error())
		return "", errors.New("Cannot get value from memory")
	}

	return value, nil
}

// GetExistsKey is used to check key in redis server is exists
func GetExistsKey(key string, channel int) (bool, error) {
	client, err := createNewClient(channel)

	if err != nil {
		log.Println(err.Error())
		return false, errors.New("Connect to redis fail")
	}

	val, err := client.Exists(key).Result()
	if err != nil {
		log.Println(err.Error())
		return false, errors.New("Cannot get exists key")
	}

	return val == 1, nil
}

// DelData is used for delete some key
func DelData(key string, channel int) error {
	client, err := createNewClient(channel)

	if err != nil {
		log.Println(err.Error())
		return errors.New("Connect to redis fail")
	}

	_, err = client.Del(key).Result()
	if err != nil {
		log.Println(err.Error())
		return errors.New("Cannot delete key")
	}

	return nil
}
