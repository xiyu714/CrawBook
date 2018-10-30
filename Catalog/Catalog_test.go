package Catalog

import (
	"fmt"
	"testing"
)

func TestCatalog(t *testing.T) {
	var c Catalog
	c.getCatalog("http://www.biquge.com.tw/18_18550/")
}

func TestDarray(t *testing.T) {
	var as [][2]string
	//不能使用["hello", "world"]的形式创建数组
	as = append(as, [2]string{"hello", "world"})

	fmt.Println(as)

}
