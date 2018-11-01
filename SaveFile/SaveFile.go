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

var abuf []string

func (s *SaveFile) save(c Catalog) {
	//var buf bytes.Buffer
	abuf = make([]string, len(c.Zhangs))

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	for i, u := range c.Zhangs {
		addOne(i, u)
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

func addOne(i int, u [2]string) {
	abuf[i] = u[1] + "\n" + tool.GetOneChapter(u[0]) + "\n"
}
