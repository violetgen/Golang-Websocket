package main

import (
	"github.com/gorilla/websocket"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send        chan Message //defining a channel that will habve a Message struct
	socket      *websocket.Conn
	findHandler FindHandler
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

// func (client *Client) subscribeChannels() {
// 	for {
// 		time.Sleep(r())
// 		client.send <- Message{"channel add", ""}
// 	}
// }

// func (client *Client) subscribeMessages() {
// 	for {
// 		time.Sleep(r())
// 		client.send <- Message{"message add", ""}
// 	}
// }

//a function that return between 0 to 1 second
// func r() time.Duration {
// 	return time.Millisecond * time.Duration(rand.Intn(1000))
// }

//This function create a new object and returns it
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}

// func main() {
// 	// msgChan := make(chan string)
// 	//defining a go routine that sends a string to a channel and the string is saved in a different variable and eventually printed, ie we removed it from the channel.
// 	//the channel is like a pipe that help manage different go routines
// 	// go func() {
// 	// 	//put stuff into the channel:
// 	// 	msgChan <- "Hello"
// 	// }()
// 	// //read from the channel and save what is the in the channel into a msg variable
// 	// msg := <-msgChan
// 	// fmt.Println(msg)

// 	client := NewClient()
// 	go client.subscribeChannels()
// 	go client.subscribeMessages()
// 	client.write() //this use the go routine created when main is called
// }
