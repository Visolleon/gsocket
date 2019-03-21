package message

import (
	"encoding/json"
	"entity"
	"errors"
	"fmt"
	"log"
	"logic"
	"player"
)

// ChatMessage 聊天消息实体
type ChatMessage struct {
	ToID     int64  `json:"toId"`     // 发送给对方的ID
	NickName string `json:"nickname"` // 发送者昵称
	Content  string `json:"content"`  // 消息内容
}

// Instance 序列化
func (msg *ChatMessage) Instance(data string) interface{} {
	json.Unmarshal([]byte(data), msg)
	return msg
}

// IsCheck 本消息是否需要验证
func (msg *ChatMessage) IsCheck() bool {
	return false
}

// Execute 集成IMessage接口
func (msg *ChatMessage) Execute(p logic.IPlayer) error {

	log.Println(fmt.Sprintf("receive chat msg : %#v", msg))

	var err error

	var entity = &entity.ChatEntity{
		UID:      p.GetID(),
		Content:  msg.Content,
		NickName: msg.NickName,
	}

	// 接收者
	acceptUser := player.GetPlayer(msg.ToID)
	if acceptUser != nil {
		acceptUser.Send(entity)
	} else {
		// 发送失败，这返回错误信息
		err = errors.New("not found user")
	}

	return err
}
