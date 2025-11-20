package server

import (
	proto "Mandatory_5_-_Auction_System/grpc"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type auctionServer struct {
	proto.UnimplementedAuctionServer
	mu    sync.Mutex
	clock int32
}

func newServer() *auctionServer {
	return &auctionServer{
		clock: 0,
	}
}

func StartServer() error {
	port := ":5050"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", port, err)
	}

	server := newServer()
	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Server listening on %s", port)
	return nil
}
