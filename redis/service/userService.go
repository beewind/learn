package service

import (
	"fmt"
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserServer struct{}

func (u *UserServer) SendCode(ctx *gin.Context) (result dto.Result) {
	phone := ctx.Query("phone")
	// 1.校验手机号
	if utils.IsPhoneInvalid(phone) {
		// 2.不符合,返回错误
		result.Fail("手机号格式错误!")
		return
	}

	// 3.符合,生成验证码
	code := utils.RandNumString(6)

	// 4.保存验证码

	/* session
	session := sessions.Default(ctx)
	session.Set("code", code)
	session.Save()
	*/

	// redis
	utils.SetCode(phone, code)

	// 模拟日志记录验证码:
	fmt.Println(phone, ":", code)

	// 5.返回验证码
	result.Ok(nil)
	return
}
func (u *UserServer) Login(ctx *gin.Context) (result dto.Result) {
	var loginform dto.LoginFormDto
	if ctx.ShouldBindJSON(&loginform) != nil {
		result.Fail("无法获取提交数据!")
		return
	}
	// 1.校验手机号
	if utils.IsPhoneInvalid(loginform.Phone) {
		result.Fail("手机号格式错误!")
		return
	}
	// 2.校验验证码
	//session := sessions.Default(ctx)
	//code := session.Get("code")
	code, err := utils.GetCode(loginform.Phone)
	if err != nil {
		result.Fail("验证失败!")
		return
	}
	if code != loginform.Code {
		// 3.不一致报错
		result.Fail("验证码错误!")
		return
	}

	// 4.一致依据手机号查询 select * from tbb_user where phone = ?
	user, err := dao.SelectOne(loginform.Phone)

	// 5.判断用户存在
	if err != nil {
		// 6.用户不存在
		user = dao.CreateUserWithPhone(loginform.Phone)
	}
	userDTO := dto.NewUserDTO(user)

	// 7.1.随机token
	newUuid := uuid.New().String()

	// 7.2.保存user到redis
	utils.SetUser(newUuid, userDTO.ToStringSlice())
	/*
		session.Set("user", userDTO)
		session.Save()
	*/

	result.Ok(newUuid)
	return
}
func (u *UserServer) SelectById(id int) entity.User {
	user, _ := dao.SelectUserById(id)
	return user
}
func (u *UserServer) SelectIdIn(id ...int) []entity.User {
	userSlice := dao.SelectUserIn(id)
	return userSlice
}
