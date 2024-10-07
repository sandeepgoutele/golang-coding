package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sandeepgoutele/golang-coding/grpc_events/events"
	"google.golang.org/grpc"
)

func main() {
	// Set up a new connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewEventServiceClient(conn)

	// Create the event to send
	event := &pb.EventRequest{
		EventId:      "event123",
		EventMessage: "This is a test event",
		Timestamp:    time.Now().UnixNano(),
	}

	// Call SendEvent RPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SendEvent(ctx, event)
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	log.Printf("Received response: %s", response.Status)
}
