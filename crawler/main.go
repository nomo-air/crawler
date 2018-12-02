package main

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/zhengai/parser"
)

var (
	startUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	})
}
