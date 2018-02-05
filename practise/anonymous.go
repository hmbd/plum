package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	f1 := func(x int) int{
		x += 2
		return x
	}
	fmt.Println(f1(1)) // "3"
	fmt.Println(f1(3)) // "5"
}
