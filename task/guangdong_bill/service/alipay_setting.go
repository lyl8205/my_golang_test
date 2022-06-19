package service

import (
	"go_test/abs"
	sanheerpflow "go_test/table/sanhe_erp_flow"
	"go_test/task/guangdong_bill/model"
	"encoding/json"
	"fmt"
	"time"
)

var AlipaySetting = `crontab:alipaysetting:%s`

type alipaySetting struct {
	abs.Redis
}

func NewAlipaySetting() *alipaySetting {
	return &alipaySetting{}
}

func (as *alipaySetting) GetAlipaySetting(appId string) (data sanheerpflow.ShAlipaySetting) {

	key := fmt.Sprintf(AlipaySetting, appId)
	redisClient := as.GetCache()
	aliSettingString := redisClient.Get(key).String()

	if aliSettingString != `` {
		if err := json.Unmarshal([]byte(aliSettingString), &data); err == nil && data.Id > 0 {
			return data
		}
	}
	aliSettingData := model.NewAlipaySetting().GetAlipaySetting(appId)
	if aliSettingData.Id > 0 {
		if aliByte, err := json.Marshal(aliSettingData); err == nil {
			_ = redisClient.Set(key, string(aliByte), 86400*time.Second).Val() //缓存一天
		}
	}
	return aliSettingData
}
