package utils

var (
	// Version 版本号, go build -ldflags "-X utils.Version '20170420'"
	Version = "20170926"

	// DebugInfo 是否Debug模式, 正式发布：go build -v -ldflags "-X utils.DebugInfo='false'"
	DebugInfo = "true"
)

// IsDebug 是否Debug模式
func IsDebug() bool {
	return DebugInfo == "true"
}
