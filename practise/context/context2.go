package main

import "fmt"

func printHello(ch chan int){
	fmt.Println("Hello from printHello")
	ch <- 2
}

func main(){
	ch := make(chan int)
	go func() {
		fmt.Println("Hello inline")
		ch <- 1
	}()

	go printHello(ch)
	fmt.Println("Hello from main")

	// channel 无缓存，当存入值后会导致 夯住，直到把值取出后
	i := <-ch
	fmt.Println("Recieved1 ", i)

	j := <-ch
	fmt.Println("Recieved2 ", j)
}
