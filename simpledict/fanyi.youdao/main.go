package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"os"
)

type DictRequest struct {
	I string `json:"i"`
	From    string `json:"from"`
	To    string `json:"to"`
	Smartresult string `json:"smartresult"`
	Client string `json:"client"`
	Salt string `json:"salt"`
	Sign string `json:"sign"`
	Lts string `json:"lst"`
	Bv string `json:"bv"`
	Doctype string `json:"doctype"`
	Version string `json:"version"`
	Keyfrom string `json:"keyfrom"`
	Action string `json:"action"`
}

type DictResponse struct {
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	ErrorCode int `json:"errorCode"`
	Type string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type int `json:"type"`
	} `json:"smartResult"`
}
func query(word string) {
	client := &http.Client{}
	//JSON序列化 request body
	//request := DictRequest{I : "good", From : "en", To : "zh-CHS", Smartresult : "dict", Client : "fanyideskweb", Salt : "16519157428258", Sign : "095430769dacfe751dae893b6c3ea187", Lts : "1651915742825", Bv : "ac3968199d18b7367b2479d1f4938ac2", Doctype : "json", Version : "2.1", Keyfrom : "fanyi.web", Action : "FY_BY_REALTlME"}
	//buf, err := json.Marshal(request)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var data = bytes.NewReader(buf)
	var back = "&from=en&to=zh-CHS&smartresult=dict&client=fanyideskweb&salt=16519157428258&sign=095430769dacfe751dae893b6c3ea187&lts=1651915742825&bv=ac3968199d18b7367b2479d1f4938ac2&doctype=json&version=2.1&keyfrom=fanyi.web"
	var front = "i="
	var input = "good"
	front = front + input + back
	fmt.Println(front)
	var data = strings.NewReader(front)
	//var data = strings.NewReader(`i=good&from=en&to=zh-CHS&smartresult=dict&client=fanyideskweb&salt=16519157428258&sign=095430769dacfe751dae893b6c3ea187&lts=1651915742825&bv=ac3968199d18b7367b2479d1f4938ac2&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME`)
	//fmt.Println(data)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=-807176797@10.108.162.138; JSESSIONID=aaap6yrNHz-w76b1uVDcy; OUTFOX_SEARCH_USER_ID_NCOO=266912926.2648601; ___rl__test__cookies=1651913886947")
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	//反序列化成结构体
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//输出一些结构体里的变量 
	fmt.Println(dictResponse.SmartResult.Entries)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
