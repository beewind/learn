package utils

import "time"

type RedisData struct {
	ExpireTime time.Time
	Data       interface{}
}
