package main

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"time"
)

func main() {
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewAnonymousClient() //twitch.NewClient("alomerry", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	client.Join("alomerry")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Raw)
	})

	time.Sleep(time.Second * 30)
}
