package db

import (
	"errors"
	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

var ErrNil = errors.New("redis: nil")

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address, // host:port of the redis server
		Password: "",      // no password set
		DB:       0,       // use default DB
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}
