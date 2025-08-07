package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New(opts ...Option) *redis.Client {
	o := options{
		addr:     "127.0.0.1:6379",
		password: "",
		db:       0,
	}

	for _, opt := range opts {
		opt(&o)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     o.addr,
		Password: o.password,
		DB:       o.db,
	})

	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}

	return rdb
}
