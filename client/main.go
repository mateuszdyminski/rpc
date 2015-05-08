package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	streams "github.com/mateuszdyminski/rpc/streams"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address = ""

func send(msg string) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := streams.NewPipelineClient(conn)

	message := &streams.Message{
		Msg: msg,
	}
	_, err = client.Send(context.Background(), message)
	return err
}

func list() error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := streams.NewPipelineClient(conn)

	stream, err := client.List(context.Background(), &streams.Request{fmt.Sprintf("clientID:%s", time.Now())})
	if err != nil {
		return err
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("received message: %+v \n", msg)
	}

	return nil
}

func main() {
	// Contact the server and print out its response.
	if len(os.Args) != 3 {
		log.Fatalf("wrong number of inputs: %d.", len(os.Args))
	}

	address = os.Args[1]

	send(os.Args[2])

	log.Printf("Message send, waiting for msgs")
	list()
}
