package main

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/persist"
	"go_crawler/crawler/scheduler"
	"go_crawler/crawler/zhengai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("profiles")
	if err != nil {
		panic(err)
	}
	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:       "http://www.zhenai.com/zhenghun/beijing",
			ParseFunc: parser.ParseCity,
		},
		{
			Url:       "http://www.zhenai.com/zhenghun/henan",
			ParseFunc: parser.ParseCity,
		},
		{
			Url:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
		},
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 200,
		Scheduler:      &scheduler.QueuedScheduler{},
		ItemChan:       itemChan,
	}
	e.Run(seed...)
}
