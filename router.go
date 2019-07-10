package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Router struct {
	rules map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

func (e *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	socket, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		// fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err.Error())
		return
	}
	client := NewClient(socket, e.FindHandler)
	go client.Write()
	client.Read()

}

// func addChannel(data interface{}) error {
// 	var channel Channel

// 	err := mapstructure.Decode(data, &channel)
// 	if err != nil {
// 		return err
// 	}
// 	channel.Id = "1"
// 	// fmt.Println("%#v\n", channel)
// 	fmt.Println("added channel")
// 	return nil
// }

// func subscribeChannel(socket *websocket.Conn) {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		message := Message{
// 			"channel add",
// 			Channel{
// 				"1", "Software Support",
// 			},
// 		}
// 		socket.WriteJSON(message)
// 		fmt.Println("sent new channel")
// 	}
// }
