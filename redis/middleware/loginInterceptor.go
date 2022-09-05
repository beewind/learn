package middleware

import (
	"redis-learn/dto"
	"redis-learn/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Filter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{
				"msg": "未登录",
			})
		} else {
			ctx.Next()
		}

	}
}
func RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			//1. 获取session
			session := sessions.Default(ctx)

			//2. 获取session中的用户
			user := session.Get("user")
			var result dto.Result
			//3.判断用户是否存在
			if user == nil {
				result.Fail("未授权访问!")
				ctx.JSON(http.StatusUnauthorized, result)
				ctx.Abort()
			} else {
				ctx.Set("user", user.(dto.UserDTO))
				ctx.Next()

			}
		*/

		// 1.获取token
		// fmt.Println(ctx.Request.Header)
		token, ok := ctx.Request.Header["Authorization"] // 注意前后端的字段,这里前端发的是authorization ,但是不知道为什么自动变化了
		if ok {
			// 2.基于token获取user
			//fmt.Println(token)
			userMap, err := utils.GetUser(token[0])
			// 3.user的非空判断
			if err == nil && len(userMap) != 0 {
				// 4.user类型转换
				var user dto.UserDTO
				user.Id, _ = strconv.Atoi(userMap["id"])
				user.NickName = userMap["nickName"]
				user.Icon = userMap["icon"]

				// 5.储存user
				ctx.Set("user", user)

				// 6.刷新token
				utils.RefreshUser(token[0])
			}

		}

		// 7.放行
		ctx.Next()
	}
}
