package player

import (
	"logic"
	"sync"

	"github.com/googollee/go-socket.io"
)

// Player 玩家信息
type Player struct {
	ID    int64  `json:"id"` // 玩家ID
	Token string // 唯一验证的Token

	RpcServerID string // 每个用户对应服务器信息

	lock   *sync.Mutex     // 锁
	socket socketio.Socket // 用户的Socket对象
}

// IsReal 是否是真实用户
func (p *Player) IsReal() bool {
	return true
}

// GetID 获取ID
func (p *Player) GetID() int64 {
	return p.ID
}

// SetID 设置ID
func (p *Player) SetID(id int64) {
	p.ID = id
}

// GetToken 获取用户Token
func (p *Player) GetToken() string {
	return p.Token
}

// SetToken 设置用户Token
func (p *Player) SetToken(token string) {
	p.Token = token
}

// GetRpcServerID 获取用户服务器信息
func (p *Player) GetRpcServerID() string {
	return p.RpcServerID
}

// SetRpcServerID 设置用户服务器信息
func (p *Player) SetRpcServerID(sid string) {
	p.RpcServerID = sid
}

// IsConnect 获取是否还在连接中
func (p *Player) IsConnect() bool {
	if p.socket == nil {
		return false
	}
	err := p.socket.Emit("ping", "")
	return err == nil
	//return p.socket != nil
}

// ReConnect 重新连接
func (p *Player) ReConnect(io socketio.Socket) {
	p.socket = io
}

// GetSocket 获取当前连接
func (p *Player) GetSocket() socketio.Socket {
	return p.socket
}

// Send Socket发送消息
func (p *Player) Send(entity logic.IEntity) error {
	err := p.Emit(entity.GetPrefix(), entity)
	return err
}

// Emit 自定义发送消息 (一般为通用通知消息等)
func (p *Player) Emit(prefix string, entity interface{}) error {
	// 排队发送消息，防止并发消息产生问题
	p.lock.Lock()
	defer p.lock.Unlock()

	err := p.socket.Emit(prefix, entity)
	return err
}

// Accept 接受消除并处理
func (p *Player) Accept(msg logic.IMessage) error {
	err := msg.Execute(p)
	return err
}

// NewPlayer 新创建玩家对象
func NewPlayer() *Player {
	return &Player{
		lock: new(sync.Mutex),
	}
}
