package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
type DictRequest struct {
	Text string `json:"text"`
	Language    string `json:"language"`
}

type DictResponse struct {
	Words []struct {
		Source int `json:"source"`
		Text string `json:"text"`
		PosList []struct {
			Type int `json:"type"`
			Phonetics []struct {
				Type int `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text string `json:"text"`
				Examples []struct {
					Type int `json:"type"`
					Sentences []struct {
						Text string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode int `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{Text: word + "\n", Language: "en"}
	//request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	//var data = strings.NewReader(`{"text":"perfect\n","language":"en"}`)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDGwp/5eSWQMbKT8gyTg&_signature=_02B4Z6wo00001HPPeRgAAIDBLIyWwhFZu1Rzz32AAH6NlrM3W5lZbolqOT1kyVllj6w9a-N18ixLzYxO6bAEGwU3oRGcRTMMnRvwBGSUrH1UpGnpCT52NxBspiuxbjxiOXKBORMcLS9eBYDo7e", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16519190838338140; ttcid=432ddd9b3c9b4189997e34b5acb4606c15; i18next=translate; tt_scid=lXj4-K9bDLrVOBqpjLqbhDDwYi7qBUtlA68r0h-Vrk9HMrF4Nf.U5P-BZDfQFZdMc882")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=en&target_language=zh&text=perfect%0A")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	//反序列化成结构体
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//输出一些结构体里的变量 
	for _, words := range dictResponse.Words {
		fmt.Println(words)
		fmt.Println("---------------------------------")
		for _, post := range words.PosList {
			phone := post.Phonetics
			fmt.Println("phonetics 0" , phone[0].Text)
			fmt.Println("phonetics 1" , phone[1].Text)
			fmt.Println("---------------------------------")
			for _, explan := range post.Explanations {
				fmt.Println("---------------------------------")
				fmt.Println(explan)
			}
		} 
	}

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

