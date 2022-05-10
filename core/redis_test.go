package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var ctx = context.Background() // go-redis必须参数
var rdb *redis.Client

func TestRedis(t *testing.T) {
	err := initRedis2()
	if err != nil {
		log.Fatal(err)
	}

	// set
	err = rdb.Set(ctx, "Key2", "value", 0).Err()
	if err != nil {
		log.Println(err)
	}

	// get
	value := rdb.Get(ctx, "Key2")
	fmt.Println(value)
}

func initRedis2() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "120.25.153.110:6379",
		Password: "austin",
		DB:       0,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}
