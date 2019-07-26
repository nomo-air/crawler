package main

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/scheduler"
	"go_crawler/crawler/zhengai/parser"
)

var (
	startUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	seed := engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 200,
		Scheduler:      &scheduler.SimpleScheduler{},
		// Scheduler:      &scheduler.QueuedScheduler{},
	}
	e.Run(seed)
}
