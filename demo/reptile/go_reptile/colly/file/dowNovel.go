package file

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"sync"
)

/**
 * @ClassName dowNovel
 * @Description TODO
 * @Author khr
 * @Date 2023/5/19 16:20
 * @Version 1.0
 */

var (
	achan    chan string
	chanTask chan string
	w        sync.WaitGroup
	c        colly.Collector
)

/*
 * @MethodName
 * @Description 错误处理
 * @Author khr
 * @Date 2023/5/19 17:20
 */

func HandleErr(err error) {
	fmt.Errorf("出现错误,err:%v\n", err)
	panic(err)
}

/*
 * @MethodName
 * @Description 检查文件是否存在
 * @Author khr
 * @Date 2023/5/19 17:20
 */

func CheckFile(filePath string) {
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		// 处理打开文件时的错误
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 文件已经存在或创建成功
	fmt.Println("文件已打开或创建成功！")
}

/*
 * @MethodName ReadTxt
 * @Description 读取txt文本里面的超链接
 * @Author khr
 * @Date 2023/5/19 17:18
 */

func ReadTxt(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		HandleErr(err)
	}
	defer file.Close()
	// 创建一个Scanner对象，用于逐行读取文本内容
	scanner := bufio.NewScanner(file)
	var bookLink []string
	// 逐行读取整个文本内容
	for scanner.Scan() {
		line := scanner.Text()
		bookLink = append(bookLink, line)
	}

	for i := 0; i < len(bookLink); i++ {
		w.Add(1)
	}
}
func dow(fullFilePath string) {
	var (
	//title string
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	})
	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// 提取h1标签的文本内容
		name := e.Text

		// 打印文本内容
		fmt.Println("标题名称", name)
		//os.WriteFile("link.txt", []byte(link), 0755)
		//matchPath(link)
	})
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		// 提取h1标签的文本内容
		//title := e.Text
	})
	c.Visit(fullFilePath)
}
