package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"plum/practise/server" // server 是包名
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", gifimage)
	http.HandleFunc("/count", counter)
	const hostPort = "localhost:8080"
	fmt.Println("请在浏览器中访问： ", "http://"+hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
func gifimage(w http.ResponseWriter, r *http.Request) {
	server.Lissajous(w) // Lissajous 是函数名，不需要文件名，通过函数名调用，函数名注意不要和文件名同名
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "请求类型： %s, URL： %s， 协议：%s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "访问的主机ip: %q\n", r.Host)
	fmt.Fprintf(w, "远程地址： %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "参数 [%q] = %q\n", k, v)
	}
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w, "你输入的URL为: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "你访问次数为: %d\n", count)
	mu.Unlock()
}
