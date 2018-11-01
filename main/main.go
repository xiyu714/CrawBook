package main

import (
	"CrawBook2/Catalog"
	"CrawBook2/SaveFile"
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	var c Catalog.Catalog

	c.GetCatalog("http://www.biquge.com.tw/18_18550/")

	var s SaveFile.SaveFile

	s.Save(c)

	fmt.Println("elapsed time:", time.Since(t1))
}
