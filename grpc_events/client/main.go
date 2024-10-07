package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc_events/events"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server (CVM)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewEventServiceClient(conn)

	// Send a unary event
	sendEvent(client)

	// Start listening for streaming events
	streamEvents(client)
}

func sendEvent(client pb.EventServiceClient) {
	// Create an event
	event := &pb.EventRequest{
		EventId: "event123",
		Status:  "Sample event",
	}

	// Call the SendEvent RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SendEvent(ctx, event)
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	log.Printf("SendEvent response: ID=%s, Message=%s, Timestamp=%d", response.EventId,
		response.EventMessage, response.Timestamp)
}

func streamEvents(client pb.EventServiceClient) {
	// Prepare the request to start the stream
	req := &pb.EventStreamRequest{
		ClientId: "hypervisor-123",
	}

	// Start streaming events
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10) // Set a long context timeout
	defer cancel()

	stream, err := client.StreamEvents(ctx, req)
	if err != nil {
		log.Fatalf("Error starting event stream: %v", err)
	}

	// Listen for incoming events in a loop
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			// Server has finished sending events
			log.Println("Stream closed by server")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving event: %v", err)
		}

		// Process the received event
		log.Printf("Streamed event: ID=%s, Message=%s, Timestamp=%d", event.EventId, event.EventMessage, event.Timestamp)
	}
}
