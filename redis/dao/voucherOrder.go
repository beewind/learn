package dao

import (
	"redis-learn/entity"
)

const TbVoucherOrder = "tb_voucher_order"

func CreateVoucherOrder(voucherOrder entity.VoucherOrder) error {
	return db.Table(TbVoucherOrder).Omit("create_time", "update_time").Create(voucherOrder).Error
}
func SelectWithUserIdAndVoucherId(userId, voucherId int) int {
	var sameRows int64
	db.Table(TbVoucherOrder).Where("user_id = ? and voucher_id = ?", userId, voucherId).Count(&sameRows)
	return int(sameRows)
}
