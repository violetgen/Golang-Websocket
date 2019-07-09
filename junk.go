package main

import (
	"encoding/json"
	"fmt"

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

// type Speaker interface {
// 	Speak()
// }
// func SomeFunc(speaker Speaker) {
// 	speaker.Speak()
// }
// func (m Message) Speak() {
// 	fmt.Println("Im a " + m.Name + " event!")
// }

func main() {
	recRawMsg := []byte(`{"name": "channel add",` + `"data": {"name":"Hardware Support"}}`)

	var recMessage Message
	//decoding:
	//the second parameter is where the message will be decoded into
	//&recMessage returns a pointer to the value
	if err := json.Unmarshal(recRawMsg, &recMessage); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)

	if recMessage.Name == "channel add" {
		channel, err := addChannel(recMessage.Data)
		//sending the channel back to the web socket
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		// log.Fatalln(err)
		return channel, err
	}
	channel.Id = "1"
	fmt.Println("%#v\n", channel)
	return channel, nil
}
