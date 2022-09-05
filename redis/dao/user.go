package dao

import (
	"fmt"
	"redis-learn/entity"
	"redis-learn/utils"
	"strconv"
	"strings"
)

const TbUser string = "tb_user"

func SelectOne(phone string) (entity.User, error) {
	var user entity.User
	//
	err := db.Table(TbUser).Where("phone = ?", phone).First(&user).Error
	return user, err
}
func SelectUserById(id int) (entity.User, error) {
	var user entity.User
	err := db.Table(TbUser).Where("id = ?", id).First(&user).Error
	return user, err
}
func FirstOrCreate(phone string) (entity.User, error) {
	var user entity.User
	err := db.Table(TbUser).Where("phone = ?", phone).FirstOrCreate(&user).Error
	return user, err
}
func CreateUserWithPhone(phone string) entity.User {
	var user entity.User
	user.Phone = phone
	user.NickName = utils.USER_NICK_NAME_PREFIX + utils.RandAllString(10)
	db.Table(TbUser).Select("phone", "nick_name").Create(&user)
	fmt.Println(user)
	return user
}

func SelectUserIn(id []int) []entity.User {
	var userSlice []entity.User
	idstr := []string{}
	idstr = append(idstr, "FIELD(id")
	for _, v := range id {
		idstr = append(idstr, strconv.Itoa(v))
	}
	idString := strings.Join(idstr, ",")
	idString += ")"
	//db.Exec("select * from ? where id in (?) order by FIELD(id,?)", TbUser, id).Scan(&userSlice)
	db.Table(TbUser).Where("id in (?)", id).Order(idString).Find(&userSlice)
	return userSlice
}
