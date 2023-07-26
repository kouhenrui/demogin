package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"sync"
	"time"
)

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	name     string
	content  []string
	paths    []string
)

func main() {
	c := colly.NewCollector()
	// 设置请求超时时间为10秒
	c.SetRequestTimeout(10 * time.Second)
	// 在请求之前设置一些配置，例如设置 User-Agent
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	})
	// 设置 OnError 回调函数处理错误
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("错误为：", err)
	})

	// 1.初始化管道
	//chanImageUrls = make(chan string, 1000000)
	//chanTask = make(chan string, 26)
	//// 2.爬虫协程
	//for i := 1; i < 27; i++ {
	//	waitGroup.Add(1)
	//	//go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	//}
	//// 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	//waitGroup.Add(1)
	//go CheckOK()
	//// 4.下载协程：从管道中读取链接并下载
	//for i := 0; i < 5; i++ {
	//	waitGroup.Add(1)
	//	go DownloadImg()
	//}
	//waitGroup.Wait()
	var bookLink []string
	_, bookLink = readFile()
	for _, y := range bookLink {
		fmt.Println("访问地址", y)
		c.OnHTML("h1", func(e *colly.HTMLElement) {
			// 提取h1标签的文本内容
			name = e.Text
			// 打印文本内容
			fmt.Println("标题名称", name)
		})
		c.OnHTML("div#content[name=content]", func(e *colly.HTMLElement) {
			txt := e.Text
			content = append(content, txt)
			//fmt.Println("Content:", content)
		})

		//https://www.xbiquge.tw/book/54523/38644957.html
		err := c.Visit("https://www.xbiquge.tw/book/54523/" + y)
		if err != nil {
			fmt.Errorf("错误是：%s", err)
		}
		err = writeToNovel(name, content)
		if err != nil {
			fmt.Printf("文本写入错误：%s", err)
		}
	}
	//
	//	//fmt.Println("访问的地址：", y)
	//
	//	c.OnHTML("h1", func(e *colly.HTMLElement) {
	//		// 提取h1标签的文本内容
	//		name = e.Text
	//		// 打印文本内容
	//		fmt.Println("标题名称", name)
	//	})
	//	c.OnHTML("dd a[href]", func(e *colly.HTMLElement) {
	//		path := e.Attr("href")
	//		paths = append(paths, path)
	//
	//	})
	//
	//	err := c.Visit(y)
	//	if err != nil {
	//		fmt.Errorf("错误是：%s", err)
	//	}
	//
	//	fmt.Println("文本标题", name)
	//	//fmt.Println("超联路径合集", paths)
	//	fmt.Println("已经获取到所有的数据")
	//	//拿到了标题和内部章节的所有路劲
	//	err = writeToNovelTxt(name, paths)
	//	if err != nil {
	//		fmt.Errorf("写入文件错误：%s", err)
	//	}
	//	fmt.Println("文件写入完毕")
	//}
}
func readFile() (error, []string) {
	// 打开文本文件
	file, err := os.Open("宇宙职业选手.txt")
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
	for scanner.Scan() {
		line := scanner.Text()
		bookLink = append(bookLink, line)
	}
	//for i := 0; i < scanner.Scan(); i++ {
	//	bookLink = append(bookLink, scanner.Text())
	//	//fmt.Println(bookLink)
	//}

	// 检查是否发生了扫描错误
	if err = scanner.Err(); err != nil {
		fmt.Println("扫描文本时发生错误:", err)
	}
	return nil, bookLink
}

func writeToNovel(name string, content []string) error {
	//文本存放地址及文本名称
	filePath := "novel/宇宙职业选手" + name + ".txt"
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		// 处理打开文件时的错误
		fmt.Println("无法打开和创建文件:", err)
		return err
	}
	//写入结束后关闭文件
	defer file.Close()
	// 文件已经存在或创建成功
	fmt.Println("文件已打开或创建成功！")
	for _, y := range content {
		//	if !checkDuplicate(filePath, y) {
		_, err = file.WriteString(y + "\n")
		if err != nil {
			fmt.Printf("写入路径错误：%s,错误路径为：%s", err, y)
		}
	}
	//}
	fmt.Println("文件写入完毕，即将关闭文件")
	return nil
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

/*
 * @MethodName writeToNovelTxt
 * @Description 将小说内部连接聚集到一起
 * @Author khr
 * @Date 2023/6/27 16:47
 */

func writeToNovelTxt(name string, urls []string) error {
	//文本存放地址及文本名称
	filePath := name + ".txt"
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		// 处理打开文件时的错误
		fmt.Println("无法打开和创建文件:", err)
		return err
	}
	//写入结束后关闭文件
	defer file.Close()
	// 文件已经存在或创建成功
	fmt.Println("文件已打开或创建成功！")
	for _, y := range urls {
		if !checkDuplicate(filePath, y) {
			_, err = file.WriteString(y + "\n")
			if err != nil {
				fmt.Printf("写入路径错误：%s,错误路径为：%s", err, y)
			}
		}
	}
	fmt.Println("文件写入完毕，即将关闭文件")
	return nil
}

/*
 * @MethodName
 * @Description 判断内容是否重复
 * @Author khr
 * @Date 2023/6/27 16:55
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
