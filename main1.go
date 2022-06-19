package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	jsoniterator "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)


type Payload1 struct {
	Mobile string `json:"mobile"`
}

type Pro struct {
	Code int64  `json:"code"` //状态码
	Msg string  `json:"msg"` //
	Data Resu  `json:"data"` //
}
type Resu struct {
	Result Resul  `json:"result"` //
}
type Resul struct {
	Prodchinfolist Prod  `json:"prodchinfolist"` //
}

type Prod struct {
	Prodchinfo []map[string]interface{}  `json:"Prodchinfo"` //
}

var pool1 = sync.Pool{
	New: func() interface{} {
		return new(Payload1)
	},
}

func main() {
	inputFilePath := "E:/log/abc.csv"
	inputFileObject, openInputFileObjectErr := os.Open(inputFilePath)
	assertOpenInputFileObjectFailure := openInputFileObjectErr != nil
	if assertOpenInputFileObjectFailure {
		developerDescription := "尝试读取要订阅的号码文件失败"
		fmt.Printf("%s:%s", developerDescription, openInputFileObjectErr.Error())
		os.Exit(1)
	}

	defer inputFileObject.Close()

	bufferReader := bufio.NewReader(inputFileObject)

	for {
		currentLine, _, readLineErr := bufferReader.ReadLine()
		if io.EOF == readLineErr {
			break
		}

		if nil != readLineErr {
			break
		}

		//fmt.Printf("%v\n", string(currentLine))
		mobile := string(currentLine[0:11])
		payload1 := pool1.Get().(*Payload1)
		payload1.Mobile = mobile
		json, _ := jsoniterator.Marshal(payload1)
		payload1.Mobile = ""
		pool1.Put(payload1)
		//fmt.Printf("%v\n", string(json))
		currentURL := new(url.URL)
		//currentURL.Scheme = "http"
		//currentURL.Host = "local.sanheerpapidev.com"
		//currentURL.Path = "/api/JiYunFlow/myTest"
		currentURL.Scheme = "https"
		currentURL.Host = "apis.samhotele.com"
		currentURL.Path = "/api/gd/common/testCcqrysubsprods"

		currentRequest := new(http.Request)
		currentRequest.Header = http.Header{"Content-Type": {"application/json;charset=utf-8"}}
		currentRequest.Method = "POST"
		currentRequest.URL = currentURL
		currentRequest.Body = ioutil.NopCloser(bytes.NewReader(json))

		client := generateClient1(true)
		response, doErr := client.Do(currentRequest)
		assertDoRequestFailure := doErr != nil
		if assertDoRequestFailure {
			fmt.Printf("%s,%s\n", mobile, doErr.Error())
			continue
		}

		responseBody, readAllErr := ioutil.ReadAll(response.Body)
		assertReadAllFailure := readAllErr != nil
		if assertReadAllFailure {
			fmt.Printf("%s,%s\n", mobile, readAllErr.Error())
			continue
		}

		pro := &Pro{}
		err := jsoniterator.Unmarshal(responseBody,pro)
		if err != nil{
			fmt.Printf("%v\n", err)
			break
		}

		//fmt.Printf("%s,%s\n", mobile, string(responseBody))
		//fmt.Printf("%v\n", pro)
		promap := pro.Data.Result.Prodchinfolist.Prodchinfo
		//fmt.Printf("%v\n", promap)
		for _,v := range promap{
			pr,_ := v["prodid"]
			if(pr == "prod.10086000034679"){
				//发模板消息
			}
			fmt.Printf("%v\n", pr)
		}
	}
}

func generateClient1(insecureSkipVerify bool) *http.Client {
	client := new(http.Client)

	if !insecureSkipVerify {
		return client
	}

	var defaultConfig tls.Config
	defaultConfig.InsecureSkipVerify = true

	var defaultTransport http.Transport
	defaultTransport.TLSClientConfig = &defaultConfig
	defaultTransport.Proxy = http.ProxyFromEnvironment
	defaultTransport.DialContext = (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext
	defaultTransport.ForceAttemptHTTP2 = true
	defaultTransport.MaxIdleConns = 100
	defaultTransport.IdleConnTimeout = 90 * time.Second
	defaultTransport.TLSHandshakeTimeout = 10 * time.Second
	defaultTransport.ExpectContinueTimeout = 1 * time.Second

	client.Transport = &defaultTransport

	return client
}

