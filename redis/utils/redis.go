package utils

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

func SetCode(phone string, code string) error {
	// 60s 过期
	return rdb.Set(ctx, LOGIN_CODE_KEY+phone, code, LOGIN_CODE_TTL).Err()
}

func GetCode(phone string) (string, error) {
	return rdb.Get(ctx, LOGIN_CODE_KEY+phone).Result()
}
func SetUser(token string, userStringSlice []string) error {
	token = LOGIN_TOKEN + token

	err := rdb.HSet(ctx, token, userStringSlice).Err()
	if err != nil {
		return err
	}
	err = refresh(token, LOGIN_TOKEN_TTL)
	return err
}
func GetUser(token string) (map[string]string, error) {
	return rdb.HGetAll(ctx, LOGIN_TOKEN+token).Result()
}
func RefreshUser(token string) error {
	return refresh(LOGIN_CODE_KEY+token, LOGIN_TOKEN_TTL)
}
func refresh(key string, ttl time.Duration) error {
	return rdb.Expire(ctx, key, ttl).Err()
}
func GetCacheShop(id string) (string, error) {
	return rdb.Get(ctx, CACHE_SHOP_KEY+id).Result()
}
func SetCacheShop(id string, shopJson string, ttl time.Duration) error {
	return rdb.Set(ctx, CACHE_SHOP_KEY+id, shopJson, ttl).Err()
}
func DelShop(id string) error {
	return rdb.Del(ctx, CACHE_SHOP_KEY+id).Err()
}
func SetCacheShopType(shopTypeList []string) error {
	return rdb.RPush(ctx, CACHE_SHOP_LIST, shopTypeList).Err()
}
func GetCacheShopType() ([]string, error) {
	return rdb.LRange(ctx, CACHE_SHOP_LIST, 0, -1).Result()
}
func TryLock(lock string) (bool, error) {
	return rdb.SetNX(ctx, lock, "1", LOCK_TTL).Result()
}
func Unlock(lock string) {
	rdb.Del(ctx, lock)
}
func SaveShop(id string, data interface{}) {
	rdb.Set(ctx, CACHE_SHOP_KEY+id, data, -1)
}

type RedisIdWorker struct{}

const BEGIN_TIMESTAMP int64 = 1661319393
const COUNT_BITS int = 32

func (r *RedisIdWorker) NextId(keyPrefix string) int64 {
	// 1.生成时间戳
	timeStamp := time.Now().Unix() - BEGIN_TIMESTAMP
	// 2.生成序列号
	date := time.Now().Format("2006:01:02")
	count, _ := rdb.Incr(ctx, "icr:"+keyPrefix+":"+date).Result()
	// 3.拼接返回
	return timeStamp<<int64(COUNT_BITS) | count
}
func SaveSeckillVoucher(voucherId string, stock string) error {
	return rdb.Set(ctx, SECKILL_STOCK_KEY+voucherId, stock, -1).Err()
}
func SIsMember(key string, member string) (bool, error) {
	//return rdb.SIsMember(ctx, key, member).Result()
	_, err := rdb.ZScore(ctx, key, member).Result()
	if err != nil {
		return false, err
	}
	return true, err
}

func Add2Set(key string, members interface{}) (int64, error) {
	return rdb.SAdd(ctx, key, members).Result()
}
func Add2ZSet(key string, member interface{}, score float64) (int64, error) {
	//return rdb.SAdd(ctx, key, value).Result()
	return rdb.ZAdd(ctx, key, redis.Z{Score: score, Member: member}).Result()
}

func RemoveFromSet(key string, members interface{}) (int64, error) {
	return rdb.SRem(ctx, key, members).Result()
}
func RemoveFromZSet(key string, members interface{}) (int64, error) {
	//return rdb.SRem(ctx, key, value).Result()
	return rdb.ZRem(ctx, key, members).Result()
}
func RangeZSet(key string, start, end int64) []int {
	z, _ := rdb.ZRangeWithScores(ctx, key, start, end).Result()
	idList := []int{}
	for k := range z {
		userId, _ := strconv.Atoi(z[k].Member.(string))
		idList = append(idList, userId)
	}
	return idList
}
