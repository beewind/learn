package service

import (
	"encoding/json"
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"

	"github.com/go-redis/redis/v9"
)

func QueryShopList() dto.Result {
	var result dto.Result
	// 1.查询cache
	cache, err := utils.GetCacheShopType()

	if err != nil && err != redis.Nil {
		return result.Fail("查询错误:" + err.Error())

	}
	// 2.存在cache
	if len(cache) != 0 {
		shopTypeList := []entity.ShopType{}
		for k := range cache {
			var temp entity.ShopType
			json.Unmarshal([]byte(cache[k]), &temp)
			shopTypeList = append(shopTypeList, temp)
		}
		return result.Ok(shopTypeList)
	}

	// 3.不存在cache
	// 4.查询数据库
	shopTypeList, err := dao.SelectShopList()
	if err != nil {
		return result.Fail(err.Error())
	}

	// 5.缓存
	shopTypeListJson := []string{}
	for k := range shopTypeList {
		temp, err := json.Marshal(shopTypeList[k])
		if err != nil {
			return result.Fail("序列化错误:" + err.Error())
		}
		shopTypeListJson = append(shopTypeListJson, string(temp))
		//fmt.Println(string(temp))
	}

	utils.SetCacheShopType(shopTypeListJson)

	// 6.返回
	return result.Ok(shopTypeList)
}
