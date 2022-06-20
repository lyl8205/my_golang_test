package model

import (
	"go_test/abs"
	"go_test/table/sanhe_gdmobile"
)

type usersGd struct {
	abs.Model
}

func NewUsersGd() *usersGd {
	return &usersGd{}
}

func (ug *usersGd) GetUsersGdById(id, limit int) (data []sanhe_gdmobile.TUsersGd) {
	ug.GetSanHeGdMobile().Table(sanhe_gdmobile.TUsersGdTN).Where("id > ?", id).Order(`id asc`).Limit(limit).Scan(&data)
	return data

}
