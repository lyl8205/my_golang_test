package main

import (
	myTool "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"
	"encoding/json"
	"fmt"
)

//广东移动检验接口返回数据格式构造
type Subsprod struct {
	Code int64  `json:"code"`
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
	res,_ := curl([]PostMobile{})
	if err := json.Unmarshal([]byte(res.Body), &subsprod); err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("%v", subsprod)
	promap := subsprod.Data.Result.Prodchinfolist.Prodchinfo
	for _,v := range promap{
		pr,_ := v["prodid"]
		//prod.10086000034679
		if(pr == "JYPT999.200828390257.0"){
			//发模板消息
			fmt.Printf("%v\n", "发模板消息")
			break
		}
		fmt.Printf("%v\n", pr)
	}
}

func curl(postdata []PostMobile) (*myTool.Response, error){
	r := myTool.NewRequest("https://apis.samhotele.com/api/gd/common/testCcqrysubsprods")
	//r := myTool.NewRequest("http://local.erpapi.com/api/JiYunFlow/myTest")

	header := make(map[string]string)
	header["Content-Type"] = "application/json;charset=utf-8"

	res,err := r.SetHeaders(header).SetPostData(postdata).Post()
	return res,err
}
