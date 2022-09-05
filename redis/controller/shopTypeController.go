package controller

import (
	"net/http"
	"redis-learn/service"

	"github.com/gin-gonic/gin"
)

func QueryShopList(ctx *gin.Context) {
	result := service.QueryShopList()
	ctx.JSON(http.StatusOK, result)
}
