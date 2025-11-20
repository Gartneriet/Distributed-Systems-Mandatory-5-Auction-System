package client

import (
	proto "Mandatory_5_-_Auction_System/grpc"
	"context"
	"log"
	"math/rand/v2"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	name                string
	clock               int32
	primaryServerPort   string
	secondaryServerPort string
	primary             bool
}

func NewClient(name string) *client {
	log.Printf("Creating new client with name: %s", name)
	return &client{name: name, clock: 0, primaryServerPort: "5050", secondaryServerPort: "6000", primary: true}
}

func (this *client) Run() {
	for {
		time.Sleep(2 * time.Second)
		log.Printf("Created client")
		conn := this.getConnection(this.primary)

		c := proto.NewAuctionClient(conn)
		log.Printf("Fetch current highest bid from server at time: %v", this.clock)

		// Call query
		result, err := c.Query(context.Background(), &proto.Empty{})
		if err != nil {
			// Error querying primary server. Switch to secondary
			log.Printf("error querying auction: %s", err)
			this.primary = false
			continue
		}

		// Increment clock
		this.clock = max(this.clock, result.Timestamp) + 1

		// If auction is over don't bid
		if result.Status == "over" {
			continue
		}

		oldBid := result.Bid

		// Send a bid
		if rand.IntN(100) > 37 {
			// Found out how much to raise
			bid := oldBid * (rand.Float32()*rand.Float32() + 1.1)
			log.Printf("Bidding: %v at time: %v", bid, this.clock)
			_, err := c.Bidding(context.Background(), &proto.Bid{
				Author:    this.name,
				Bid:       bid,
				Timestamp: this.clock,
			})
			if err != nil {
				log.Printf("error querying bidding: %s", err)
			}
			this.clock++
		}
	}
}

func (this *client) getConnection(port bool) *grpc.ClientConn {
	if port {
		conn, err := grpc.NewClient("localhost:"+this.primaryServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Error creating new client: %s", err)
		}
		return conn
	} else {
		conn, err := grpc.NewClient("localhost:"+this.secondaryServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Error creating new client: %s", err)
		}
		return conn
	}
}
