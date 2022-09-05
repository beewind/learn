package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/service"

	"github.com/gin-gonic/gin"
)

var userServer service.UserServer

func SendCode(ctx *gin.Context) {
	// 获取url中的phone参数
	result := userServer.SendCode(ctx)
	ctx.JSON(http.StatusOK, result)
}

func Login(ctx *gin.Context) {
	result := userServer.Login(ctx)
	ctx.JSON(http.StatusOK, result)
}
func Me(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// user := session.Get("user")
	// var result dto.Result
	// if user == nil {
	// 	result.Fail("错误请求!")
	// 	ctx.JSON(http.StatusBadRequest, result)
	// 	return
	// }
	// result.Ok(user)
	// ctx.JSON(http.StatusOK, result)
	user, ok := ctx.Get("user")
	var result dto.Result

	if !ok {
		result.Fail("错误请求!")
		ctx.JSON(http.StatusBadRequest, result)
	}

	result.Ok(user)
	ctx.JSON(http.StatusOK, result)
}
