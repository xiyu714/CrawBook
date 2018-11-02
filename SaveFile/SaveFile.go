package SaveFile

import (
	. "CrawBook2/Book"
	"CrawBook2/tool"
	"fmt"
	"os"
	"strings"
	"sync"
)

type SaveFile struct {
}

var abuf []string
var wg sync.WaitGroup

func (s *SaveFile) Save(c Book) {
	//var buf bytes.Buffer
	abuf = make([]string, len(c.Zhangs))

	file, err := os.Create(c.Name + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	for i, u := range c.Zhangs {
		wg.Add(1)
		go addOne(i, u)

		//buf.WriteString(u[1])
		//fmt.Println(u[1])
		//buf.WriteString("\n")
		//buf.WriteString(tool.GetOneChapter(u[0]))
		//buf.WriteString("\n")
	}
	wg.Wait()
	str := strings.Join(abuf, "")
	file.Write([]byte(str)) //这里可能有问题
}

func addOne(i int, u [2]string) {
	abuf[i] = u[1] + "\r\n" + tool.GetOneChapter(u[0]) + "\r\n\r\n"
	fmt.Println(u[1])
	wg.Done()
}
