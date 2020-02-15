package util

import (
	"strings"
)

func RootURL(content string)([]string,[]string) {
	contentUr :=strings.Split(content,`<li class="fl"><a href=".`)
	contentTi :=strings.Split(content,`" target="_blank">`)

	conLen := len(contentUr)
	url := make([]string, conLen-1)
	for i := 1; i<conLen;i++ {
		subIndex := strings.Index(contentUr[i],`"`)
		if subIndex !=-1 {
			//fmt.Println( URL+ contentUr[i][1:subIndex])
			url[i-1] =  URL+ contentUr[i][1:subIndex]
		}

	}

	conTLen := len(contentTi)
	title := make([]string, conTLen-1)
	for i := 1; i<conTLen;i++ {
		subIndex := strings.Index(contentTi[i],`</a></li>`)
		if subIndex !=-1 {
			//fmt.Println(contentTi[i][:subIndex])
			title[i-1] = contentTi[i][:subIndex]
		}

	}
	return url,title
}

func SecondURL(content string,URL string)([]string,[]string) {
	contentUr :=strings.Split(content,`<a href="./`)
	contentTi :=strings.Split(content,`</a></h6>`)

	conLen := len(contentUr)
	url := make([]string, conLen-1)
	for i := 1; i<conLen;i++ {
		subIndex := strings.Index(contentUr[i],`"`)
		if subIndex !=-1 {
			//fmt.Println( URL+ contentUr[i][1:subIndex])
			url[i-1] =  URL+ contentUr[i][:subIndex]
			//fmt.Println(url[i-1])
		}

	}

	conTLen := len(contentTi)
	title := make([]string, conTLen-1)
	for i := 0; i<conTLen-1;i++ {
		subIndex2 := strings.LastIndex(contentTi[i],`">`)
		if  subIndex2 !=-1 {
			//fmt.Println(contentTi[i][:subIndex])
			title[i] = contentTi[i][subIndex2+2:]
			//fmt.Println(title[i])

		}

	}
	return url,title
}