package Catalog

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCatalog(t *testing.T) {
	var c Catalog
	c.GetCatalog("http://www.biquge.com.tw/18_18550/")
	fmt.Println(c.Zhangs[2])
}

func TestDarray(t *testing.T) {
	var as [][2]string
	//不能使用["hello", "world"]的形式创建数组
	as = append(as, [2]string{"hello", "world"})

	fmt.Println(as)
}

func TestUrl(t *testing.T) {
	u, _ := url.Parse("http://www.biquge.com.tw/18_18550/")
	r, _ := url.Parse("/18_18550/8439408.html")
	s := u.ResolveReference(r)
	if "http://www.biquge.com.tw/18_18550/8439408.html" != s.String() {
		t.Fail()
	} else {
		fmt.Println("ok")
	}
}

//testing.T 用来控制测试
