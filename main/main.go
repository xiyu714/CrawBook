package main

import (
	"CrawBook2/tool"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	//获取html Response
	resp, err := http.Get("http://www.biquge.com.tw/18_18550/8437893.html")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//获得*Node
	doc, err := html.Parse(resp.Body) //resp.Body ioreadcloser
	if err != nil {
		fmt.Println(err)
	}

	parse(doc)

	str := tool.Utf8(buf.String())

	//存入文件
	file, err := os.Create("output/test.txt")
	if err != nil {
		fmt.Println("open", err)
	}

	strR := strings.NewReader(str)
	strR.WriteTo(file)
}

var hasfind = false
var buf bytes.Buffer

func parse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "content" {
				hasfind = true
				parseTxt(&buf, n)
				break
			}
		}
	}

	if !hasfind {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}
}

func parseTxt(buf *bytes.Buffer, n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "br" { //c.Data tag name for element nodes, content of text
			buf.WriteString(c.Data + "\n")
		}
	}
}
