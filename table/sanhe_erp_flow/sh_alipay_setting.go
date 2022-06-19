package sanhe_erp_flow

const ShAlipaySettingTN = `sh_alipay_setting`

type ShAlipaySetting struct {
	Id                 int    `json:"id,omitempty" gorm:"column:id;type:smallint(5);unsigned;not null;primaryKey;autoIncrement"`
	AppTitle           string `json:"app_title,omitempty" gorm:"column:app_title;type:varchar(80);not null;default:''"`            // 支付宝应用描述
	AlipayAccount      string `json:"alipay_account,omitempty" gorm:"column:alipay_account;type:varchar(128);not null;default:''"` // 支付宝账户
	Appid              string `json:"appid,omitempty" gorm:"column:appid;type:varchar(128);not null"`                              // 支付宝应用id
	Pid                string `json:"pid,omitempty" gorm:"column:pid;type:varchar(128);not null;default:''"`                       // 发卡卷会用到
	RsaPrivateKey      string `json:"rsa_private_key,omitempty" gorm:"column:rsa_private_key;type:text;not null"`                  // 商户私钥
	AlipayRsaPublicKey string `json:"alipay_rsa_public_key,omitempty" gorm:"column:alipay_rsa_public_key;type:text;not null"`      // 支付宝公钥
	AuthToken          string `json:"auth_token,omitempty" gorm:"column:auth_token;type:varchar(60);not null;default:''"`          // 用户授权token
}
