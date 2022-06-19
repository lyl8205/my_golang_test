package main

import (
	myTool "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"
	"encoding/json"
	"fmt"
	"go_test/abs"
	"go_test/table/sanhe_gdmobile"
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

func main()  {
	var subsprod Subsprod
	var mobileLists []sanhe_gdmobile.TUsersGdMobile
	//获取数据
	usersGd := NewUsersGd()
	usersGd.GetSanHeGdMobile().Table(sanhe_gdmobile.TUsersGdTN).Where("id > ?", 0).Order(`id asc`).Limit(5).Scan(&mobileLists)

	fmt.Printf("%+v\n", mobileLists)

	postData := make(map[string]interface{})
	//postData["mobile"] = "13580756197"

	for _,v := range mobileLists {
		data, _ := json.Marshal(&v)
		if err := json.Unmarshal(data, &postData); err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("postData:%v\n", postData)

		res,_ := curl(postData)
		if err := json.Unmarshal([]byte(res.Body), &subsprod); err != nil {
			fmt.Printf("subsprod解析：%v\n", err)
		}
		fmt.Printf("subsprod返回：%v\n", subsprod)
		promap := subsprod.Data.Result.Prodchinfolist.Prodchinfo
		for _,v := range promap{
			pr,_ := v["prodid"]
			//prod.10086000034679
			if(pr == "prod.10086000001994"){
				//发模板消息
				fmt.Printf("%v\n", "发模板消息")
			}
		}
	}
}

func curl(postData map[string]interface{}) (*myTool.Response, error){
	r := myTool.NewRequest("https://apis.samhotele.com/api/gd/common/testCcqrysubsprods")
	//r := myTool.NewRequest("http://local.erpapi.com/api/JiYunFlow/myTest")

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

