package main

import (
	"fmt"
	"provice-city-area-street/conf"
	"provice-city-area-street/fetcher"
	"provice-city-area-street/handler"
	"provice-city-area-street/parser"
)

func main() {
	url := conf.URL

	//step1. 获取所有省份数据
	provinceUrl := fmt.Sprintf(url, "中国", 1, conf.Key)
	fmt.Println(provinceUrl)
	provinceData, err := fetcher.Fetch(provinceUrl)

	if err != nil {
		panic("获取省份数据失败")
	}

	respData, err := parser.Parse(provinceData)
	fmt.Println(respData)
	if err != nil {
		panic("数据解析失败")
	}

	error := handler.Handle(*respData)
	if error != nil {
		panic("数据处理失败")
	}
}
