package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Name string      `json:"name"` //if we want the "Name" returned to be in lower case
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//upgrader switch the http to websocket
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(res, "Hello dear")
	// var socket *websocket.Conn
	socket, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println("this is the error last", err)
		}
	}
}
