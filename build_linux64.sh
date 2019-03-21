
#!/bin/sh
# 如果有权限问题：可使用sudo -E ./build_linux64.sh
export fullstamp=$(date +%Y%m%d%H%M)
echo $GOPATH
echo '------------go env---------------'
export GOROOT= # Go语言的安装目录的绝对路径
export GOTOOLDIR= # Go工具目录的绝对路径
export GOPATH=$GOPATH:`pwd` # 工作区目录的绝对路径
export GOARCH=amd64 # 构建环境的目标计算架构
export GOBIN= # 存放可执行文件的目录的绝对路径
# export GOCHAR= # 程序构建环境的目标计算架构的单字符标识
export GOEXE=exe # 可执行文件的后缀
export GOHOSTARCH=amd64 # 程序运行环境的目标计算架构
export GOOS=linux # 程序构建环境的目标操作系统
export GOHOSTOS=centos # 程序运行环境的目标操作系统
export GORACE= # 用于数据竞争检测的相关选项

# 非Plan 9操作系统的环境信息
export CC=gcc # 操作系统默认的C语言编译器的命令名称
# export GOGCCFLAGS=-fPIC -m64 -fmessage-length=0 # Go语言在使用操作系统的默认C语言编译器对C语言代码进行编译时加入的参数

export CXX=g++
export CGO_ENABLED=0  # cgo工具是否可用
go env

echo '\033[34m-----------go get----------------\033[0m'
go get -u -v

echo '\033[34m-----------go build----------------\033[0m'
go clean
go build -v -i -ldflags "-s -w -X main.Version='$fullstamp'"

echo '\033[32m┌────────────────────┐\033[0m'
echo '\033[32m│  Build completed!  │\033[0m'
echo '\033[32m└────────────────────┘\033[0m'