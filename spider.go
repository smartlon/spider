package main

import (
	"fmt"
	"github.com/smartlon/spider/util"
)

func main() {
	content := util.Get(util.URL)
	url, title := util.RootURL(content)
	total := 0
	for k,v := range url {
		conPath := util.PATH+"//"+title[k]
		//util.Mkdir(conPath)
		content2 := util.Get(v)
		url2,title2 := util.SecondURL(content2,v)
		url2Len := len(url2)
		for k2:=1; k2<url2Len;k2++ {
			conPath2 := conPath+"//"
			util.Pdf(url2[k2],conPath2,title2[k2-1])
			//writtern := util.WriteFile(url2[k2],conPath2,title2[k2-1])
			fmt.Printf("url = %s , title = %s  path = %s  number = %d\n",url2[k2],title2[k2-1],conPath2,0)
			total++
			//content3 := util.Get(url2[k2])
			//fmt.Println(content3)
			//util.Mkdir(conPath2)
		}
	}
	fmt.Println(total)

}


