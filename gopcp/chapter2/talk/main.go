package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"plum/gopcp/chapter2/talk/chatbot"
)

var chatbotName string // 机器人的名字

func init() {
	// 命令行执行方法  go run main.go -chatbot simple.cn
	// 可以在命令行中指定语言
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "The chatbot's name for dialogue")
}

/**
检查程序中发生的错误，判断程序是否需要退出
return:
	是否有错误
*/
func checkError(chatbot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}

func main() {
	flag.Parse()

	// nil 可以表示任何指针类型和interface ,  就像表示空指针可以代表 *int , *string
	// 在实例化时传入的是 nil, 在使用过程中进行了判断, 不等于 nil 怎么样怎么样
	chatbot.Register(chatbot.NewSimpleEN("simple.en", nil))
	chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	myChatbot := chatbot.Get(chatbotName)
	if myChatbot == nil {
		err := fmt.Errorf("程序发生了致命错误， %s\n", chatbotName)
		checkError(nil, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := myChatbot.Begin()
	checkError(myChatbot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(myChatbot, err, true)
	fmt.Println(myChatbot.Hello(input[:len(input)-1]))
	for {
		input, err = inputReader.ReadString('\n')
		if checkError(myChatbot, err, false) {
			continue
		}
		output, end, err := myChatbot.Talk(input)
		if checkError(myChatbot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
		if end {
			err = myChatbot.End()
			checkError(myChatbot, err, false)
			os.Exit(0)
		}
	}
}
