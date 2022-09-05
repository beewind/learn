package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//:id/:isFollow
var followServer service.FollowService

func FollowController(ctx *gin.Context) {
	userDto, _ := ctx.Get("user")
	userId := userDto.(dto.UserDTO).Id
	followUserId_str := ctx.Param("id")
	isFollow_str := ctx.Param("isFollow")
	followUserId, _ := strconv.Atoi(followUserId_str)
	isFollow, _ := strconv.ParseBool(isFollow_str)
	result := followServer.Follow(followUserId, userId, isFollow)
	ctx.JSON(http.StatusOK, result)
}

func IsFollow(ctx *gin.Context) {
	// /or/not/:id
	userDto, _ := ctx.Get("user")
	userId := userDto.(dto.UserDTO).Id
	followUserId_str := ctx.Param("id")
	followUserId, _ := strconv.Atoi(followUserId_str)
	result := followServer.IsFollow(followUserId, userId)
	ctx.JSON(http.StatusOK, result)
}
