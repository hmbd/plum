package chatbot

import "errors"

// 定义聊天的接口类型
type Talk interface{
	Hello(username string) string
	Talk(heard string)(saying string, end bool, err error)
}

// Chatbot 定义聊天机器人的接口类型
type Chatbot interface{
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	ErrInvalidChatbotName = errors.New("无效的机器人名称") // 代表无效的机器人名称错误
	ErrInvalidChatbot = errors.New("无效的机器人") // 代表无效的机器人错误
	ErrExistingChatbot = errors.New("机器人已经存在") // 代表同名机器人错误
)

// 代表名称---机器人的的映射
var chatbotMap = map[string]Chatbot{}

/**
注册聊天机器人
return:
	返回值就代表操作结果
*/
func Register(chatbot Chatbot) error{
	if chatbot == nil{
		return ErrInvalidChatbot
	}
	name := chatbot.Name()
	if name == ""{
		return ErrInvalidChatbotName
	}
	if _, ok := chatbotMap[name]; ok{
		return ErrExistingChatbot
	}
	chatbotMap[name] = chatbot
	return nil
}

/*
通过名字获取机器人
*/
func Get(name string) Chatbot{
	return chatbotMap[name]
}