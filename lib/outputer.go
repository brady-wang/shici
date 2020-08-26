package lib

import (
	"fmt"
	"io/ioutil"
)

func SaveContent(data []map[string]string,author string)  {
	for _,v := range data{
		name := v["title"]+v["time"]
		file := ioutil.WriteFile("./txt/"+author+"/"+name+".txt", []byte(v["content"]), 0777)
		if file != nil {
			fmt.Printf("写入文件%s 失败\n", name)
		} else{
			fmt.Printf("写入文件%s 成功\n", name)
		}
		fmt.Printf("%s\n",name)
	}
}