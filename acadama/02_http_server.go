// 阿代码 Day1练习 #2: HTTP Server
// Go语言基础 - 使用net/http创建简单服务器
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Response 返回结构
type Response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

// rootHandler 根路径处理
func rootHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message:   "阿代码的Go HTTP服务器运行中!",
		Timestamp: time.Now().Unix(),
		Status:    "ok",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// healthHandler 健康检查
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health: OK\nTimestamp: %d", time.Now().Unix())
}

func main() {
	fmt.Println("=== 阿代码 Go语言学习 Day1 - HTTP Server ===")
	
	// 注册路由
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	
	port := ":8080
	fmt.Printf("服务器启动在 http://localhost%s\n", port)
	fmt.Println("按Ctrl+C停止服务器")
	
	// 启动服务器
	log.Fatal(http.ListenAndServe(port, nil))
}
