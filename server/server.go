package server

import (
	proto "Mandatory_5_-_Auction_System/grpc"
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type auctionServer struct {
	proto.UnimplementedAuctionServer
	mu         sync.Mutex
	clock      int32
	highestBid float32
	status     string
	name       string
}

func newServer(name string) *auctionServer {
	return &auctionServer{
		clock:      0,
		highestBid: 10,
		name:       name,
	}
}

func StartServer(port string, name string) error {

	log.Printf("Starting server on port %s", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", port, err)
	}

	server := newServer(name)
	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Server listening on %s", port)
	return nil

}

func (this *auctionServer) Bidding(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	// Lock and unlock
	this.mu.Lock()
	defer this.mu.Unlock()

	// Log receive
	log.Printf("Received bid: %v at time: %v", msg.Bid, msg.Timestamp)
	log.Printf("Current bid: %v", this.highestBid)

	// End auction depending on clock
	if this.clock > 100 {
		this.status = "over"
		log.Printf("SOLD! FOR %v DOLLARS!", this.highestBid)
	}

	// Time is up - no more bids
	if this.status == "over" {
		log.Printf("Auction over. No more bidding is allowed. Sold for %v", this.highestBid)
		return &proto.Ack{}, nil
	}

	// Increment clock
	this.clock = max(this.clock, msg.Timestamp) + 1

	// Compare to current highest bid
	// If higher than 110%, then update the highest bid and Ack
	if msg.Bid > (1.1 * this.highestBid) {
		this.highestBid = msg.Bid
		log.Printf("New highest bid: %v", this.highestBid)
	}

	if this.name == "server1" {
		// Simulated crashing
		if rand.Float32()*rand.Float32() > 0.5 {
			crashServer(this)
		}
		// Backup
		log.Printf("Performing backup...")
		this.Backup()
		log.Printf("Backup completed")
	}

	// Return acknowledgement
	return &proto.Ack{}, nil
}

func (this *auctionServer) Query(ctx context.Context, msg *proto.Empty) (*proto.Result, error) {
	// Lock and unlock
	this.mu.Lock()
	defer this.mu.Unlock()

	log.Printf("Received new query at time: %v", this.clock)

	// Increment clock
	this.clock++

	log.Printf("Sending result to client at time: %v", this.clock)

	// Return result
	return &proto.Result{
		Status:    this.status,
		Bid:       this.highestBid,
		Timestamp: this.clock,
	}, nil
}

func (this *auctionServer) Backup() {
	conn, err := grpc.NewClient("localhost:6000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error creating new client: %s", err)
	}
	client := proto.NewAuctionClient(conn)

	_, err = client.CallBackup(context.Background(), &proto.Result{
		Status:    this.status,
		Bid:       this.highestBid,
		Timestamp: this.clock,
	})
	if err != nil {
		log.Printf("Error backing up: %s", err)
	}

}

func (this *auctionServer) CallBackup(ctx context.Context, msg *proto.Result) (*proto.Ack, error) {
	// Lock and unlock
	this.mu.Lock()
	defer this.mu.Unlock()

	// Send fields to backup server
	this.clock = msg.Timestamp
	this.highestBid = msg.Bid
	this.status = msg.Status

	// Return acknowledgement
	return &proto.Ack{}, nil
}

func crashServer(this *auctionServer) {
	log.Printf("Change da world, my final message. Goodbye.")
	panic(this)
}
