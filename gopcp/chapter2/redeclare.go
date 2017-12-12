package main

import "fmt"

var v = "1, 2, 3" // 可以不显示声明类型，自动识别

func main()  {
	v := []int{1,2,3} // 声明并初始化一个数组 声明时数组长度可省略
	if v != nil{
		var v = 123
		fmt.Printf("v: %v\n", v)
	}
}

