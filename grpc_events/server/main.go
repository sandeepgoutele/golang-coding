package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "grpc_events/events"

	"github.com/bxcodec/faker/v3"
	"google.golang.org/grpc"
)

// server is used to implement EventServiceServer.
type server struct {
	pb.UnimplementedEventServiceServer
}

// SendEvent handles the unary RPC call
func (s *server) SendEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	log.Printf("Received event: ID=%s, Status=%s", req.EventId, req.Status)

	// Process the event and return a response
	return &pb.EventResponse{
		EventId:      "event123",
		EventMessage: "This is a test event",
		Timestamp:    time.Now().Unix(),
	}, nil
}

// StreamEvents implements the server-side streaming method
func (s *server) StreamEvents(req *pb.EventStreamRequest, stream pb.EventService_StreamEventsServer) error {
	log.Printf("Client connected: %s", req.ClientId)

	// Continuously generate and send random events
	for {
		// Simulate a random delay between events
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		// Generate a random event using faker
		event := &pb.EventResponse{
			EventId:      faker.UUIDDigit(),
			EventMessage: faker.Sentence(),
			Timestamp:    time.Now().Unix(),
		}

		// Send the event to the client
		if err := stream.Send(event); err != nil {
			return err
		}

		log.Printf("Sent event: ID=%s, Message=%s", event.EventId, event.EventMessage)
	}
}

func main() {
	// Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &server{})

	log.Printf("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
