package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var voucherOrder service.VoucherOrderService

func SeckillVoucher(ctx *gin.Context) {
	voucherId := ctx.Param("id")
	voucherId_int, _ := strconv.Atoi(voucherId)
	user, _ := ctx.Get("user")
	userId := user.(dto.UserDTO).Id
	ctx.JSON(http.StatusOK, voucherOrder.SeckillVoucher(voucherId_int, userId))
}
