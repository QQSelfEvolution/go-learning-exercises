// 小匠 Day1练习 #1: 变量与基础类型
// Go语言基础 - 变量声明和基础类型
package main

import "fmt"

func main() {
    fmt.Println("=== 小匠 Go语言学习 Day1 ===")
    
    // 变量声明方式
    var name string = "小匠"
    var age int = 24
    var height float64 = 175.5
    
    fmt.Printf("我叫%s，今年%d岁，身高%.1fcm\n", name, age, height)
    
    // 短变量声明
    city := "北京"
    isStudent := true
    
    fmt.Printf("所在城市: %s, 是学生: %v\n", city, isStudent)
    
    // 多个变量
    x, y := 10, 20
    sum := x + y
    product := x * y
    
    fmt.Printf("%d + %d = %d\n", x, y, sum)
    fmt.Printf("%d * %d = %d\n", x, y, product)
    
    // 常量
    const Pi = 3.14159
    const StatusOk = 200
    
    fmt.Printf("圆周率: %.5f, HTTP状态码: %d\n", Pi, StatusOk)
    
    fmt.Println("=== Day1练习1完成 ===")
}
