package service

import (
	"crontab/abs"
	"crontab/config"
	"crontab/library/sdk/sanhe"
	sanheerpflow "crontab/table/sanhe_erp_flow"
	"fmt"

	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"

	"crontab/table/sanhe_gdmobile"
	"crontab/task/anhui_bill/model"
	"log"
	"strconv"
	"time"

	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool/logger"
)

type billSending struct {
	abs.Service
}

var (
	SendTime          = time.Now().AddDate(0, -1, 0)
	ExecutionAuantity = 60
	CommonKey         = config.Key.BillKey.CommonKey[`ah`]
)

func NewBillSending() *billSending {
	return &billSending{}
}

type MiniTemplateRes struct {
	User sanhe_gdmobile.TUsersAh
	Res  bool
}

// 发送
func (bs *billSending) Send(limit int) {
	key := fmt.Sprintf(config.Key.BillKey.TemplateBillKey, CommonKey.KeyNum, SendTime.Format(`200601`))
	redisCache := bs.GetCache()
	idStr := redisCache.Get(key).Val()
	id, _ := strconv.Atoi(idStr)
	userList := model.NewUsers().GetUsersById(id, limit)
	userNum := len(userList)
	if userNum <= 0 {
		logger.Use(`bill` + CommonKey.KeyNum).Info(fmt.Sprintf(`Send:暂无数据发送`))
		log.Println(`暂无数据发送`)
		return
	}
	bs.SendData(userList, userNum)
}

func (bs *billSending) SendTest() {
	var userList []sanhe_gdmobile.TUsersAh
	userList = append(userList, sanhe_gdmobile.TUsersAh{
		Id:        0,
		Mobile:    `13570929276`,
		Uid:       `2088002780919824`,
		NewMobile: `13570929276`,
	})
	bs.SendData(userList, len(userList))
}

func (bs *billSending) SendData(userList []sanhe_gdmobile.TUsersAh, userNum int) {
	// 发送小程序模板
	successNum, failUser, err := bs.SendTemplateMessage(userList, userNum, 1)
	if err != nil {
		logger.Use(`bill` + CommonKey.KeyNum).Info(err.Error())
		log.Println(err.Error())
		return
	}
	bs.setRedisId(userList[userNum-1])
	bs.Statistics(successNum)
	failUserNum := len(failUser)
	logger.Use(`bill` + CommonKey.KeyNum).Info(fmt.Sprintf(`Send:小程序模板发送成功:%d,发送失败:%d,总数是:%d`, successNum, failUserNum, userNum))
	log.Printf(`Send:小程序模板发送成功:%d,发送失败:%d,总数是:%d`, successNum, failUserNum, userNum)
	if failUserNum <= 0 {
		logger.Use(`bill` + CommonKey.KeyNum).Info(fmt.Sprintf(`Send:发送完毕`))
		log.Println(`发送完毕`)
		return
	}
	//发送生活号
	successNum, failUser, err = bs.SendTemplateMessage(failUser, failUserNum, 2)
	if err != nil {
		logger.Use(`bill` + CommonKey.KeyNum).Info(err.Error())
		log.Println(err.Error())
		return
	}
	bs.Statistics(successNum)
	logger.Use(`bill` + CommonKey.KeyNum).Info(fmt.Sprintf(`Send::生活号模板发送成功:%d,发送失败:%d,总数是:%d`, successNum, len(failUser), userNum))
	log.Printf(`Send:生活号模板发送成功:%d,发送失败:%d,总数是:%d`, successNum, len(failUser), userNum)
}

// 统计发送成功量
func (bs *billSending) Statistics(num int) {
	key := fmt.Sprintf(config.Key.BillKey.TemplateBillNumKey, CommonKey.KeyNum, SendTime.Format(`200601`))
	redisCache := bs.GetCache()
	_ = redisCache.HIncrBy(key, CommonKey.KeyNum, int64(num)).Err()
	if redisCache.TTL(key).Val() == -1*time.Second {
		_ = redisCache.Expire(key, time.Hour*24*15).Err()
	}
}

