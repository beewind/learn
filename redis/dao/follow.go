package dao

import "redis-learn/entity"

const TbFollow = "tb_follow"

func CreateFollow(followUserId, userId int) error {
	var follow entity.Follow
	follow.FollowUserId = followUserId
	follow.UserId = userId
	return db.Table(TbFollow).Select("follow_user_id", "user_id").Create(&follow).Error
}
func DelFollow(followUserId, userId int) error {
	return db.Table(TbFollow).Where("follow_user_id = ? and user_id = ?", followUserId, userId).Delete(&entity.Follow{}).Error
}
func SelectFollow(followUserId, userId int) (entity.Follow, error) {
	var follow entity.Follow
	err := db.Table(TbFollow).Where("follow_user_id = ? and user_id = ?", followUserId, userId).First(&follow).Error
	return follow, err
}
