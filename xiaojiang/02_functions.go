// 小匠 Day1练习 #2: 函数定义
// Go语言基础 - 函数定义与多返回值
package main

import (
	"errors"
	"fmt"
)

/// 计算器函数
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

// 安全除法 - 返回结果和错误
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 带命名的返回值
func swap(a, b int) (int, int) {
	return b, a
}

// 可变参数函数
func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Println("=== 小匠 Go语言学习 Day1 - 函数 ===")
	
	// 基础函数调用
	a, b := 15, 4
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	
	// 安全除法
	result, err := divide(a, b)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, result)
	}
	
	// 测试除零
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("捕获除零错误: %v\n", err)
	}
	
	// 交换变量
	x, y := 100, 200
	fmt.Printf("交换前: x=%d, y=%d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("交换后: x=%d, y=%d\n", x, y)
	
	// 可变参数
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("sum(1,2,3,4,5) = %d\n", sum)
	
	sum2 := sumAll(10, 20, 30)
	fmt.Printf("sum(10,20,30) = %d\n", sum2)
	
	fmt.Println("=== Day1练习2完成 ===")
}
