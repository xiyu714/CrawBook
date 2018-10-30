package Catalog

import (
	"CrawBook2/tool"
	"github.com/PuerkitoBio/goquery"
)

//不能给默认值
type Catalog struct {
	zhangs [][2]string
}

func (c Catalog) getCatalog(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		println(err)
	}

	doc.Find("#list > dl:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, selection *goquery.Selection) {
			str := selection.Text()
			value, ok := selection.Attr("href")
			if !ok {
				println("no")
			}
			str = tool.Utf8(str)
			println(str)
			println(value)
			c.zhangs = append(c.zhangs, [2]string{value, str})
		})
	})

}
