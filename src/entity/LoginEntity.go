package entity

// LoginEntity 登录返回消息
type LoginEntity struct {
	UUID    string `json:"uuid"`
	Message string `json:"message"` // 错误消息
	Status  int    `json:"status"`  // 登陆状态：1-成功，0-失败
}

// GetPrefix 获取消息头
func (msg *LoginEntity) GetPrefix() string {
	return "_login"
}
