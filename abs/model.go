package abs

import (
	"go_test/config"

	mysql2 "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/mysql"
	"gorm.io/gorm"
)

type Model struct {
	mysql2.Model
}

var sanHeGdMobile, shanHeErpFlow mysql2.DbCollector

func (m *Model) GetSanHeGdMobile() *gorm.DB {
	return m.NewClient(&sanHeGdMobile, config.Mysql.SanHeGdMobile)
}

func (m *Model) GetShanHeErpFlow() *gorm.DB {
	return m.NewClient(&shanHeErpFlow, config.Mysql.ShanHeErpFlow)
}
