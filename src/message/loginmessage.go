package message

import (
	"encoding/json"
	"entity"
	"fmt"
	"logic"
	"player"
	"utils"
)

// LoginMessage 登陆接收的消息
type LoginMessage struct {
	UserID int64  `json:"id"`
	Token  string `json:"token"`
}

// Instance 序列化
func (msg *LoginMessage) Instance(data string) interface{} {
	json.Unmarshal([]byte(data), msg)
	return msg
}

// IsCheck 本消息是否需要验证
func (msg *LoginMessage) IsCheck() bool {
	return false
}

// Execute 处理登陆相关信息处理
func (msg *LoginMessage) Execute(currentPlayer logic.IPlayer) error {
	var err error

	fmt.Println("Login::: ID:", msg.UserID, "Token:", msg.Token)

	ret := &entity.LoginEntity{
		Status: 0,
	}

	// NOTE: 测试用
	if msg.Token == "TokenTest" && utils.IsDebug() {
		currentPlayer.SetID(msg.UserID)
		currentPlayer.SetToken(msg.Token)

		checkReconnect(currentPlayer)

		ret.Status = 1
		err = currentPlayer.Emit("login", ret)
	} else {
		ret.Message = "login fail"
		ret.Status = -1
	}
	err = currentPlayer.Emit("_login", ret)
	return err
}

// checkReconnect 检测是否是重新链接
func checkReconnect(currentPlayer logic.IPlayer) {
	cp := currentPlayer.(*player.Player)

	// 判断用户是否存在
	cacheUser := player.GetPlayer(currentPlayer.GetID())
	if cacheUser != nil {
		// 让原有的socket断开
		if cacheUser.IsConnect() {
			// 向原来的socket发送取代消息
			cacheUser.GetSocket().Emit("_system_kick_out", "")
		}
		// 原有用户对象链接上现在的socket
		cacheUser.ReConnect(cp.GetSocket())
	}

	// 加入在线缓存
	player.CachePlayer(cp)
}
