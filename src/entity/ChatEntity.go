package entity

// ChatEntity 聊天发送消息
type ChatEntity struct {
	UID      int64  `json:"uid"`      // 发送者UUID
	NickName string `json:"nickname"` // 发送者昵称
	Content  string `json:"content"`  // 消息内容
}

// GetPrefix 获取消息头
func (msg *ChatEntity) GetPrefix() string {
	return "_chat"
}
