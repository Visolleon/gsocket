# gsocket

管理WebSocket消息体的服务框架，很好的规范上报消息和下发消息的结构和管理，避免杂乱无章的消息头和消息体，多用于一些socket广播消息较多的场景，如：游戏、多人多频道聊天等；支持socket.io和websocket


* 项目目录说明

    ```
    ├─docs           # 开发文档
    ├─log            # 运行日志
    ├─scripts        # 发布启动脚本
    │  └─init
    │      └─centos
    ├─src        # 源码
    │  ├─config         # 游戏配置文件
    │  ├─entity         # 下发消息体
    │  ├─hub            # 消息注册
    │  ├─message        # 上报消息体处理
    │  ├─player         # 用户信息在线缓存
    │  ├─utils          # 通用库
    │  ├─web            # WEB
    └─main.go  # 主程序入口
    ```

* 通讯上报和下发处理
    - 上报消息：在message文件夹中定义Message，定义的类需继承IMessage接口，定义完成后，需要hub中注册消息


    ```go
    // ChatMessage 聊天消息实体
    type ChatMessage struct {
        ToID    string `json:"toId"`    // 接收方的ID
        Content string `json:"content"` // 消息内容
    }

    // Instance 序列化
    func (msg *ChatMessage) Instance(data string) interface{} {
        json.Unmarshal([]byte(data), msg)
        return msg
    }

    // IsCheck 本消息是否需要验证
    func (msg *ChatMessage) IsCheck() bool {
        return true
    }

    // Execute 集成IMessage接口
    func (msg *ChatMessage) Execute(p logic.IPlayer) error {
        // todo something
        return nil
    }

    ```


    - 下发消息：在entity文件夹中定义entity消息，需要继承IEntity接口，为统一接口，因此每个下发的entity都需要硬编码定义消息头，防止后续版本API混乱。


    ```go
    // ChatEntity 聊天发送消息
    type ChatEntity struct {
        PlayerID   string `json:"id"`      // 发送者ID
        PlayerName string `json:"name"`    // 发送者昵称
        Content    string `json:"content"` // 消息内容
    }

    // GetPrefix 获取消息头
    func (msg *ChatEntity) GetPrefix() string {
        return "chat"
    }
    ```


## 部署说明-Linux(CentOS)

	
* 创建 gserveruser 用户
    ```sh
    $useradd -s /sbin/nologin gserveruser
    ```
	
* 源码编译(开发模式, 部署请忽略)
	- Windows Build To Linux64:

			$build_linux64.bat

	
* 配置启动脚本(CentOS):
    ```sh
    $chown -R gserveruser:gserveruser /home/gsocket
    $chmod 755 scripts/init/centos/gsocket
    $cp scripts/init/centos/gsocket /etc/rc.d/init.d/
    $chkconfig --add gsocket
    ```
			
* 加 `读写` 权限
    ```sh
    $chmod 755 gsocket
    ```
			
* 运行服务
    ```sh
    $service gsocket start
    ```