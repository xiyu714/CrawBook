package SaveFile

import (
	. "CrawBook2/Catalog"
	"CrawBook2/tool"
	"bytes"
	"fmt"
	"os"
)

type SaveFile struct {
}

func (s *SaveFile) save(c Catalog) {
	var buf bytes.Buffer

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, u := range c.Zhangs {
		buf.WriteString(u[1])
		fmt.Println(u[1])
		buf.WriteString("\n")
		buf.WriteString(tool.GetOneChapter(u[0]))
		buf.WriteString("\n")
	}

	file.Write(buf.Bytes())
}
