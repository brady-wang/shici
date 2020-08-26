package author

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"test/shici/lib"
	"time"
)



func GetAuthor() []map[string] string {
	authorBaseUrl := "https://www.shicimingju.com/category/all"
	var data  = make([]map[string]string ,0)


	for i:=1;i<=652;i++{
		authorUrl := authorBaseUrl +"_"+ strconv.Itoa(i)
		doc := getDoc(authorUrl)
		dt := parseData(doc)
		data = append(data, dt)
		fmt.Print(data)
	}
	return data
}


func parseData(doc *html.Node ) map[string] string  {


	list := htmlquery.Find(doc,"//div[@id='main_left']/div[@class='card zuozhe_card']")

	var mp = make(map[string]string,0)
	for _,li :=  range list{
		authorNode := htmlquery.FindOne(li,"//div[@class='zuozhe_list_item']/h3/a")
		author := htmlquery.InnerText(authorNode)

		urlNode := htmlquery.FindOne(li,"//div[@class='zuozhe_list_item']/h3/a")
		url := htmlquery.SelectAttr(urlNode,"href")
		url = "https://www.shicimingju.com"+url
		mp["author"] = author
		mp["url"] = url
	}

	return mp
}


// 返回doc类型 给htmlquery 解析
func getDoc(url string) *html.Node  {
	time.Sleep(time.Millisecond * time.Duration(lib.DelayTime))
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

