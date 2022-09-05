package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/service"
	"redis-learn/utils"

	"github.com/gin-gonic/gin"
)

var shopSever service.ShopSever

func QueryShopById(ctx *gin.Context) {
	var result dto.Result
	var shopSever service.ShopSever
	idStr := ctx.Param("id")
	//idInt, err := strconv.Atoi(idStr)

	if idStr == "" || utils.IsNumInvalid(idStr) {
		result.Fail("参数错误!")
		ctx.JSON(http.StatusBadRequest, result)
	} else {
		result := shopSever.QueryById(idStr)
		ctx.JSON(http.StatusOK, result)
	}

}
func SaveShop(ctx *gin.Context) {
	var shop entity.Shop
	var result dto.Result
	if ctx.BindJSON(&shop) != nil {
		result.Fail("获取参数失败!")
	} else {
		result = shopSever.Save(shop)
	}
	ctx.JSON(http.StatusOK, result)
}
func UpdateShop(ctx *gin.Context) {
	var shop entity.Shop
	var result dto.Result
	if ctx.BindJSON(&shop) != nil {
		result.Fail("获取参数失败!")
	} else {
		result = shopSever.Update(shop)
	}
	ctx.JSON(http.StatusOK, result)
}
