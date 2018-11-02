package main

import (
	"CrawBook2/Book"
	"CrawBook2/SaveFile"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t1 := time.Now()
	//设置命令行解析
	parseArguments()

	var c Book.Book

	//c.GetBook("http://www.biquge.com.tw/18_18550/")
	c.GetBook(url)

	var s SaveFile.SaveFile

	s.Save(c)

	fmt.Println("elapsed time:", time.Since(t1))
}

var url string

func parseArguments() {
	var h bool
	flag.StringVar(&url, "u", "", "小说网站地址")
	flag.BoolVar(&h, "h", false, "show help")
	flag.Parse()
	if h {
		flag.Usage()
	}
	if url == "" {
		fmt.Println("请输入网站地址")
		flag.Usage()
		os.Exit(-1)
	}
}
