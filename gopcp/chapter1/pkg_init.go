package main

import (
	"fmt"
	"runtime"
)

// 代码包初始化函数，在main之前运行，别的包可以单独调用
func init() {
	fmt.Printf("字典 Map: %v\n", m)
	// 通过系统内置包 runtime 获取当前机器的操作系统和计算结构
	// 查询到之后通过 Sprintf 方法进行格式化并复制给变量info 等价于 python 中的 "OS %s".format("Windows")
	info = fmt.Sprintf("操作系统OS: %s, 系统类型 Arch: %s", runtime.GOOS, runtime.GOARCH)
}

// 非局部变量，map 类型，且已初始化
var m = map[int]string{1: "A", 2: "B"}

// 非局部变量， string 类型，值进行了声明，位初始化
var info string

func main() {
	fmt.Println(info)
}
