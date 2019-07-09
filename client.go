package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send chan Message //defining a channel that will habve a Message struct
}

func (client *Client) write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
	}
}

func (client *Client) subscribeChannels() {
	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages() {
	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

//a function that return between 0 to 1 second
func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

//This function create a new object and returns it
func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}

func main() {
	// msgChan := make(chan string)
	//defining a go routine that sends a string to a channel and the string is saved in a different variable and eventually printed, ie we removed it from the channel.
	//the channel is like a pipe that help manage different go routines
	// go func() {
	// 	//put stuff into the channel:
	// 	msgChan <- "Hello"
	// }()
	// //read from the channel and save what is the in the channel into a msg variable
	// msg := <-msgChan
	// fmt.Println(msg)

	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}
