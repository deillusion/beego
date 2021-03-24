package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)


var client = &http.Client{}
func main() {
	//生成client 参数为默认
	
	fmt.Println("started")
	//生成要访问的url
	url := "https://cdn1.vectorstock.com/i/1000x1000/70/80/red-sakura-flower-seamless-pattern-vector-14257080.jpg"
	path := "/opt/gocode/src/xjfz/images/red-sakura-flower-seamless-pattern-vector-14257080.jpg"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	reqest.Header.Add("Cookie", "bft0fi8q7inc3nl8pkfbuig1no; _gid=GA1.2.1884800018.1616491379; _ga=GA1.2.605282523.1616491379; _ga_QBXWNL4G7Q=GS1.1.1616497965.2.1.1616498311.0")
	reqest.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\"")
	reqest.Header.Add("sec-ch-ua-mobile", "?0")
	reqest.Header.Add("sec-fetch-dest", "document")
	reqest.Header.Add("sec-fetch-mode", "navigate")
	reqest.Header.Add("sec-fetch-site", "none")
	reqest.Header.Add("sec-fetch-user", "?1")
	reqest.Header.Add("upgrade-insecure-requests", "1")
	reqest.Header.Add("if-modified-since", "Sat, 06 Jan 2018 01:18:05 GMT")
	reqest.Header.Add("if-none-match", "bd7326c5c743ad38b3db0ba1ee379306")
	fmt.Println("headers get")
	if err != nil {
		panic(err)
	}
	
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	fmt.Println("response gets")
	body, _ := ioutil.ReadAll(response.Body)
	out, _ := os.Create(path)
    io.Copy(out, bytes.NewReader(body))
   //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//stdout := os.Stdout
	//_, err = io.Copy(stdout, response.Body)
   
   //返回的状态码
	//status := response.StatusCode

	//fmt.Println(status)

}