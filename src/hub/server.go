package hub

import (
	"errors"
	"logic"
	"message"
	"reflect"
	"runtime/debug"
	"time"

	log "git.happyseven.cn/golibs/logger"

	"player"

	"fmt"

	"github.com/googollee/go-socket.io"
)

// hubCaches 注册的消息缓存
var hubCaches = make(map[string]interface{})

// Register 注册消息处理
func Register(name string, msg interface{}) error {
	var err error
	if hubCaches[name] == nil {
		hubCaches[name] = msg
	} else {
		err = errors.New("Message [" + name + "] duplicate register.")
	}
	return err
}

// StartServer 新开始一个消息处理服务
func StartServer(so socketio.Socket) {
	log.Println("client on connect")
	// 当前的用户实体
	var currentPlayer = player.NewPlayer()
	currentPlayer.ReConnect(so)
	var err error

	// 从Hub里读取注册的消息
	for k, m := range hubCaches {
		// 单个消息体进行反射新对象
		// 由于目前单个人的消息是走单线程，并且由于游戏机制，因此不存在同个消息同时并发的问题
		// 因此下面的这种方式在此情况下是可行的
		t := reflect.ValueOf(m).Type()
		h := reflect.New(t).Interface()
		hh := h.(logic.IMessage)
		// log.Printf("hh == m: %v\n", hh == m)
		func(key string, m logic.IMessage) {
			so.On(key, func(data string) {
				fmt.Printf("Message Come: %s\n", data)
				msg, success := m.Instance(data).(logic.IMessage)

				// log.Printf("msg == m: %v\n", msg == m)
				if success {
					if currentPlayer.GetID() > 0 {
						// 表示已经登陆
						err = currentPlayer.Accept(msg)
						if err != nil {
							if err.Error() == "" {
								message.SendSystemTip(currentPlayer, "system error")
							} else {
								message.SendSystemTip(currentPlayer, err.Error())
							}
							log.Println(err)
						}
					} else {
						if !msg.IsCheck() {
							// 未登陆状态
							msg.Execute(currentPlayer)
						} else {
							log.Println("not login:" + key)
						}
					}
				} else {
					log.Println(err)
				}
			})
		}(k, hh)
	}

	so.On("disconnection", func() {
		go func(uid int64) {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
					log.Printf("Panic: %v\n", err)
				}
			}()

			p := player.GetPlayer(uid)
			if p != nil {
				if !p.IsConnect() {
					log.Printf("玩家【%d】断开连接啦", uid)
					time.Sleep(10 * time.Second)
					p = player.GetPlayer(uid)
					if p != nil {
						if !p.IsConnect() {
							log.Printf("玩家【%d】掉线退出\n", uid)
							if p.GetSocket() != nil {
								p.GetSocket().Disconnect()
							}
							p.ReConnect(nil)
							// 清理缓存
							player.UnCachePlayer(p)
						}
					}
				}
			}
		}(currentPlayer.GetID())
	})
}
