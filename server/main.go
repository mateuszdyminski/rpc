package main

import (
	"log"
	"net"
	"sync"

	streams "github.com/mateuszdyminski/rpc/streams"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// streamService is used to implement streams.Pipeline.
type streamService struct {
	clients  map[string]chan *streams.Message
	messages []*streams.Message
	m        sync.Mutex
}

// Send sends message which will be broadcasting
func (s *streamService) Send(ctx context.Context, in *streams.Message) (*streams.Response, error) {
	log.Printf("Received msg: %+v \n", in)
	s.broadcast(in)
	return new(streams.Response), nil
}

func (s *streamService) broadcast(in *streams.Message) {
	s.m.Lock()
	s.messages = append(s.messages, in)
	for k := range s.clients {
		s.clients[k] <- in
	}
	s.m.Unlock()
}

// List opens stream and sends all received messages to the clients
func (s *streamService) List(in *streams.Request, stream streams.Pipeline_ListServer) error {
	toSend := make(chan *streams.Message)
	s.m.Lock()
	log.Printf("Registering client %+v", in)
	s.clients[in.Origin] = toSend
	s.m.Unlock()

	for {
		msg := <-toSend
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	streams.RegisterPipelineServer(s, &streamService{clients: make(map[string]chan *streams.Message)})
	log.Printf("starting... \n")
	s.Serve(lis)
}
