package lib

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"net/http"
	"time"
)

const DelayTime = 10

// 返回doc类型 给htmlquery 解析
func GetDoc(url string) *html.Node  {
	time.Sleep(time.Millisecond * time.Duration(DelayTime))
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "UM_distinctid=1742844da1da8f-0e6a940f833873-3323765-1fa400-1742844da1e1f0; Hm_lvt_6181c6e4531be00fa696eb631fbd2a4b=1598403436; CNZZDATA1278992071=1940026881-1598401404-%7C1598406805; Hm_lpvt_6181c6e4531be00fa696eb631fbd2a4b=1598407702")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	doc ,_ := htmlquery.Parse(resp.Body)
	return doc
}