package SaveFile

import (
	. "CrawBook2/Catalog"
	"testing"
)

func TestSaveFile(t *testing.T) {
	var s SaveFile
	var c Catalog

	c.GetCatalog("http://www.biquge.com.tw/18_18550/")
	s.save(c)
}
