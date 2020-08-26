package main

import (
	"fmt"
	"strconv"
	"test/shici/author"
	"test/shici/lib"
)


func main() {
	authorList := author.GetAuthor()
	fmt.Println(authorList)
	for _,v := range authorList{
		author := v["author"]
		authorUrl := v["url"]
		lib.CreateDir(author)
		crawlByAuthor(authorUrl,author)
	}

}

func crawlByAuthor(authorUrl,author string)  {
	totalPage := getAuthorPage(authorUrl)
	for i:=1;i<=totalPage;i++{
		url := "https://www.shicimingju.com/chaxun/zuozhe/"+ strconv.Itoa(i) +".html"
		fmt.Println("crawling page " + url)
		var data = make([]map[string]string,0)
		doc := lib.GetDoc(url)
		data = lib.ParseDetailData(doc)
		lib.SaveContent(data,author)
		//author.GetAuthor()
	}
}

func getAuthorPage(url string) int {
	doc := lib.GetDoc(url)
	count := lib.GetAuthorCount(doc)
	return count
}









