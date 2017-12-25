package main

import "fmt"
import "math"

// 接口是定义了类型一系列方法的列表
// 如果一个类型实现了该接口的所有方法，那么该类型就符合该接口
type Shape interface {
	area() float64
}

type Shape2 interface {
	Shape // 接口嵌入 和 Shape等价的
}

type Rectangles struct {
	height float64
	wight  float64
}

type Rectangles1 struct {
	Rectangles // 直接将 Rectangles 引入到 Rectangles1
}

type Circle struct {
	radius float64
}

type Circle1 struct {
	circle Circle // 定义 circle 类型为 Circle
}

func (r Rectangles) area() float64 {
	return r.height * r.wight
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea1(shape Shape) float64 {
	return shape.area()
}

func getArea2(shape Shape2) float64 {
	return shape.area()
}

// 指针方法， 会对原值有影响, 方法属于类型
func (circle *Circle) Bigger1() {
	circle.radius ++
}

// 值方法， 修改不会影响到原值， 方法属于类型对象本身
func (circle Circle) Bigger2() {
	circle.radius ++
}

func test() {
	r1 := Rectangles{20, 10}
	c1 := Circle{4}
	fmt.Println("长方形面积1： ", getArea1(r1), getArea2(r1))
	fmt.Println("圆面积1： ", getArea1(c1), getArea2(c1))

	r2 := Rectangles1{r1}
	c2 := Circle1{c1}
	// r2.height 、 r.wight 可以直接访问
	fmt.Println("长方形面积2： ", getArea1(r2), getArea2(r2))

	// 需要通过 c2.circle.radius 才能访问到半径
	fmt.Println("圆面积2： ", getArea1(c2.circle), getArea2(c2.circle))
}

func test2() {
	c := Circle{4}
	fmt.Println("圆面积3： ", getArea1(c))
	c.Bigger1()
	fmt.Println("圆面积4： ", getArea1(c))
	c.Bigger2()
	fmt.Println("圆面积5： ", getArea1(c))
}
func main() {
	test()
	test2()
}
