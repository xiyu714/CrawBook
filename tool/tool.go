package tool

import (
	"CrawBook/github.com/axgle/mahonia"
	"strings"
)

func Utf8(str string) string {
	decoder := mahonia.NewDecoder("gbk")
	s := decoder.ConvertString(str)
	bs := strings.Replace(s, "聽", " ", -1)
	bs = strings.Replace(bs, "\n\n\n\n", "\n", -1)
	bs = strings.Replace(bs, "小?說", "", -1)

	return bs
}
