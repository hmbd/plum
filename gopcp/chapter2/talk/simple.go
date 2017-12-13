package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入您的名字：")
	// 直到输入的数据为 \n (即回车键) 时，输入结束
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("输入发生了错误: %s，程序结束\n", err)
		// 异常退出
		os.Exit(1)
	} else {
		// 把最后的 结束标识符(即 \n) 从数据中取出
		name := input[:len(input)-1]
		fmt.Printf("你好， %s! 需要继续吗?\n", name)
	}

	// 如果内部不 break 的话，就是个 死循环
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("输入发生了错误， %s", err)
			// 跳出本次循环
			continue
		}
		input = input[:len(input)-1]
		// 把数据转成小写
		switch  input {
		case "":
			// 输入为空
			continue
		case "exit", "bye":
			fmt.Println("你选择了退出")
			// 正常结束程序
			os.Exit(0)
		default:
			fmt.Println("抱歉，请再次输入")
		}
	}
}
