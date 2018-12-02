package parser

import (
	"go_crawler/crawler/engine"
	"regexp"
)

const cityRegex = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	parseResult := engine.ParseResult{}

	reg := regexp.MustCompile(cityRegex)
	match := reg.FindAllSubmatch(contents, -1)

	for _, m := range match {
		parseResult.Items = append(parseResult.Items, string(m[2])) // 城市
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:       string(m[1]), // url
			ParseFunc: engine.NilParseFunc,
		})
	}

	return parseResult
}
