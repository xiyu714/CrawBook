package Book

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

var abuf []string
var wg sync.WaitGroup

const max = 1000

var n = 0

func (c *Book) Save() {
	//var buf bytes.Buffer
	abuf = make([]string, len(c.Zhangs))
	filename := c.Name + ".txt"
	outputFile := path.Join("output", filename)

create:
	file, err := os.Create(outputFile)
	if err != nil {
		if _, ok := err.(*os.PathError); ok { //通过类型断言判断错误类型
			os.Mkdir("output", os.ModePerm)
			goto create
		} else {
			log.Fatal(err)
		}
	}
	for i, u := range c.Zhangs {
		wg.Add(1)
		n++
		go addOne(i, u)
		if n > max {
			wg.Wait()
			n = 0
		}

		//buf.WriteString(u[1])
		//fmt.Println(u[1])
		//buf.WriteString("\n")
		//buf.WriteString(tool.getOneChapter(u[0]))
		//buf.WriteString("\n")
	}
	wg.Wait()
	str := strings.Join(abuf, "")
	file.Write([]byte(str)) //这里可能有问题
}

func addOne(i int, u [2]string) {
	abuf[i] = u[1] + "\r\n" + getOneChapter(u[0]) + "\r\n\r\n"
	fmt.Println(u[1])
	wg.Done()
}
