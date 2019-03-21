package logic

// IEntity 发送到客户端消息接口
// 所有发送消息都将即成此接口
type IEntity interface {
	GetPrefix() string // 获取消息前缀
}
