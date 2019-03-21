package logic

// IMessage 消息接口
// 所有来自于客户端消息都将即成此接口
type IMessage interface {

	// 消息转化抽象
	Instance(json string) interface{}

	// 本消息是否需要验证登录
	IsCheck() bool

	// 消息执行抽象
	Execute(p IPlayer) error
}
