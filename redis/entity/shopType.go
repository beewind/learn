package entity

import "redis-learn/utils"

type ShopType struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Sort       int    `json:"sort"`
	CreateTime utils.Time
	UpdateTime utils.Time
}
