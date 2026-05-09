// 阿代码 Day1练习 #3: 文件读写
// Go语言基础 - 文件操作
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== 阿代码 Go语言学习 Day1 - 文件读写 ===")
	
	filename := "day1_practice.txt"
	
	// 写入文件
	content := []string{
		"=== 阿代码 Go语言练习 ===",
		fmt.Sprintf("时间: %s", time.Now().Format("2006-01-02 15:04:05")),
		"Day1练习: 文件读写",
		"学习目标: 掌握Go语言文件操作",
		"",
	}
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()
	
	for _, line := range content {
		fmt.Fprintln(file, line)
	}
	fmt.Printf("已写入文件: %s\n", filename)
	
	// 读取文件
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()
	
	fmt.Println("\n--- 文件内容 ---")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 高亮显示标题
		if strings.HasPrefix(line, "===") {
			fmt.Printf("\n\033[33m%s\033[0m\n", line)
		} else {
			fmt.Println(line)
		}
	}
	
	fmt.Println("\n=== Day1练习3完成 ===")
}
