package SaveFile

import (
	. "CrawBook2/Catalog"
	"CrawBook2/tool"
	"fmt"
	"os"
	"strings"
)

type SaveFile struct {
}

func (s *SaveFile) save(c Catalog) {
	//var buf bytes.Buffer
	abuf := make([]string, len(c.Zhangs))

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	for i, u := range c.Zhangs {
		abuf[i] = u[1] + "\n" + tool.GetOneChapter(u[0]) + "\n"
		fmt.Println(u[1])

		//buf.WriteString(u[1])
		//fmt.Println(u[1])
		//buf.WriteString("\n")
		//buf.WriteString(tool.GetOneChapter(u[0]))
		//buf.WriteString("\n")
	}
	str := strings.Join(abuf, "")
	file.Write([]byte(str))
}
