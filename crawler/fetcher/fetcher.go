package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("抓取出错了, 返回码[%d]", resp.StatusCode)
	}
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(resp.Body).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	return body, err
}

func determineEncoding(r io.Reader) encoding.Encoding {
	data, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(data, "")
	return e
}
