package entity

import "redis-learn/utils"

type Shop struct {
	//SerialVersionUid int
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	TypeId     int64      `json:"typeId"`
	Images     string     `json:"images"` // 多个图片用','隔开
	Area       string     `json:"area"`
	Address    string     `json:"address"`
	X          float64    `json:"x"`
	Y          float64    `json:"y"` //经纬度
	AvgPrice   int64      `json:"avgPrice"`
	Sold       int        `json:"sold"`
	Comments   int        `json:"comments"` // 评论数量
	Score      int        `json:"score"`    // 评分
	OpenHours  string     `json:"openHours"`
	CreateTime utils.Time `json:"createTime"`
	UpdateTime utils.Time `json:"updateTime"`
	Distance   float64    `json:"omit"`
}
