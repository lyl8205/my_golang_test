package sanhe_gdmobile

import "time"

const TUsersSdTN = `t_users_sd`

type TUsersSd struct {
	Id               int        `json:"id,omitempty" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`             // 自增id
	Mobile           string     `json:"mobile,omitempty" gorm:"column:mobile;type:varchar(30) DEFAULT;null"`                      // 支付宝绑定的手机号
	Uid              string     `json:"uid,omitempty" gorm:"column:uid;type:varchar(50) DEFAULT;null"`                            // 支付宝用户uid
	NewMobile        string     `json:"new_mobile,omitempty" gorm:"column:new_mobile;type:varchar(30) DEFAULT;null"`              // 需要更换手机号码
	IsUpdate         int        `json:"is_update,omitempty" gorm:"column:is_update;type:tinyint(4);default:0"`                    // 是否切换过手机号 0 不需要|1需要
	UserInfo         string     `json:"user_info,omitempty" gorm:"column:user_info;type:varchar(600) DEFAULT;null"`               // 支付宝用户信息
	CreateTime       *time.Time `json:"create_time,omitempty" gorm:"column:create_time;type:datetime DEFAULT;null"`               // 创建时间
	UpdateTime       *time.Time `json:"update_time,omitempty" gorm:"column:update_time;type:datetime DEFAULT;null"`               // 更新时间
	Scopes           string     `json:"scopes,omitempty" gorm:"column:scopes;type:varchar(30);not null;default:''"`               // 授权类型|topup_service、auth_base、auth_user
	From             int        `json:"from,omitempty" gorm:"column:from;type:tinyint(4);not null;default:0"`                     // 来源|1私域 2插件
	TopupServiceTime *time.Time `json:"topup_service_time,omitempty" gorm:"column:topup_service_time;type:datetime DEFAULT;null"` // 新增授权时间
	Code             string     `json:"code,omitempty" gorm:"column:code;type:varchar(6) DEFAULT;null"`                           // 手机验证码
}
