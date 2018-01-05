package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

func newRect(x, y, width, height float64) *Rect{
	return &Rect{x, y, width, height}
}
func main() {
	b := new(bool)
	fmt.Println("bool: ", *b)
	i := new(int)
	fmt.Println("int: ", *i)
	s := new(string)
	fmt.Println("string: ", *s)

	rect1 := Rect{0, 1, 2, 3}
	rect2 := &Rect{width: 100, height: 200}
	fmt.Println(rect1, rect2)
}
