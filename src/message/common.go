package message

import "logic"

// SendSystemTip 发送系统提示
func SendSystemTip(currentPlayer logic.IPlayer, tip string) {
	currentPlayer.Emit("systemtip", tip)
}
