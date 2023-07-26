package main

import (
	"fmt"
	"novel-co/click"
)

/**
 * @ClassName main
 * @Description TODO
 * @Author khr
 * @Date 2023/5/30 15:52
 * @Version 1.0
 */
func main() {
	qu, e := click.Ch.Prepare("select * from default.novel")
	fmt.Println(qu)
	fmt.Println(e, "e")
	//con := click.CliKhouseCon
	//defer con.Close()
	//con := click.CliKhouseCon
	//row, err := click.CliKhouseCon.Query("insert into novel (title,content,book_id) values (\"第一章\",\"jhiiiiiiiiiiiii\",22) ")
	//fmt.Println("effect rows", row)
	//panic(err)
	//fmt.Println("effect err", err)
	//// 创建一个新的Colly爬虫
	//c := colly.NewCollector()
	//
	//// 设置用户代理（可选）
	////c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
	//// 在请求之前设置一些配置，例如设置 User-Agent
	////c.OnRequest(func(r *colly.Request) {
	////	r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	////})
	//// 在请求之前执行的回调函数
	//c.OnRequest(func(r *colly.Request) {
	//	r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	//	fmt.Println("Visiting", r.URL.String())
	//})
	//
	//// 在HTML元素匹配的选择器上执行的回调函数
	//c.OnHTML("div.content", func(e *colly.HTMLElement) {
	//
	//	fmt.Println(e.Text)
	//	// 提取小说内容并保存到文件
	//	//err := saveToFile("novel.txt", e.Text)
	//	//if err != nil {
	//	//	log.Println("Failed to save novel:", err)
	//	//}
	//})
	//
	//// 启动爬虫并指定要访问的URL
	//err := c.Visit("https://www.biquge.co/0_760/291754.html") // 将 "xxx" 替换为你想要爬取的小说链接
	//if err != nil {
	//	log.Println("Failed to visit website:", err)
	//}
}

//// 保存内容到文件
//func saveToFile(filename string, content string) error {
//	file, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	_, err = file.WriteString(content)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Novel saved to", filename)
//	return nil
//}
