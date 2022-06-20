package main

import (
	myTool "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"
	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool/logger"
	"encoding/json"
	"fmt"
	"go_test/abs"
	"go_test/config"
	"go_test/library/sdk/sanhe"
	sanheerpflow "go_test/table/sanhe_erp_flow"
	"go_test/table/sanhe_gdmobile"
	"go_test/task/guangdong_bill/service"
	"time"
)

//广东移动检验接口返回数据格式构造
type Subsprod struct {
	Code int  `json:"code"`
	Msg string  `json:"msg"`
	Data SubResu  `json:"data"`
}
type SubResu struct {
	Result SubResul  `json:"result"`
}
type SubResul struct {
	Prodchinfolist SubProd  `json:"prodchinfolist"`
}

type SubProd struct {
	Prodchinfo []map[string]interface{}  `json:"Prodchinfo"`
}

type PostMobile struct {
	Mobile string
}

type MiniTemplateRes struct {
	User sanhe_gdmobile.TUsersGdMobile
	Result  bool
}



func main()  {
	var (
		ExecutionAuantity = 1
		limitNum = 1
		starId = 0
		subsprod Subsprod
		mobileLists []sanhe_gdmobile.TUsersGdMobile
	)
	//获取数据
	usersGd := NewUsersGd()
	usersGd.GetSanHeGdMobile().Table(sanhe_gdmobile.TUsersGdTN).Where("id > ?", starId).Order(`id asc`).Limit(limitNum).Scan(&mobileLists)

	fmt.Printf("%+v\n", mobileLists)
	var uset []sanhe_gdmobile.TUsersGdMobile
	mobileLists = append(uset, sanhe_gdmobile.TUsersGdMobile{
		Mobile:    `13580756197`,
		Uid:       `2088802919102703`,
	})
	fmt.Printf("%+v\n", mobileLists)

	//计算要开启的goroutine
	corNum := 1
	if limitNum > ExecutionAuantity {
		corNum = limitNum / ExecutionAuantity
		if limitNum%ExecutionAuantity > 0 {
			corNum++
		}
	}
	corNumInt := int(corNum)
	miniTemplateRes := make(chan MiniTemplateRes, limitNum)
	for i := 0;i < corNumInt;i++ {
		start := i * ExecutionAuantity
		end := (i + 1) * ExecutionAuantity
		if end > limitNum {
			end = limitNum
		}
		go func(start,end int){
			defer func() {
				myTool.SafeDefer()
			}()
			mobile := mobileLists[start:end]

			for _,v := range mobile {
				postData := make(map[string]interface{})
				postValue := make(map[string]string)
				data, _ := json.Marshal(&v)
				if err := json.Unmarshal(data, &postData); err != nil {
					fmt.Printf("%v", err)
				}
				if err := json.Unmarshal(data, &postValue); err != nil {
					fmt.Printf("%v", err)
				}
				fmt.Printf("postData:%v\n", postData)

				res,_ := curl(postData)
				if err := json.Unmarshal([]byte(res.Body), &subsprod); err != nil {
					fmt.Printf("subsprod解析：%v\n", err)
				}
				//fmt.Printf("subsprod返回：%v\n", subsprod)
				promap := subsprod.Data.Result.Prodchinfolist.Prodchinfo

				//发小程序模板消息
				for _,vv := range promap{
					pr,_ := vv["prodid"]
					//prod.10086000034679
					if(pr == "prod.10086000034679"){
						fmt.Printf("发小程序消息:%v\n", postData["mobile"])
						//发模板消息
						sendres,_ := sendMessage(postValue["mobile"],postValue["uid"],1)
						fmt.Printf("发小程序消息返回:%v\n", sendres)
						if !sendres {
							fmt.Printf("发生活号消息:%v\n", postData["mobile"])
							sendres,_ = sendMessage(postValue["mobile"],postValue["uid"],2)
							fmt.Printf("发生活号消息返回:%v\n", sendres)
						}

						minRes := MiniTemplateRes{v, sendres}
						miniTemplateRes <- minRes
						break
					}
				}
				//fmt.Printf("sendres:%v\n", sendres)
				//minRes := MiniTemplateRes{v, sendres}
				//miniTemplateRes <- minRes
			}

		}(start,end)
	}

	counter := 0
	successNum := 0
	failNum := 0
	failUser := make([]sanhe_gdmobile.TUsersGdMobile, 0)

Loop:
	for {
		select {
			case resmini := <-miniTemplateRes:
				counter++
				if resmini.Result {
					successNum++
				}else{
					failNum++
					failUser = append(failUser,resmini.User)
				}

				if counter == limitNum {
					fmt.Printf(`进入,successNum==:%d,failNum==:%d,counter:%d`, successNum,failNum,counter)
					close(miniTemplateRes)
					break Loop
				}
			case <-time.After(60 * time.Second):
				close(miniTemplateRes)
				logger.Use(`gd`).Info(fmt.Sprintf(`超时关闭,successNum==:%d,sendType:%d,counter:%d`, successNum, 1, counter))
				break
		}
	}
	fmt.Println("程序结束")
}

func curl(postData map[string]interface{}) (*myTool.Response, error){
	//r := myTool.NewRequest("https://apis.samhotele.com/api/gd/common/testCcqrysubsprods")
	r := myTool.NewRequest("http://local.erpapi.com/api/JiYunFlow/myTest")

	header := make(map[string]string)
	header["Content-Type"] = "application/json;charset=utf-8"

	fmt.Printf("执行：%v\n", postData["mobile"])
	res,err := r.SetHeaders(header).SetPostData(postData).Post()
	return res,err
}

type usersGd struct {
	abs.Model
}

func NewUsersGd() *usersGd {
	return &usersGd{}
}

func sendMessage(mobile,uid string,sendType int) (res bool, err error) {
	var (
		SendTime          = time.Now().AddDate(0, -1, 0)
		alipaySetting sanheerpflow.ShAlipaySetting
		CommonKey         = config.Key.BillKey.CommonKey[`gd`]
	)
	y, m, _ := SendTime.Date()
	reqParam := sanhe.RequestParam{
		Year:  y,
		Month: int(m),
	}
	switch sendType {
	case 1:
		alipaySetting = service.NewAlipaySetting().GetAlipaySetting(CommonKey.AppletAppId)
		reqParam.TemplateId = CommonKey.MiNiTemplateId
		reqParam.PageUrl = CommonKey.MiNiPageUrl
		reqParam.Mobile = mobile
		reqParam.ToUserId = uid
	case 2:
		alipaySetting = service.NewAlipaySetting().GetAlipaySetting(CommonKey.LifeNumberAppId)
		reqParam.TemplateId = CommonKey.LifeTemplateId
		reqParam.PageUrl = CommonKey.LifePageUrl
		reqParam.Mobile = mobile
		reqParam.ToUserId = uid
	}
	if alipaySetting.Id <= 0 {
		myTool.PushSimpleMessage(`支付宝配置有误`, false)
		return false, fmt.Errorf(`支付宝配置有误`)
	}
	aliClient := sanhe.NewClient(alipaySetting.Appid, alipaySetting.RsaPrivateKey)

	if aliClient.AliClient.AppId == `` {
		myTool.PushSimpleMessage(`支付宝配置初始化失败`, false)
		return false, fmt.Errorf(`支付宝配置初始化失败`)
	}

	switch sendType {
	case 1:
		res = aliClient.SendMiniTemplateMessageBill(reqParam)
	case 2:
		res = aliClient.SendMessageSingleBill(reqParam)
	}
	return res,nil
}


