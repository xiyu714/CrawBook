package Book

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"net/http"
	"strings"
)

func getOneChapter(url string) (str string) {
	decoder := mahonia.NewDecoder("gbk")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	rd := decoder.NewReader(resp.Body) //对io流进行解码

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		println(err)
	}

	//获取文档中的内容
	doc.Find(w.V网站配置详情[当前网站].V内容).Each(func(i int, selection *goquery.Selection) {
		str, _ = selection.Html()
		//str = strings.Replace(str, "<br/>\n<br/>\n", "\r\n", -1)
		//str = strings.Replace(str, "小?說", "", -1)
		for _, rs := range b.Filter {
			str = strings.Replace(str, rs[0], rs[1], -1)
		}
	})

	return
}
