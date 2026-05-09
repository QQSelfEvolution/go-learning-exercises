// 小匠 Day1练习 #3: 结构体
// Go语言基础 - 结构体定义与方法
package main

import (
	"fmt"
	"time"
)

// Person 结构体 - 人员信息
type Person struct {
	Name   string
	Age    int
	Email  string
}

// Student 结构体 - 继承Person
type Student struct {
	Person       // 嵌入结构体
	StudentID    string
	Grade        int
	EnrolledDate time.Time
}

// Teacher 结构体
type Teacher struct {
	Name      string
	Subject   string
	Students  []*Student // 教的学生
}

// NewStudent 构造函数
func NewStudent(name string, age int, email string, studentID string, grade int) *Student {
	return &Student{
		Person: Person{
			Name:  name,
			Age:   age,
			Email: email,
		},
		StudentID:    studentID,
		Grade:        grade,
		EnrolledDate: time.Now(),
	}
}

// Greet 方法 - 打招呼
func (p Person) Greet() string {
	return fmt.Sprintf("你好，我是%s，今年%d岁", p.Name, p.Age)
}

// Study 方法
func (s *Student) Study() string {
	return fmt.Sprintf("%s正在读%d年级", s.Name, s.Grade)
}

// Teach 方法
func (t *Teacher) Teach() string {
	return fmt.Sprintf("%s老师正在教%s", t.Name, t.Subject)
}

func main() {
	fmt.Println("=== 小匠 Go语言学习 Day1 - 结构体 ===")
	
	// 创建学生
	student1 := NewStudent("张三", 18, "zhangsan@example.com", "S001", 3)
	student2 := &Student{
		Person: Person{Name: "李四", Age: 19, Email: "lisi@example.com"},
		StudentID: "S002",
		Grade: 4,
		EnrolledDate: time.Now(),
	}
	
	// 调用方法
	fmt.Println(student1.Greet())
	fmt.Println(student1.Study())
	
	fmt.Println(student2.Greet())
	fmt.Println(student2.Study())
	
	// 创建老师
	teacher := &Teacher{
		Name:    "王老师",
		Subject: "Go语言",
		Students: []*Student{student1, student2},
	}
	
	fmt.Println(teacher.Teach())
	fmt.Printf("王老师教%d个学生\n", len(teacher.Students))
	
	// 显示学生信息
	fmt.Println("\n--- 学生列表 ---")
	for _, s := range teacher.Students {
		fmt.Printf("- %s (ID: %s, 年级: %d)\n", s.Name, s.StudentID, s.Grade)
	}
	
	fmt.Println("=== Day1练习3完成 ===")
}
