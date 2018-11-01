package Catalog

import (
	"CrawBook/github.com/axgle/mahonia"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

//不能给默认值
type Catalog struct {
	Zhangs [][2]string
}

func (c *Catalog) GetCatalog(rurl string) {
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
}
