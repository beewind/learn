package service

import (
	"encoding/json"
	"log"
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type ShopSever struct{}

func (s *ShopSever) QueryById(id string) dto.Result {
	// 缓存穿透
	// return queryByIdPassThrough(id)

	// 互斥锁解决缓存击穿
	return queryByIdMutex(id)

	// 逻辑过期时间解决缓存击穿
	// return queryWithLogicalExpire(id)
}

// 互斥锁解决缓冲击穿
func queryByIdMutex(id string) dto.Result {

	var result dto.Result
	// 1.从redis查询商铺缓存
	cache, err := utils.GetCacheShop(id)

	if err != nil && err != redis.Nil {
		result.Fail(err.Error())
		return result
	}
	if err == nil {
		// 2.判断是否零值
		if cache == "" {
			return result.Fail("店铺不存在")
		} else {
			// 3.存在,不为空值返回
			var shop entity.Shop
			json.Unmarshal([]byte(cache), &shop)
			result.Ok(shop)
			return result
		}

	}

	// 4.实现缓存重建,更据id查询数据库
	lockKey := "lock:shop:" + id

	// 4.1 获取互斥锁
	isLock, _ := utils.TryLock(lockKey)
	defer utils.Unlock(lockKey)
	// 4.2 判断是否获取
	if !isLock {
		time.Sleep(50 * time.Millisecond)
		return queryByIdMutex(id)
	}
	// 4.3 获取成功
	// 4.3.1 查看缓存
	// cache, err = utils.GetCacheShop(id)
	// if err == nil && cache == "" {
	// 	return result.Fail("店铺不存在")
	// }
	// 4.3.2 查询数据库
	shop, err := dao.SelectShopById(id)

	// 5.不存在,返回错误信息
	if err != nil {
		result.Fail("店铺不存在!")
		utils.SetCacheShop(id, "", utils.CACHE_NIL_TTL)
		return result
	}

	shopJson, err := json.Marshal(shop)
	if err != nil {
		log.Println(err)
		result.Fail("解析错误!" + err.Error())
		return result
	}
	cache = string(shopJson)
	// 6.存在,写入redis
	err = utils.SetCacheShop(id, cache, utils.CACHE_SHOP_TTL)
	if err != nil {
		log.Println(err)
		result.Fail("错误发生!" + err.Error())
		return result
	}
	// 7.返回
	return result.Ok(shop)
}

// 逻辑过期
func queryWithLogicalExpire(id string) dto.Result {

	var result dto.Result
	// 1.从redis查询商铺缓存
	cache, err := utils.GetCacheShop(id)
	var redisData utils.RedisData
	// 2.判断缓存命中
	if err != nil {
		// 3.未命中
		return result.Fail("获取失败")
	}
	json.Unmarshal([]byte(cache), &redisData)
	// 4.是否过期
	//  未过期
	if redisData.ExpireTime.After(time.Now()) {
		return result.Ok(redisData)
	}
	// 过期

	// 5.获取锁
	lockKey := "lock:shop:" + id
	isLock, _ := utils.TryLock(lockKey)

	if !isLock {
		return result.Ok(redisData)
	}
	go func() {
		saveShop2Redis(id, 200*time.Millisecond)
		utils.Unlock(lockKey)
	}()
	return result.Ok(redisData)

}

// 解决缓冲穿透
func queryByIdPassThrough(id string) dto.Result {

	var result dto.Result
	// 1.从redis查询商铺缓存
	cache, err := utils.GetCacheShop(id)

	if err != nil && err != redis.Nil {
		result.Fail(err.Error())
		return result
	}
	if err == nil {
		// 2.判断是否零值
		if cache == "" {
			return result.Fail("店铺不存在")
		} else {
			// 3.存在,不为空值返回
			var shop entity.Shop
			json.Unmarshal([]byte(cache), &shop)
			result.Ok(shop)
			return result
		}

	}

	// 4.实现缓存重建,更据id查询数据库
	shop, err := dao.SelectShopById(id)

	// 5.不存在,返回错误信息
	if err != nil {
		result.Fail("店铺不存在!" + err.Error())
		utils.SetCacheShop(id, "", utils.CACHE_NIL_TTL)
		return result
	}

	shopJson, err := json.Marshal(shop)
	if err != nil {
		log.Println(err)
		result.Fail("解析错误!" + err.Error())
		return result
	}
	cache = string(shopJson)
	// 6.存在,写入redis
	err = utils.SetCacheShop(id, cache, utils.CACHE_SHOP_TTL)
	if err != nil {
		log.Println(err)
		result.Fail("错误发生!" + err.Error())
		return result
	}
	// 7.返回
	return result.Ok(shop)
}

func (s *ShopSever) Update(shop entity.Shop) dto.Result {
	var result dto.Result
	// 1.更新数据库
	if shop.Id == 0 {
		return result.Fail("商店id不能为空")
	}
	dao.UpdateShopById(shop)
	// 2.删除缓存
	id := strconv.Itoa(int(shop.Id)) //可能溢出
	utils.DelShop(id)
	return result.Ok(nil)
}

// func (s *ShopSever) QueryShopByType(typeId, current, x, y) {

// }

func (s *ShopSever) Save(shop entity.Shop) dto.Result {
	err := dao.SaveShop(shop)
	var result dto.Result
	if err != nil {
		return result.Fail("添加失败!")
	} else {
		return result.Ok(nil)
	}
}
func saveShop2Redis(id string, ttl time.Duration) {
	// 1.查询店铺数据
	shop, _ := dao.SelectShopById(id)

	// 2.逻辑过期时间
	var redisData utils.RedisData
	redisData.Data = shop
	redisData.ExpireTime = time.Now().Add(ttl)

	// 3.写入redis
	b, _ := json.Marshal(redisData)
	utils.SaveShop(id, string(b))
}
