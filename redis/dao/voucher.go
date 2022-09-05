package dao

import (
	"fmt"
	"redis-learn/entity"

	"gorm.io/gorm"
)

const TbVoucher string = "tb_voucher"
const TbSeckillVoucher string = "tb_seckill_voucher"

func SelectVoucherOfShop(shopId int) ([]entity.Voucher, error) {
	res := []entity.Voucher{}
	err := db.Table(TbVoucher).Where("shop_id = ?", shopId).Find(&res).Error
	return res, err
}
func SelectVoucherSeckillId(voucherId int) (entity.SeckillVoucher, error) {
	var voucherSeckill entity.SeckillVoucher
	err := db.Table(TbSeckillVoucher).Where("voucher_id = ?", voucherId).First(&voucherSeckill).Error
	return voucherSeckill, err
}
func SelectVoucherId(id int) (entity.Voucher, error) {
	var voucher entity.Voucher
	err := db.Table(TbVoucher).Where("id = ?", id).First(&voucher).Error
	return voucher, err
}
func AddVoucher(voucher entity.Voucher) (int, error) {
	err := db.Table(TbVoucher).Create(&voucher).Error
	return voucher.Id, err
}
func AddSeckillVoucher(voucher entity.Voucher) (int, error) {
	err := db.Table(TbVoucher).Create(&voucher).Error
	if err != nil {
		return 0, err
	}
	var seckillVoucher entity.SeckillVoucher

	seckillVoucher.VoucherId = voucher.Id
	seckillVoucher.Stock = voucher.Stock
	seckillVoucher.BeginTime = voucher.BeginTime
	seckillVoucher.EndTime = voucher.EndTime

	err = db.Table(TbSeckillVoucher).Omit("create_time", "update_time").Create(&seckillVoucher).Error
	if err != nil {
		return 0, err
	}
	return seckillVoucher.VoucherId, err

}

// 库存-1
func SubSeckillVoucher(voucherId int) error {
	result := db.Table(TbSeckillVoucher).Where("voucher_id = ? and stock > 0", voucherId).Update("stock", gorm.Expr("stock - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("优惠券已售罄")
	}
	return nil
}
