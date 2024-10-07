package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sandeepgoutele/golang-coding/grpc_events/events"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEventServiceServer
}

func (svr *server) SendEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	log.Printf("Received event: ID=%s, Message=%s, Timestamp=%d", req.EventId, req.EventMessage, req.Timestamp)

	response := &pb.EventResponse{
		Status: "Event received successfully",
	}
	return response, nil
}

func main() {
	// Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	svr := grpc.NewServer()
	pb.RegisterEventServiceServer(svr, &server{})

	log.Print("Server is listening on port 50051...")
	if err := svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
