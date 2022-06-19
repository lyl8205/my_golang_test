package model

import (
	"crontab/abs"
	sanheerpflow "crontab/table/sanhe_erp_flow"
)

type alipaySetting struct {
	abs.Model
}

func NewAlipaySetting() *alipaySetting {
	return &alipaySetting{}
}

func (as *alipaySetting) GetAlipaySetting(appId string) (data sanheerpflow.ShAlipaySetting) {
	as.GetShanHeErpFlow().Table(sanheerpflow.ShAlipaySettingTN).Where(`appid = ?`, appId).Limit(1).Scan(&data)
	return data
}
