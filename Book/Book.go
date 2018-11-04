package Book

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

//不能给默认值
type Book struct {
	Name   string
	Zhangs [][2]string
	Filter [][2]string
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
var 当前网站 string
var b *Book

func (c *Book) GetBook(rurl string) {
	okw, 网站 := isSupport("http://www.biquge.com.tw/18_18550/")
	if !okw {
		fmt.Println("此程序当前不支持此网站")
		os.Exit(-1)
	}

	b = c

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
	doc.Find(网站.V章节).Children().Children().Each(func(i int, s *goquery.Selection) {
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
	doc.Find(网站.V书名).Each(func(i int, s *goquery.Selection) {
		c.Name = s.Text()
	})
	c.Filter = w.V网站配置详情[当前网站].V过滤
}

func isSupport(url string) (ok bool, p 网站配置) {
	url = getDomain(url)
	for key, details := range w.V网站配置详情 {
		if url == key {
			return true, details
		}
	}
	ok = false
	return
}

func getDomain(rurl string) (domain string) {
	u, _ := url.Parse(rurl)
	当前网站 = u.Host
	domain = u.Host
	return
}

func init() {
	yfile, err := os.Open("config/配置文件.yaml")
	if err != nil {
		log.Fatal(err)
	}
	ydecoder := yaml.NewDecoder(yfile)
	ydecoder.Decode(&w)
}
