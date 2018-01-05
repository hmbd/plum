package main

import "fmt"

type ByteSlice []byte

func Append(slice, data []byte) []byte {
	length := len(slice)
	if length+len(data) > cap(slice) {
		newSlice := make([]byte, (length+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0: length+len(data)]
	for i, c := range data {
		slice[length+i] = c
	}
	return slice
}

func (p *ByteSlice) Append1(data []byte) {
	slice := *p
	length := len(slice)
	if length+len(data) > cap(slice) {
		newSlice := make([]byte, (length+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0: length+len(data)]
	for i, c := range data {
		slice[length+i] = c
	}
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	length := len(slice)
	if length+len(data) > cap(slice) {
		newSlice := make([]byte, (length+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:length+len(data)]
	for i, c := range data {
		slice[length+i] = c
	}
	*p = slice
	return len(data), nil
}
func test1() {
	var a ByteSlice = []byte{1, 2, 3}
	b := []byte{4}
	Append(a, b)
	fmt.Println(a)
	a = Append(a, b)
	fmt.Println(a)
	c := Append(b, a)
	fmt.Println(c)
}
func test2() {
	var a ByteSlice = []byte{1, 2, 3}
	var b ByteSlice = []byte{1, 2, 3}
	c := []byte{4}
	(&a).Append1(c)
	b.Append1(c)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

func test3() {
	//type Writer interface{
	//	Write(p []byte)(n int, err error)
	//}
	var a ByteSlice = []byte{1, 2, 3, 4, 5}
	c := []byte{9}
	number, err := a.Write(c)
	if err != nil {
		fmt.Println("写入数据发生了错误")
	}
	fmt.Printf("字节数： %d, 数据： %v\n", number, a)
	var b ByteSlice
	fmt.Fprintf(&b, "aa%dbb", 7)
	fmt.Println(b)
}
func main() {
	test1()
	test2()
	test3()
}
