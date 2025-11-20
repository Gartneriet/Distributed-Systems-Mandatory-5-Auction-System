package main

import (
	"Mandatory_5_-_Auction_System/client"
	"Mandatory_5_-_Auction_System/server"
	"flag"
	"log"
	"time"
)

func main() {
	nodeType := flag.String("type", "", "type of node")
	name := flag.String("name", "", "Name of the server")
	port := flag.String("port", "", "Port of the server")
	flag.Parse()
	if nodeType != nil && *nodeType == "client" {
		startClient(*name)
	} else if nodeType != nil && *nodeType == "server" {
		startServer(*port, *name)
	}
}

func startServer(port string, name string) {
	log.Printf("server listening on port %s", port)
	_ = server.StartServer(port, name)
}
func startClient(name string) {
	log.Printf("client listening on %s", name)
	time.Sleep(2 * time.Second)

	c := client.NewClient(name)
	c.Run()
}
