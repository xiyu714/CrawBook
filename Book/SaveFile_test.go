package Book

import (
	"testing"
)

func TestSaveFile(t *testing.T) {
	var c Book

	c.GetBook("http://www.biquge.com.tw/18_18550/")
	c.Save()
}
