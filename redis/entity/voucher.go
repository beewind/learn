package entity

import "redis-learn/utils"

type Voucher struct {
	Id          int        `json:"id"`
	ShopId      string     `json:"shopId"`
	Title       string     `json:"title"`
	SubTitle    string     `json:"subTitle"`
	Rules       string     `json:"rules"`
	PayValue    string     `json:"payValue"`
	ActualValue int        `json:"actualValue"`
	Type        int        `json:"type"`
	Status      int        `json:"status"`
	Stock       int        `json:"stock"`
	BeginTime   utils.Time `json:"beginTime"`
	EndTime     utils.Time `json:"endTime"`
	CreateTime  utils.Time `json:"createTime"`
	UpdateTime  utils.Time `json:"updateTime"`
}
