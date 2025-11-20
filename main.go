package main

import (
	"Mandatory_5_-_Auction_System/client"
	"Mandatory_5_-_Auction_System/server"
	"flag"
	"time"
)

func main() {
	nodeType := flag.String("type", "", "Path to node YAML config")
	if *nodeType == "client" {
		startClient()
	} else if *nodeType == "server" {
		startServer()
	}
}

func startServer() {
	_ = server.StartServer()
}
func startClient() {
	time.Sleep(5 * time.Second)

	c := client.NewClient()
	c.Run()
}
