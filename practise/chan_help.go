package main

import (
	"fmt"
	"time"
)

// channel的作用就是在多线程中传递数据
// channel 是先进先出的，如果给channel赋值了，那么就必须多去它的值，不然就会造成阻塞(因为赋值了，导致阻塞，也就无法读取它的值)，这个只对无缓冲的channel有效
// 对于有缓冲的channel，发送方会一直阻塞直到数据被拷贝到缓冲区，如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复

func test1() {
	ch := make(chan int)

	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	ch <- 1 // 把1存入,此时会阻塞，直到值被取出
	fmt.Println(2)
}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // 把i存入channel
		fmt.Println("发送数据: ", i)
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println("接收到数据： ", value)
	}
}

func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	// 等待把值取完,只能存一个，取一个
	time.Sleep(1 * time.Second)

	ch1 := make(chan int, 10)
	go produce(ch1)
	go consumer(ch1)
	// 等待把值取完，可以同时存十个，取十个
	time.Sleep(1 * time.Second)
}
