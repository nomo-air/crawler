package engine

import (
	"github.com/gpmgo/gopm/modules/log"
	"go_crawler/crawler/fetcher"
	"time"
)

func Run(seeds ...Request) {
	requests := []Request{}
	for _, request := range seeds {
		requests = append(requests, request)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(request.Url)
		log.Warn("Fetching %s", request.Url)
		if err != nil {
			log.Warn("请求[%s]失败：%s", request.Url, err)
			continue
		}

		parseResult := request.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Warn("Got Item %v", item)
		}
		time.Sleep(time.Millisecond * 10)

	}
}
