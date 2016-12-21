package main

import (
	//"io/ioutil"
	"fmt"
	"time"
	//"net/http"
	//"context"
)

type SpiderManage struct {
	Limit int
	Url string
}

func NewSpiderManage() *SpiderManage {
	return &SpiderManage{}
}

func (sm *SpiderManage) Qnum(limit int) *SpiderManage {
	sm.Limit = limit
	return sm
}

func (sm *SpiderManage) Init(url string) *SpiderManage {
	sm.Url = url
	return sm
}

func (sm *SpiderManage) Start() *SpiderManage {
	return sm
}

func (sm *SpiderManage) Stop() *SpiderManage {
	return sm
}


func main() {
	fmt.Println(time.Now())

	sm := NewSpiderManage()
	sm.Qnum(20)
	sm.Init("http://www.baidu.com")
	sm.Start()
	fmt.Println(sm)

	/*
	client := &http.Client{
		//	CheckRedirect:redirectPolicyFunc,
	}

	req,err := http.NewRequest("POST","http://recommended.ptdev.cn/aggregation/rdd/b",nil)
	req.Header.Add("Cookie","session_id=186F95F161A942EFA737EC85B504DEF7CA25F6CB")
	res,err := client.Do(req)
	if err != nil {
		panic("get error")
	}
	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	*/
}

