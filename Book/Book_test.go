package Book

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
	"testing"
)

func TestCatalog(t *testing.T) {
	var c Book
	c.GetBook("http://www.biquge.com.tw/18_18550/")
	fmt.Println(c.Zhangs[2])
}

func TestDarray(t *testing.T) {
	var as [][2]string
	//不能使用["hello", "world"]的形式创建数组
	as = append(as, [2]string{"hello", "world"})

	fmt.Println(as)
}

func TestUrl(t *testing.T) {
	u, _ := url.Parse("http://www.biquge.com.tw/18_18550/")
	r, _ := url.Parse("/18_18550/8439408.html")
	s := u.ResolveReference(r)
	if "http://www.biquge.com.tw/18_18550/8439408.html" != s.String() {
		t.Fail()
	} else {
		fmt.Println("ok")
	}
}

func TestDomain(t *testing.T) {
	u := getDomain("http://www.biquge.com.tw/18_18550/")
	fmt.Println(u) //输出www.biquge.com.tw
}

func TestIsSupport(t *testing.T) {
	ok, p := isSupport("http://www.biquge.com.tw/18_18550/")
	fmt.Println(ok)
	fmt.Println(p)
}

func Test网站输出(t *testing.T) {
	w := 网站信息{
		map[string]网站配置{"www.biquge.com.tw": {
			"#info > h1:nth-child(1)",
			[][2]string{{"hello", "world"}},
			"#list > dl:nth-child(1)",
			"#content",
		},
		},
	}

	yfile, _ := os.Create("../config/配置文件.yaml")
	yencoder := yaml.NewEncoder(yfile)
	defer yencoder.Close()
	yencoder.Encode(w)
}

func Test网站读取(t *testing.T) {
	var w 网站信息

	yfile, _ := os.Open("../config/配置文件.yaml")
	ydecoder := yaml.NewDecoder(yfile)
	ydecoder.Decode(&w)

	fmt.Println(w)
}

func TestNo(t *testing.T) { //(*´･д･)?会不会对Book包初始化

}

//testing.T 用来控制测试
