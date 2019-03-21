package web

import (
	"config"
	"fmt"
	"hub"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

// Init web路由
func Init() {
	// SocketIO服务
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.SetMaxConnection(100000)
	server.On("connection", hub.StartServer)
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("Error:", err)
	})

	http.Handle("/socket.io/", server)

	http.Handle("/", http.FileServer(http.Dir("statics")))

	log.Println(fmt.Sprintf("Server start at port: %d", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
