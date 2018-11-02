package Book

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"net/http"
	"net/url"
)

//不能给默认值
type Book struct {
	Name   string
	Zhangs [][2]string
}

func (c *Book) GetBook(rurl string) {
	decoder := mahonia.NewDecoder("gbk")

	resp, err := http.Get(rurl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	rd := decoder.NewReader(resp.Body)

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		println(err)
	}

	//获得章节目录
	var str string
	var value string
	var ok bool
	doc.Find("#list > dl:nth-child(1)").Children().Children().Each(func(i int, s *goquery.Selection) {
		str, err = s.Html()
		if err != nil {
			fmt.Println(err)
		}
		str = s.Text()
		value, ok = s.Attr("href")
		if !ok {
			println("no")
		}
		u, _ := url.Parse(rurl)
		r, _ := url.Parse(value)
		aurl := u.ResolveReference(r).String()

		c.Zhangs = append(c.Zhangs, [2]string{aurl, str})
	})

	//获取书名
	doc.Find("#info > h1:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		c.Name = s.Text()
	})
}
