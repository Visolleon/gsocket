package entity

// ErrorEntity 通用错误消息
type ErrorEntity struct {
	Prefix  string `json:"-"`       // 消息头
	Status  int    `json:"status"`  // 错误状态
	Message string `json:"message"` // 错误信息
}

// GetPrefix 获取消息头
func (msg *ErrorEntity) GetPrefix() string {
	if msg.Prefix == "" {
		msg.Prefix = "_error"
	}
	return msg.Prefix
}
