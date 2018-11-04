package Book

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"gopkg.in/yaml.v2"
	"net/http"
	"net/url"
	"os"
)

//不能给默认值
type Book struct {
	Name   string
	Zhangs [][2]string
}

type 网站信息 struct {
	V网站配置详情 map[string]网站配置 `yaml:"网站配置详情"`
}

type 网站配置 struct {
	V书名 string      `yaml:"书名"`
	V过滤 [][2]string `yaml:"过滤,flow"` //key,tag  ,后面不能有空格
	V章节 string      `yaml:"章节"`
}

var w 网站信息

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

func init() {
	var w 网站信息

	yfile, _ := os.Open("../config/配置文件.yaml")
	ydecoder := yaml.NewDecoder(yfile)
	ydecoder.Decode(&w)
}
