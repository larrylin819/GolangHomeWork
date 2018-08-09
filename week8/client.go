package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

func httpGet(username string) {
	resp, err := http.Get("http://127.0.0.1:8888/search/" + username)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost(user string, pwd string) {
	v := url.Values{}
	v.Set("user", user)
	v.Set("pwd", pwd)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}//客户端,被Get,Head以及Post使用
	reqest, err := http.NewRequest("POST", "http://127.0.0.1:8888/login", body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//给一个key设定为响应的value.
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest)//发送请求
	defer resp.Body.Close()//一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func main() {
	httpGet("larry")
	httpPost("larry", "1234")
	httpGet("abcdef")
	httpPost("abcdef", "abcdef")
}
