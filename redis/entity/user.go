package entity

import "redis-learn/utils"

type User struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	Phone      string     `json:"phone"`
	Password   string     `json:"password"`
	NickName   string     `json:"nickname"`
	Icon       string     `json:"icon"`
	CreateTime utils.Time `json:"createTime"`
	UpdateTime utils.Time `json:"updateTime"`
}
