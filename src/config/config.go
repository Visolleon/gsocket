package config

import (
	"log"

	"git.happyseven.cn/golibs/utils"

	"gopkg.in/ini.v1"
)

var (
	//Cfg default ini
	Cfg *ini.File

	// Port socket listen port
	Port int

	// MaxPlayerCount 最大游戏人数
	MaxPlayerCount int
)

// Init 初始化配置文件
func Init() {
	loadConfigIni()

	// 监控Config文件，全部重新加载
	go utils.Watcher("config.ini", func(fileName string) {
		log.Println("config.ini 发生变化，全部重新加载")
		loadConfigIni()
	})
}

// LoadConfigIni 载入配置文件
func loadConfigIni() {
	log.Println("init config")
	var err error
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Printf("Setting init load config.ini error, %s", err)
		// panic(err)
		return
	}
	Port = Cfg.Section("server").Key("PORT").MustInt(8080)

	MaxPlayerCount = Cfg.Section("game").Key("MAXPLAYER").MustInt(2)
}
