package SaveFile

import (
	. "CrawBook2/Book"
	"testing"
)

func TestSaveFile(t *testing.T) {
	var s SaveFile
	var c Book

	c.GetBook("http://www.biquge.com.tw/18_18550/")
	s.Save(c)
}
