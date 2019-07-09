package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
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
		// msgType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)

		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			go subscribeChannel(socket)
		}
		// fmt.Println(string(msg))
		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println("this is the error last", err)
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	// fmt.Println("%#v\n", channel)
	fmt.Println("added channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)
		message := Message{
			"channel add",
			Channel{
				"1", "Software Support",
			},
		}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}
