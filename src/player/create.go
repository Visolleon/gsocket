package player

import "sync"

// 用户信息缓存,可根据房间进行缓存
var pCache = &playerCache{
	caches: make(map[int64]*Player),
	lock:   new(sync.Mutex),
}

// GetPlayer 根据ID获取用户信息
func GetPlayer(id int64) *Player {
	return pCache.GetPlayer(id)
}

// CachePlayer 用户登陆完成缓存用户信息
func CachePlayer(p *Player) {
	pCache.Cache(p)
}

// UnCachePlayer 用户退出清理缓存
func UnCachePlayer(p *Player) {
	pCache.UnCache(p)
}

// ExecuteByRPCServerID 分用户是哪个区来执行相关操作
func ExecuteByRPCServerID(serverID string, handler ExecuteHandler) error {
	return pCache.ExecuteByRPCServerID(serverID, handler)
}

// ExecuteByTeamID 分用户是哪个队伍来执行相关操作
func ExecuteAll(handler ExecuteHandler) error {
	return pCache.ExecuteAll(handler)
}
