package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var voucherService service.VoucherService

func AddSeckillVoucher(ctx *gin.Context) {
	var voucher entity.Voucher
	var result dto.Result
	if ctx.BindJSON(&voucher) != nil {
		ctx.JSON(http.StatusOK, result.Fail("数据读取失败"))
	} else {
		result = voucherService.AddSeckillVoucher(voucher)
		ctx.JSON(http.StatusOK, result)
	}
}
func AddVoucher(ctx *gin.Context) {
	var voucher entity.Voucher
	var result dto.Result
	if ctx.BindJSON(&voucher) != nil {
		ctx.JSON(http.StatusOK, result.Fail("数据读取失败!"))
	} else {
		result = voucherService.AddVoucher(voucher)
		ctx.JSON(http.StatusOK, result)
	}

}
func QueryVoucherOfShop(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shopId"))
	result := voucherService.QueryVoucherOfShop(shopId)
	ctx.JSON(http.StatusOK, result)
}
