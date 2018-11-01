package main

import (
	"CrawBook2/Book"
	"CrawBook2/SaveFile"
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	var c Book.Book

	c.GetBook("http://www.biquge.com.tw/18_18550/")

	var s SaveFile.SaveFile

	s.Save(c)

	fmt.Println("elapsed time:", time.Since(t1))
}
