package utils

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

const (
	LOGIN_CODE_KEY    string        = "login:code:"
	LOGIN_CODE_TTL    time.Duration = 120 * time.Second
	LOGIN_TOKEN       string        = "login:token:"
	LOGIN_TOKEN_TTL   time.Duration = 30 * time.Minute
	CACHE_SHOP_KEY    string        = "cache:shop:"
	CACHE_SHOP_TTL    time.Duration = 30 * time.Minute
	CACHE_SHOP_LIST   string        = "cache:shop:list"
	CACHE_NIL_TTL     time.Duration = 5 * time.Minute
	LOCK_TTL          time.Duration = time.Second
	SECKILL_STOCK_KEY string        = "seckill:stock:"
	LOCK              string        = "lock:"
)

var rdb *redis.Client

var ctx = context.Background()

var UNLOCK_LUA string

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.101:6379",
		Password: "zhuyao",
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err.Error())
	}
	log.Println("redis 连接成功")
	UNLOCK_LUA, err = readUnlockLua("./unlock.lua")
	if err != nil {
		panic(err.Error())
	}
}

func readUnlockLua(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	rd := bufio.NewReader(f)
	script, err := io.ReadAll(rd)
	if err != nil {
		return "", err
	}
	return string(script), nil
}
