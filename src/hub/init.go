package hub

import (
	"message"
)

// Init 初始化消息处理
func Init() {
	// 用户相关
	Register("login", message.LoginMessage{})

	// 聊天
	Register("chat", message.ChatMessage{})
}
