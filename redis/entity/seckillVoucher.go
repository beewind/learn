package entity

import "redis-learn/utils"

type SeckillVoucher struct {
	VoucherId  int        `json:"voucherId"`
	Stock      int        `json:"stock"`
	CreateTime utils.Time `json:"createTime"`
	BeginTime  utils.Time `json:"beginTime"`
	EndTime    utils.Time `json:"endTime"`
	UpdateTime utils.Time `json:"updateTime"`
}
