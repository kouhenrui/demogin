package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"regexp"
)

var c *colly.Collector

func main() {
	// 创建一个新的 Colly 收集器
	c = colly.NewCollector()
	var bookLink = []string{}
	// 在请求之前设置一些配置，例如设置 User-Agent
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	})

	// 定义变量来保存小说内容
	//var novelContent strings.Builder

	//path := "https://www.xbiquge.tw/"
	//// 访问小说章节页面
	//err := c.Visit(path)
	//
	//fmt.Printf("开始访问网站")
	//if err != nil {
	//	fmt.Printf("访问错误:%s", err)
	//}
	//// 设置 OnHTML 回调函数来处理匹配的 HTML 元素
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	fmt.Println(link)
	//	//os.WriteFile("link.txt", []byte(link), 0755)
	//	//matchPath(link)
	//})
	//c.OnHTML("h1", func(e *colly.HTMLElement) {
	//	// 提取h1标签的文本内容
	//	text := e.Text
	//
	//	// 打印文本内容
	//	fmt.Println(text)
	//	//os.WriteFile("link.txt", []byte(link), 0755)
	//	//matchPath(link)
	//})
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	fmt.Println(link)
	//})

	// 设置 OnError 回调函数处理错误
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("请求出错:", r.Request.URL, "\n错误:", err)
	})

	_, bookLink = readFile()
	for _, y := range bookLink {
		doOneFile(y)
		//fmt.Println(y)
		//c.Visit(y)
	}

}

/*
 * @MethodName matchPath
 * @Description 路径模糊匹配
 * @Author khr
 * @Date 2023/5/18 9:50
 */

func matchPath(link string) {
	filePath := "out.txt"
	// 打开文件并以追加模式写入
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 定义要匹配的正则表达式
	pattern := regexp.MustCompile(`^https://www.xbiquge.tw/book/[^html]*$`)
	if pattern.MatchString(link) {

		if !checkDuplicate(link, filePath) {

			file.WriteString(link + "\n")
			fmt.Println("文本未重复,开始写入文件")
		}

	}
}

/*
 * @MethodName checkDuplicate
 * @Description 检测文本内容是否重复
 * @Author khr
 * @Date 2023/5/18 9:42
 */
func checkDuplicate(value, filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == value {
			return true
		}
	}

	return false
}

/*
 * @MethodName
 * @Description 读取文本内容
 * @Author khr
 * @Date 2023/5/18 9:55
 */
func readFile() (error, []string) {
	// 打开文本文件
	file, err := os.Open("out.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return err, nil
	}
	defer file.Close()

	// 创建一个Scanner对象，用于逐行读取文本内容
	scanner := bufio.NewScanner(file)
	bookLink := []string{}
	// 逐行读取整个文本内容
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	bookLink = append(bookLink, line)
	//}
	for i := 0; i < 2 && scanner.Scan(); i++ {
		bookLink = append(bookLink, scanner.Text())
		//fmt.Println(bookLink)
	}

	// 检查是否发生了扫描错误
	if err = scanner.Err(); err != nil {
		fmt.Println("扫描文本时发生错误:", err)
	}
	return nil, bookLink
}

/*
 * @MethodName
 * @Description
 * @Author khr
 * @Date 2023/5/18 10:17
 */
func doOneFile(links string) {
	fmt.Println(links, "这是一个超链")
	c.Visit(links)
	var name string
	var paths []string
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		// 提取h1标签的文本内容
		name = e.Text

		// 打印文本内容
		fmt.Println("标题名称", name)
		//os.WriteFile("link.txt", []byte(link), 0755)
		//matchPath(link)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		path := e.Attr("href")
		//name := e.Text
		//fmt.Println("超联路径", path)
		paths = append(paths, path)

		//fmt.Println("超联路径合集", paths)
		//filePath := "novel/" + name + ".txt"
		//addFile(filePath)
		//writeToFile(filePath, paths)

	})
	fmt.Println("文本标题", name)
	fmt.Println("超联路径合集", paths)
}
func writeToFile(fileName string, paths []string) {
	// 打开文件并以追加模式写入
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	pattern := regexp.MustCompile(`.*\.html$`)
	for _, y := range paths {
		if pattern.MatchString(y) {
			if !checkDuplicate(fileName, y) {

				file.WriteString(y + "\n")
			}
		}
	}

}
func addFile(filePath string) {

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
