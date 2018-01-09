package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 开了 N 个 goroutine 同时访问每个url，在访问结束后把结果写到通道中，此时会阻塞住，等待从通道中读取内容
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("总耗时%.2fs\n", time.Since(startTime).Seconds())
}

func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s 发生了错误，错误信息：%v\n", url, err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("读取 %s 内容发生了错误，错误信息: %v\n", url, err)
		return
	}
	secs := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("耗时 %.2f 字节数: %7d , 网址： %s\n", secs, nbytes, url)
}
