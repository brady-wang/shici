package lib

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)



func GetAuthorCount(doc *html.Node ) int  {

	list := htmlquery.Find(doc,"//*[@id='list_nav_all']/a")
	return len(list)
}


func ParseDetailData(doc *html.Node ) []map[string]string  {
	var data = make([]map[string]string,0)

	list := htmlquery.Find(doc,"//div[@class='card shici_card']//div[@class='shici_list_main']/h3")
	fmt.Println(len(list))
	for _,li :=  range list{
		var mp = make(map[string]string,0)
		title := htmlquery.FindOne(li,".//a")
		href := htmlquery.FindOne(li,".//a")
		mp["title"] = htmlquery.InnerText(title)
		mp["href"] = htmlquery.SelectAttr(href,"href")
		data = append(data, mp)

	}

	for _,v := range data{
		for kk,vv := range v{
			if kk == "href"{
				content := content(vv)
				v["content"] = content
			}
		}
	}
	return data
}

// 获取诗词内容
func content(url string) string  {
	time.Sleep(time.Millisecond * time.Duration(DelayTime))
	url = "https://www.shicimingju.com"+url
	fmt.Println("crawling page "+url)
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
	list := htmlquery.Find(doc,"//div[@id='zs_content']//text()")

	var data string
	for _,li :=  range list{
		data +=  strings.Replace(strings.Replace(htmlquery.InnerText(li), " ", "", -1) , "\n", "", -1)
	}
	return data
}