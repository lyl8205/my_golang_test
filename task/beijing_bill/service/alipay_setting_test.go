package service

import "testing"

func TestNewAlipaySetting_GetAlipaySetting(t *testing.T) {
	appId := `2019031863556379`
	alipaySettingData := NewAlipaySetting().GetAlipaySetting(appId)
	t.Log("alipaySettingData==", alipaySettingData)
}
