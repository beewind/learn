package dto

import (
	"redis-learn/entity"
	"strconv"
)

type UserDTO struct {
	Id       int    `json:"id" redis:"id"`
	NickName string `json:"nickName" redis:"nickName"`
	Icon     string `json:"icon" redis:"icon"`
}

func NewUserDTO(e entity.User) UserDTO {
	return UserDTO{
		Id:       e.Id,
		NickName: e.NickName,
		Icon:     e.Icon,
	}
}
func (u *UserDTO) ToStringSlice() (res []string) {
	res = append(res, "id")
	res = append(res, strconv.Itoa(u.Id))

	res = append(res, "nickName")
	res = append(res, u.NickName)

	res = append(res, "icon")
	res = append(res, u.Icon)
	return

}
