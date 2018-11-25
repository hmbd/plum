package main

import (
	"os"
	"fmt"
)

func main() {
	abort := make(chan int)
	go func() {
		input, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			input = 0
		}
		abort <- input
	}()

	fmt.Println("Conmmecing contdown. Press return to abort.")

	ch := make(chan int, 2)

	for i := 0; i < 10; i ++ {
		select {
		case ch <- i:

		case x := <-ch:
			fmt.Println(x)
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
