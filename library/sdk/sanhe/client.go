package sanhe

import (
	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/alipay"
)

type client struct {
	AliClient *alipay.Client
}

func NewClient(appid, privateKey string) *client {
	ali, _ := alipay.NewClient(appid, privateKey, true)
	return &client{AliClient: ali}
}
