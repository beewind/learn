package controller

import (
	"net/http"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var blogServer service.BlogServer

func QueryHotBlog(ctx *gin.Context) {
	current := ctx.Query("current")
	user, ok := ctx.Get("user")
	var userId int
	if ok {
		userId = user.(dto.UserDTO).Id
	} else {
		userId = 0
	}
	result := blogServer.QueryHotBlog(current, userId)
	ctx.JSON(http.StatusOK, result)
}

func QueryBlogById(ctx *gin.Context) {
	blogId := ctx.Param("id")
	user, ok := ctx.Get("user")
	var userId int
	if ok {
		userId = user.(dto.UserDTO).Id
	}
	result := blogServer.GetById(blogId, userId)
	ctx.JSON(http.StatusOK, result)

}
func SaveBlog(ctx *gin.Context) {
	var blog entity.Blog
	var result dto.Result
	user, _ := ctx.Get("user")
	if ctx.BindJSON(&blog) != nil {
		result.Fail("无法获取提交blog")
	} else {
		blog.UserId = user.(dto.UserDTO).Id
		result = blogServer.SaveBlog(blog)

	}
	ctx.JSON(http.StatusOK, result)
}
func BlogLike(ctx *gin.Context) {
	var result dto.Result
	blogId, _ := strconv.Atoi(ctx.Param("id"))
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusOK, result.Ok(nil))
	}
	userId := user.(dto.UserDTO).Id
	result = blogServer.LikeBlog(blogId, userId)
	ctx.JSON(http.StatusOK, result)
}
func BlogLikes(ctx *gin.Context) {
	blogId := ctx.Param("id")
	blogId_int, _ := strconv.Atoi(blogId)
	result := blogServer.QueryBlogLikes(blogId_int)
	ctx.JSON(http.StatusOK, result)
}
