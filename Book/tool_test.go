package Book

import "testing"

func TestGetOneChapter(t *testing.T) {
	str := getOneChapter("http://www.biquge.com.tw/18_18550/8437893.html")

	println(str)
}