// 设置最后查询到的ID
func (bs *billSending) setRedisId(user sanhe_gdmobile.TUsersAh) {
	key := fmt.Sprintf(config.Key.BillKey.TemplateBillKey, CommonKey.KeyNum, SendTime.Format(`200601`))
	redisCache := bs.GetCache()
	_ = redisCache.Set(key, user.Id, time.Hour*24*15).Err()
}

// 发送模板消息
//sendType 1.移动小程序,2.移动生活号
func (bs *billSending) SendTemplateMessage(userList []sanhe_gdmobile.TUsersAh, userNum, sendType int) (successNum int, failUser []sanhe_gdmobile.TUsersAh, err error) {
	var (
		alipaySetting sanheerpflow.ShAlipaySetting
	)
	y, m, _ := SendTime.Date()
	reqParam := sanhe.RequestParam{
		Year:  y,
		Month: int(m),
	}
	switch sendType {
	case 1:
		alipaySetting = NewAlipaySetting().GetAlipaySetting(CommonKey.AppletAppId)
		reqParam.TemplateId = CommonKey.MiNiTemplateId
		reqParam.PageUrl = CommonKey.MiNiPageUrl
	case 2:
		alipaySetting = NewAlipaySetting().GetAlipaySetting(CommonKey.LifeNumberAppId)
		reqParam.TemplateId = CommonKey.LifeTemplateId
		reqParam.PageUrl = CommonKey.LifePageUrl

	}
	if alipaySetting.Id <= 0 {
		tool.PushSimpleMessage(`支付宝配置有误`, false)
		return successNum, failUser, fmt.Errorf(`支付宝配置有误`)
	}
	aliClient := sanhe.NewClient(alipaySetting.Appid, alipaySetting.RsaPrivateKey)

	if aliClient.AliClient.AppId == `` {
		tool.PushSimpleMessage(`支付宝配置初始化失败`, false)
		return successNum, failUser, fmt.Errorf(`支付宝配置初始化失败`)
	}
	corNum := 1
	if userNum > ExecutionAuantity {
		corNum = userNum / ExecutionAuantity
		if userNum%ExecutionAuantity > 0 {
			corNum++
		}
	}
	corNumInt := int(corNum)
	miniTemplateRes := make(chan MiniTemplateRes, userNum)
	for i := 0; i < corNumInt; i++ {
		start := i * ExecutionAuantity
		end := (i + 1) * ExecutionAuantity
		if end > userNum {
			end = userNum
		}
		go func(start, end int) {
			defer func() {
				tool.SafeDefer()
			}()
			user := userList[start:end]
			// var success, err int
			for _, v := range user {
				if v.NewMobile == `` {
					continue
				}
				var res bool
				reqParam.Mobile = v.NewMobile
				reqParam.ToUserId = v.Uid
				switch sendType {
				case 1:
					res = aliClient.SendMiniTemplateMessageBill(reqParam)
				case 2:
					res = aliClient.SendMessageSingleBill(reqParam)
				}
				minRes := MiniTemplateRes{v, res}
				miniTemplateRes <- minRes
			}
			// log.Printf("Mobile:%s 成功数:%d,失败数:%d,总数:%d,start:%d, end:%d", user[0].Mobile, success, err, len(user), start, end)
		}(start, end)
	}

	counter := 0
	failUser = make([]sanhe_gdmobile.TUsersAh, 0)

Loop:
	for {
		select {
		case res := <-miniTemplateRes:
			counter++
			if res.Res {
				successNum++
			} else {
				failUser = append(failUser, res.User)
			}
			if counter == userNum {
				close(miniTemplateRes)
				break Loop
			}
		case <-time.After(120 * time.Second):
			close(miniTemplateRes)
			logger.Use(`bill` + CommonKey.KeyNum).Info(fmt.Sprintf(`超时关闭,successNum==:%d,sendType:%d,counter:%d`, successNum, sendType, counter))
			break Loop
		}
	}
	return successNum, failUser, nil
}
