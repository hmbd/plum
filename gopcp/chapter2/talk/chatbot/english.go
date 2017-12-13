package chatbot

import "fmt"
import "strings"

// 代表针对英文的聊天机器人
type simpleEN struct {
	name string
	talk Talk
}

// 创建针对英文聊天的机器人
func NewSimpleEN(name string, talk Talk) Chatbot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}

// Chatbot接口实现的一部分
func (robot *simpleEN) Name() string {
	return robot.name
}

// Chatbot接口实现的一部分
func (robot *simpleEN) Begin() (string, error) {
	return "Please input you name: ", nil
}

// Talk 接口实现的一部分
func (robot *simpleEN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("Hello, %s!What can id do for you?", userName)
}

// Talk 接口实现的一部分
func (robot *simpleEN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "exit", "bye":
		saying = "Bye!"
		end = true
		return
	default:
		saying = "Sorry, I didn't catch you."
		return
	}
}

// Chatbot 接口实现的一部分
func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("An error occurred: %s\n", err)
}

// Chatbot 接口实现的一部分
func (robot *simpleEN) End() error {
	return nil
}
