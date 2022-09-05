package service

import (
	"log"
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"
	"strconv"
	"time"
)

type BlogServer struct{}

var userServer UserServer

func (b *BlogServer) QueryHotBlog(current string, userId int) dto.Result {
	current_int, _ := strconv.Atoi(current)
	page := dao.QueryHotBlog(current_int, utils.MAX_PAGE_SIZE)
	for k := range page {
		b.queryBlogUser(&page[k])
		page[k].IsLike = b.isBlogLiked(page[k], userId)
	}
	var result dto.Result
	return result.Ok(page)
}
func (b *BlogServer) GetById(id string, userId int) dto.Result {
	var result dto.Result
	id_int, _ := strconv.Atoi(id)
	// 1.查询blog
	blog, err := dao.SelectBlogById(id_int)
	if err != nil {
		return result.Fail("笔记不存在")
	}
	// 2.查询blog的作者信息
	b.queryBlogUser(&blog)

	// 3.是否被点赞
	if userId != 0 {
		blog.IsLike = b.isBlogLiked(blog, userId)
	}

	return result.Ok(blog)
}
func (b *BlogServer) isBlogLiked(blog entity.Blog, userId int) bool {
	// 判断是否点过赞
	key := "blog:liked:" + strconv.Itoa(blog.Id)
	userId_str := strconv.Itoa(userId)
	isLike, err := utils.SIsMember(key, userId_str)
	if err != nil {
		log.Println(err)
	}
	return isLike
}
func (b *BlogServer) SaveBlog(blog entity.Blog) dto.Result {
	var result dto.Result
	// 1.保存探店笔记
	err := dao.SaveBlog(&blog)
	if err != nil {
		return result.Fail("添加失败!")
	}
	// 2.查询笔记作者的所有粉丝 select * from tb_follow where follow_user_id = ?

	// 3.推送笔记给所有的粉丝
	// 3.1 获取粉丝id
	// 3.2 推送

	// 4.返回id
	return result.Ok(blog.Id)
}
func (b *BlogServer) queryBlogUser(blog *entity.Blog) {
	upId := blog.UserId
	user := userServer.SelectById(upId)
	blog.Icon = user.Icon
	blog.Name = user.NickName

}
func (b *BlogServer) LikeBlog(blogId, userId int) dto.Result {
	var result dto.Result
	// 1.获取登录用户

	// 2.判断当前用户是否点赞 //sismember
	key := "blog:liked:" + strconv.Itoa(blogId)
	userId_str := strconv.Itoa(userId)
	isLike, err := utils.SIsMember(key, userId_str)
	if err != nil {
		log.Println(err)
	}

	// 3.如果未点赞
	if !isLike {
		// 3.1 数据库点赞数加一
		dao.AddBlogLike(blogId)
		// 3.2 保存用户id到set
		utils.Add2ZSet(key, userId, float64(time.Now().Unix()))
	} else {
		// 4.如果用户已点赞
		// 4.1 数据库点赞数减一
		dao.SubBlogLike(blogId)
		// 4.2 从set删除用户id
		utils.RemoveFromZSet(key, userId)
	}
	return result.Ok(nil)
}

func (b *BlogServer) QueryBlogLikes(blogId int) dto.Result {
	key := "blog:liked:" + strconv.Itoa(blogId)
	likedUserIds := utils.RangeZSet(key, 0, 4)
	var userDtoList = []dto.UserDTO{}
	userList := dao.SelectUserIn(likedUserIds)
	/* for userId := range likedUserIds {
		user := userServer.SelectById(userId)
		userDto := dto.NewUserDTO(user)
		userDtoList = append(userDtoList, userDto)
	} */
	for k := range userList {
		userdto := dto.NewUserDTO(userList[k])
		userDtoList = append(userDtoList, userdto)
	}
	var result dto.Result
	return result.Ok(userDtoList)
}
