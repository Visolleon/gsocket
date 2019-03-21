package entity

// CommonEntity 通用信息返回
type CommonEntity struct {
	Prefix  string `json:"-"` // 消息头
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

// GetPrefix 获取通用返回消息头
func (msg *CommonEntity) GetPrefix() string {
	if msg.Prefix == "" {
		msg.Prefix = "_common"
	}
	return msg.Prefix
}
