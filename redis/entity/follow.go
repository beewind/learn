package entity

import "redis-learn/utils"

type Follow struct {
	Id           int        `json:"id"`
	UserId       int        `json:"userId"`
	FollowUserId int        `json:"followUserId"`
	CreateTime   utils.Time `json:"createTime"`
}
