package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 生成 0,1,2,3... 这样的序列 存入 channel中
	go func() {
		for x := 0; x <= 100; x ++ {
			naturals <- x
		}
		close(naturals)
	}()

	// 从channel中取值收计算平台存入另一个channel
	//go func() {
	//	for {
	//		x := <-naturals
	//		squares <- x * x
	//	}
	//}()

	// 没有办法直接测试一个channel是否被关闭，
	// 但是接收操作有一个变体形式：它多接收一个结果，多接收的第二个结果是一个布尔值ok，
	// true  表示成功从channels接收到值，
	// false 表示channels已经被关闭并且里面没有值可接收。
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//for {
	//	fmt.Println(<-squares)
	//}

	// range 可以处理channel没有值时退出循环，和python中的for循环可以处理 StopIteration 一样
	for x := range squares {
		fmt.Println(x)
	}
}
