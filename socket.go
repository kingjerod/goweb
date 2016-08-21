package socket

import (
	"log"
	"github.com/googollee/go-socket.io"
)

func Handle() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(socket socketio.Socket) {
		log.Println("User connected")
		socket.Join("chat")
		socket.On("chat message", func(msg string) {
			log.Println("emit:" + msg)
			socket.Emit("chat message", msg)
			socket.BroadcastTo("chat", "chat message", msg)
		})
		socket.On("disconnection", func() {
			log.Println("User disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	return server;
}
