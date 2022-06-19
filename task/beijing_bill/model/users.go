package model

import (
	"crontab/abs"
	"crontab/table/sanhe_gdmobile"
)

type users struct {
	abs.Model
}

func NewUsers() *users {
	return &users{}
}

func (ug *users) GetUsersById(id, limit int) (data []sanhe_gdmobile.TUsersBj) {
	ug.GetSanHeGdMobile().Table(sanhe_gdmobile.TUsersBjTN).Where("id > ?", id).Order(`id asc`).Limit(limit).Scan(&data)
	return data

}
