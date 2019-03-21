package player

import (
	"errors"
	"sync"

	log "git.happyseven.cn/golibs/logger"
)

// playerCache 玩家信息缓存
type playerCache struct {
	caches map[int64]*Player
	lock   *sync.Mutex
}

// Cache 缓存用户信息
func (pc *playerCache) Cache(player *Player) error {
	pc.lock.Lock()
	defer pc.lock.Unlock()

	var err error

	if pc.caches[player.GetID()] == nil {

	} else {
		err = errors.New("缓存已经存在，请检测是否是重复登陆")
	}
	pc.caches[player.GetID()] = player
	log.Printf("当前在线数： %d\n", len(pc.caches))
	return err
}

// UnCache 取消缓存
func (pc *playerCache) UnCache(player *Player) {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	log.Printf("清理用户为： %d\n", player.GetID())

	delete(pc.caches, player.GetID())
	log.Printf("剩下当前在线数： %d\n", len(pc.caches))
}

// GetPlayer 获取在线用户 (调用频率较高,也无需加锁)
func (pc *playerCache) GetPlayer(id int64) *Player {
	return pc.caches[id]
}

// ExecuteHandler 执行Handler
type ExecuteHandler func(p *Player) error

// ExecuteByRPCServerID 分用户是哪个区来执行相关操作
func (pc *playerCache) ExecuteByRPCServerID(gameServerID string, handler ExecuteHandler) error {
	var err error
	for _, v := range pc.caches {
		if v.GetRpcServerID() == gameServerID {
			err = handler(v)
		}
	}
	return err
}

// ExecuteAll 处理所有人
func (pc *playerCache) ExecuteAll(handler ExecuteHandler) error {
	var err error
	for _, v := range pc.caches {
		err = handler(v)
	}
	return err
}
