package Catalog

import (
	"CrawBook/github.com/axgle/mahonia"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

//不能给默认值
type Catalog struct {
	zhangs [][2]string
}

func (c Catalog) getCatalog(url string) {
	decoder := mahonia.NewDecoder("gbk")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	rd := decoder.NewReader(resp.Body)

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		println(err)
	}

	//var str string
	//var value string
	//var ok bool
	//doc.Find("#list > dl:nth-child(1)").Each(func(i int, s *goquery.Selection) {
	//	s.Children().Each(func(j int, selection *goquery.Selection) {
	//		//strh, err := selection.Html()
	//		//if err != nil {
	//		//	fmt.Println(err)
	//		//}
	//		//-- 测试能否读取内部html
	//		str = selection.Text()
	//		value, ok = selection.Attr("href")
	//		if !ok {
	//			println("no")
	//		}
	//		println(str)
	//		println(value)
	//		c.zhangs = append(c.zhangs, [2]string{value, str})
	//	})
	//})
	//-- 之所以出现这样的问题，是因为我根本就不知道我在操作什么

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
		c.zhangs = append(c.zhangs, [2]string{value, str})
	})
	fmt.Println(c.zhangs[2])

}
