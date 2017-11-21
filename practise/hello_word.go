package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fmtStdin() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f /  %d /  %s"
	)
	fmt.Println("请输入您的名字: ")

	// 在标准输出中,以空格为分隔,获取输入值,多余值则会忽略,以 `回车` 结束
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("您好 %s | %s!\n", firstName, lastName)

	// 注意input和format格式,不能有其它字符在
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("Sscanf 格式化后的输出为: ", f, i, s)

	fmt.Scanf("%d %s %f", &i, &s, &f)
	fmt.Println("Scanf 格式化后的输出为: ", i, s, f)

}

func bufioStdin() {
	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字: ")

	// 使用 `S` 字符来标识输入结束, 注意 `S` 只能使用引号
	input, err = inputReader.ReadString('S')

	// `abc`作为过滤字符, 输入中如果有 `abc`中任意一个,都会被过滤
	input = strings.Trim(input, "abc")
	if err == nil {
		fmt.Printf("您好: %s", input)
	}
}

func main() {
	fmtStdin()
	bufioStdin()
}
