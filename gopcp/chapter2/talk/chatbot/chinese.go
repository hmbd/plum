package chatbot

import "fmt"
import "strings"

// 代表针对中文的聊天机器人
type simpleCN struct {
	name string
	talk Talk
}

// 创建针对中文聊天的机器人
func NewSimpleCN(name string, talk Talk) Chatbot {
	// simple 同时实现了接口 Talk 和 Chatbot 的方法  返回哪一类interface 都可以
	return &simpleCN{
		name: name,
		talk: talk,
	}
}

// Chatbot接口实现的一部分
func (robot *simpleCN) Name() string {
	return robot.name
}

// Chatbot接口实现的一部分
func (robot *simpleCN) Begin() (string, error) {
	return "请输入你的名字: ", nil
}

// Talk 接口实现的一部分
func (robot *simpleCN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("你好, %s!要继续操作吗?", userName)
}

// Talk 接口实现的一部分
func (robot *simpleCN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "退出", "再见":
		saying = "再见!"
		end = true
		return
	default:
		saying = "抱歉，请再次输入."
		return
	}
}

// Chatbot 接口实现的一部分
func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("程序发生了一个错误: %s\n", err)
}

// Chatbot 接口实现的一部分
func (robot *simpleCN) End() error {
	return nil
}
