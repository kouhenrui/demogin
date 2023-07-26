package main

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"sync"
)

var (
	c      *colly.Collector
	path   = "https://www.xbiquge.tw/book/55710/39995548.html"
	titles string

	contents string

	data      []string
	dataMutex sync.Mutex
)
var ctx = context.Background()

func main() {
	c = colly.NewCollector()
	// 在请求之前设置一些配置，例如设置 User-Agent
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		// 提取h1标签的文本内容
		title := e.Text
		err := writeToClick(title + "\n")
		if err != nil {
			fmt.Println("写入数据库错误")
		} else {
			fmt.Println("写入数据成功")
		}
		//fmt.Println(title)
		//dataMutex.Lock()
		////titles = title
		//data = append(data, title)
		//defer dataMutex.Unlock()
	})
	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		// 提取h1标签的文本内容
		content := e.Text
		err := writeToClick(content + "\n")
		if err != nil {
			fmt.Println("写入数据库错误")
		} else {
			fmt.Println("写入数据成功")
		}
		//fmt.Println(content)
		//dataMutex.Lock()
		//contents = content
		//data = append(data, content)
		//defer dataMutex.Unlock()
	})
	//fmt.Println(len(data))
	//if len(data) > 0 {
	//	fmt.Println("拿到了完整的章节标题和内容")
	//	for _, i := range data {
	//		fmt.Print(i)
	//		err := writeToClick(i)
	//
	//		if err != nil {
	//			fmt.Println("写入数据库错误")
	//		} else {
	//			fmt.Println("写入数据成功")
	//		}
	//	}
	//
	//}

	err := c.Visit(path)
	if err != nil {
		fmt.Printf("访问错误:%s", err)
	}
}
func writeToClick(data string) error {
	file, err := os.OpenFile("novel.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return err
	}
	_, err = file.Write([]byte(data))
	//err := click.ClickDB.Exec(ctx, "insert into default.novel (id,title,content) values (3,?,?)", title, content)
	//_, err := click.ClickDB.Query(ctx, "insert into default.novel (id,title,content) values (3,titles,contents)")
	if err != nil {
		fmt.Println(err, "e")
		return err
	}

	return nil
}
