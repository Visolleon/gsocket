package logic

import "github.com/googollee/go-socket.io"

// IPlayer 玩家实体的Interface:
type IPlayer interface {
	IsReal() bool // 是否真实用户

	GetID() int64 // 获取ID
	SetID(id int64)
	GetToken() string
	SetToken(string)

	GetRpcServerID() string
	SetRpcServerID(string)

	IsConnect() bool           // 获取是否还在连接中
	ReConnect(socketio.Socket) // 重新连接

	Send(IEntity) error             // Socket发送消息
	Emit(string, interface{}) error // 发送自定义消息 (一般为通用通知消息等)
	Accept(IMessage) error          // 收到消息处理消息
}
