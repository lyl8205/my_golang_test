package config

type billKey struct {
	TemplateBillNumKey, TemplateBillKey string
	CommonKey                           map[string]billCommonKey
}

type billCommonKey struct {
	KeyNum, AppletAppId, LifeNumberAppId, MiNiTemplateId, LifeTemplateId, MiNiPageUrl, LifePageUrl string
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
			MiNiTemplateId:  `f59b2d5ad69b4a9eadc86399ce321988`,
			LifeTemplateId:  `7d4df44b83be4fc1ba8c82cc90259270`,
			MiNiPageUrl:     `pages/index/index&query=s%3D1`,
			LifePageUrl:     `alipays://platformapi/startapp?appId=2019060665494484`,
		},
		`bj`: {
			KeyNum:          `bj`,
			AppletAppId:     `2018122962713328`,
			LifeNumberAppId: `2013103100001891`,
			MiNiTemplateId:  `97449fafc2b946e593ece1ec5bddc57c`,
			LifeTemplateId:  `54712ee384c04400a3da993a23b91ba4`,
			MiNiPageUrl:     `pages/query/bill/bill&query=column%3Dxcx_zdfc%26channel%3Dbj_xcx%26source%3D1%26mobile=`,
			LifePageUrl: `alipays://platformapi/startapp?appId=2018122962713328&page=pages/query/bill/bill&query=co
			lumn%3Dxcx_zdfc%26channel%3Dbj_xcx%26source%3D1%26mobile=`,
		},
		`sd`: {
			KeyNum:          `sd`,
			AppletAppId:     `2019062865730087`,
			LifeNumberAppId: `2014040800004702`,
			MiNiTemplateId:  `3c05f3b864a745d79d2789bad8ff0671`,
			LifeTemplateId:  `a9281ab1e98541368a294eee8831a01d`,
			MiNiPageUrl:     `pages/home/home&query=s%3D1`,
			LifePageUrl:     `alipays://platformapi/startapp?appId=2019060665494484`,
		},
		`ah`: {
			KeyNum:          `ah`,
			AppletAppId:     `2019032863723430`,
			LifeNumberAppId: `2013112700002210`,                                      //待换
			MiNiTemplateId:  `97449fafc2b946e593ece1ec5bddc57c`,                      //待换
			LifeTemplateId:  `54712ee384c04400a3da993a23b91ba4`,                      //待换
			MiNiPageUrl:     `pages/index/index&query=s%3D1`,                         //待换
			LifePageUrl:     `alipays://platformapi/startapp?appId=2019060665494484`, //待换
		},
	},
}}
