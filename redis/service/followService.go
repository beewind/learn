package service

import (
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/utils"
	"strconv"

	"gorm.io/gorm"
)

type FollowService struct{}

func (f *FollowService) Follow(followUserId, UserId int, isFollow bool) dto.Result {
	var result dto.Result
	// 1.判断关注还是取关
	userId_str := strconv.Itoa(UserId)
	key := "follows" + userId_str
	if isFollow {
		// 2.关注
		err := dao.CreateFollow(followUserId, UserId)
		if err == nil {
			utils.Add2Set(key, followUserId)
		}
	} else {
		// 3.取关
		err := dao.DelFollow(followUserId, UserId)
		if err == nil {
			utils.RemoveFromSet(key, followUserId)
		}
	}
	return result.Ok(nil)
}
func (f *FollowService) IsFollow(followUserId, userId int) dto.Result {
	var result dto.Result
	_, err := dao.SelectFollow(followUserId, userId)
	if err == gorm.ErrRecordNotFound {
		return result.Ok(false)
	}
	return result.Ok(true)
}
