package main

import "fmt"

func printHello(){
	fmt.Println("Hello from pringHello")
}

func main()  {
	go func() {
		fmt.Println("Hello inline")
	}()
	go printHello()

	fmt.Println("Hello from main")
}
