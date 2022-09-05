package entity

import "redis-learn/utils"

type VoucherOrder struct {
	Id        int `json:"id"`
	UserId    int `json:"userId"`
	VoucherId int `json:"voucherId"`
	// 1:余额支付 2:支付宝 3:微信
	PayType int `json:"payType"`
	// 订单状态:1:未支付 2:已支付 3:已核销 4:已取消 5:退款中 6:已退款
	Status int `json:"status"`
	// 下单时间
	CreateTime utils.Time `json:"createTime"`
	// 支付时间
	PayTime utils.Time `json:"payTime"`
	// 核销时间
	UseTime utils.Time `json:"useTime"`
	// 退款时间
	RefundTime utils.Time `json:"refundTime"`
	// 更新时间
	UpdateTime utils.Time `json:"updateTime"`
}
