package Catalog

import (
	"CrawBook2/tool"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

//不能给默认值
type Catalog struct {
	id int
}

func (c Catalog) getCatalog(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		println(err)
	}

	doc.Find("#list > dl:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, selection *goquery.Selection) {
			str := selection.Text()
			str = tool.Utf8(str)
			fmt.Println(str)
		})
	})

}
