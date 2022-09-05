package utils

import (
	"time"

	"github.com/google/uuid"
)

type SimpleLock struct {
	uuid uuid.UUID
	name string
}

func NewSimpleLock(id string) SimpleLock {
	return SimpleLock{
		uuid: uuid.New(),
		name: LOCK + id,
	}
}
func (s *SimpleLock) TryLock(timeout time.Duration) bool {
	// 存入redis的value是固定的值,无法分辨线程安全,会删除其他线程的锁 --> 存入uuid,判断解决删除其他线程的锁
	ok, _ := rdb.SetNX(ctx, s.name, s.uuid.String(), timeout).Result()
	return ok
}
func (s *SimpleLock) Unlock() {
	// 线程a当获取key-value 判断为自己的锁后被阻塞,(key-value超时失效),切换线程b添加新的锁,线程a恢复运行,del线程b的锁
	/*	版本一:
		lockUUID, _ := rdb.Get(ctx, s.name).Result()

		if lockUUID == s.uuid.String() {
			rdb.Del(ctx, s.name)
		}
	*/

	// 版本二:lua,保证删除锁的原子性
	rdb.Eval(ctx, UNLOCK_LUA, []string{s.name}, s.uuid.String())

}
