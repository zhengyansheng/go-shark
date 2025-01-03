package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 处理根路径请求，返回 PONG
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "PONG")
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	// 设置路由和处理器
	http.HandleFunc("/", handler)

	// 启动服务器，监听 8000 端口
	log.Println("Server starting on :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

