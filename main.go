package main

import (
	myTool "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"
	"encoding/json"
	"fmt"
	"go_test/abs"
	"go_test/table/sanhe_gdmobile"
	//"go_test/task/guangdong_bill/service"
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
		ExecutionAuantity = 60
		limitNum = 200
		starId = 0
		subsprod Subsprod
		mobileLists []sanhe_gdmobile.TUsersGdMobile
	)
	//获取数据
	usersGd := NewUsersGd()
	usersGd.GetSanHeGdMobile().Table(sanhe_gdmobile.TUsersGdTN).Where("id > ?", starId).Order(`id asc`).Limit(limitNum).Scan(&mobileLists)

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
				data, _ := json.Marshal(&v)
				if err := json.Unmarshal(data, &postData); err != nil {
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
				for _,v := range promap{
					pr,_ := v["prodid"]
					//prod.10086000034679
					if(pr == "JYPT999.200828390257.0"){
						//发模板消息
						//service.NewBillSending().SendMobileData()
						//fmt.Printf("%v\n", "发模板消息")
					}
				}
				var result bool
				result = true
				minRes := MiniTemplateRes{v, result}
				miniTemplateRes <- minRes
			}

		}(start,end)
	}

	counter := 0
	successNum := 0
	failUser := make([]sanhe_gdmobile.TUsersGdMobile, 0)

Loop:
	for {
		select {
			case resmini := <-miniTemplateRes:
				counter++
				if resmini.Result {
					successNum++
				}else{
					failUser = append(failUser,resmini.User)
				}
				
				if counter == limitNum {
					fmt.Printf(`进入,successNum==:%d,sendType:%d,counter:%d`, successNum, 1, counter)
					close(miniTemplateRes)
					break Loop
				}
			//case <-time.After(60 * time.Second):
			//	close(miniTemplateRes)
			//	fmt.Println("超时关闭")
			//	logger.Use(`gd`).Info(fmt.Sprintf(`超时关闭,successNum==:%d,sendType:%d,counter:%d`, successNum, 1, counter))
			//	done<-true
			//	break
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


