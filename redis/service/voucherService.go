package service

import (
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"
	"strconv"
)

type VoucherService struct{}

func (v *VoucherService) QueryVoucherOfShop(shopId int) dto.Result {
	var result dto.Result
	vouchers, err := dao.SelectVoucherOfShop(shopId)
	if err != nil {
		return result.Fail(err.Error())
	}
	return result.Ok(vouchers)
}
func (v *VoucherService) AddSeckillVoucher(voucher entity.Voucher) dto.Result {
	var result dto.Result
	id, err := dao.AddSeckillVoucher(voucher) // 数据库save
	// redis save
	if err != nil {
		return result.Fail("添加失败!")
	}
	voucherId_str := strconv.Itoa(voucher.Id)
	voucherStock_str := strconv.Itoa(voucher.Id)
	err = utils.SaveSeckillVoucher(voucherId_str, voucherStock_str)
	if err != nil {
		return result.Fail("添加失败!")
	}
	return result.Ok(id)
}
func (v *VoucherService) AddVoucher(voucher entity.Voucher) dto.Result {
	var result dto.Result
	id, err := dao.AddVoucher(voucher)
	if err != nil {
		return result.Fail("添加失败:" + err.Error())
	}
	return result.Ok(id)
}
