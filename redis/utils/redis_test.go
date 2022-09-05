package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNum(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(TryLock("lock"))
		time.Sleep(time.Millisecond * 300)
	}
}
func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
func TestRedis(t *testing.T) {
	lock := NewSimpleLock("123")
	lock.TryLock(time.Minute)
	lock.Unlock()
}
