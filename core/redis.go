// Package core
// @Description: 初始化redis
// @Author: Quan

package core

import (
	"bytedance-douyin/global"
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

// Redis Init redis connection. If it failed, program will be exit
func Redis() (rdb *redis.Client) {
	rdb, err := initRedis()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		os.Exit(-1)
	}
	return rdb
}

func initRedis() (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     global.GVA_CONFIG.Redis.Addr + ":" + global.GVA_CONFIG.Redis.Port,
		Password: global.GVA_CONFIG.Redis.Password,
		DB:       global.GVA_CONFIG.Redis.Db,
		PoolSize: global.GVA_CONFIG.Redis.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return rdb, err
}
