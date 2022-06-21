package config

type billKey struct {
	TemplateBillNumKey, TemplateBillKey string
	CommonKey                           map[string]billCommonKey
}

type billCommonKey struct {
	KeyNum, AppletAppId, LifeNumberAppId, MiNiTemplateId, LifeTemplateId, MiNiPageUrl, LifePageUrl,LifeProductName,MiNiProductName string
}

type redisKey struct {
	BillKey billKey
}

var Key = redisKey{BillKey: billKey{
	TemplateBillNumKey: `crontab:tempbillnum:hash:%s:%s:`,
	TemplateBillKey:    `crontab:%stempbill:string:%s`,
	CommonKey: map[string]billCommonKey{
		`gd`: {
			KeyNum:          `gd`,
			AppletAppId:     `2019031863556379`,
			LifeNumberAppId: `2013112700002210`,
			//MiNiTemplateId:  `f59b2d5ad69b4a9eadc86399ce321988`,
			MiNiTemplateId:  `d8fd871671c34d038262e742c81a9121`,
			//LifeTemplateId:  `7d4df44b83be4fc1ba8c82cc90259270`,
			LifeTemplateId:  `f697b31685e94030a654d4f52cc3448f`,
			MiNiPageUrl:     `alipays://platformapi/startapp?appId=2019031863556379&page=pages/index/index&query=source%3D1%26channel%3Dgd_xcx%26column%3Dmbxx%26pubsrc%3Dqylqtx_mbxx`,
			LifePageUrl:     `alipays://platformapi/startapp?appId=2019031863556379&page=pages/index/index&query=source%3D1%26channel%3Dgd_xcx%26column%3Dmbxx%26pubsrc%3Dqylqtx_mbxx`,
			LifeProductName:     `芒果TV月会员`,
			MiNiProductName:     `0元领芒果TV月会员`,
		},
	},
}}
