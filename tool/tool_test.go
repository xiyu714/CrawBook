package tool

import "testing"

func TestGetOneChapter(t *testing.T) {
	str := GetOneChapter("http://www.biquge.com.tw/18_18550/8437893.html")

	println(str)
}
