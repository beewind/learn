package entity

import "redis-learn/utils"

type Blog struct {
	Id         int        `json:"id"`
	ShopId     int        `json:"shopId"`
	UserId     int        `json:"userId"`
	Icon       string     `json:"icon" gorm:"-"`
	Name       string     `json:"name" gorm:"-"`
	IsLike     bool       `json:"isLike" gorm:"-"`
	Title      string     `json:"title"`
	Images     string     `json:"images"`
	Content    string     `json:"content"`
	Liked      int        `json:"liked"`
	Comments   int        `json:"comments"`
	CreateTime utils.Time `json:"createTime"`
	UpdateTime utils.Time `json:"updateTime"`
}
