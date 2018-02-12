package main

import (
	"math"
	"fmt"
	"image/color"
)

type Point struct {
	x, y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint1 struct {
	*Point
	Color color.RGBA
}

type Path []Point

/*
包级别的函数
*/
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

/*
Point 类的方法
*/
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

/*
我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型(底层类型是指[]Point这个slice，Path就是命名类型)不是指针或者interface。
*/
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

func testColorPoint() {

	var cp ColoredPoint
	cp.x = 1
	fmt.Println(cp.Point.x)
	fmt.Println(cp.x)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 2}, red}
	var q = ColoredPoint{Point{3, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(3)
}

func testColorPoint1() {

	var cp ColoredPoint1
	cp.x = 1
	fmt.Println(cp.Point.x)
	fmt.Println(cp.x)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint1{&Point{1, 2}, red}
	var q = ColoredPoint1{&Point{3, 4}, blue}
	fmt.Println(p.Distance(*q.Point))
	p.ScaleBy(2)
	q.ScaleBy(3)
	fmt.Println(*p.Point, *q.Point)
}

func testPoint() {

	p := Point{x: 1, y: 2}
	q := Point{x: 4, y: 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(q.Distance(p))
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p1 := Point{1, 2}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println(p1)

	p2 := Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2)

	//不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到
	//Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
	Point{1, 2}.Distance(p2) //
}
func main() {
	testPoint()
	testColorPoint()
	testColorPoint1()
}
