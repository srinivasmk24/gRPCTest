package main

import (
    "log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"placnetworks.com/chat"
)

func main() {
    var conn = *gRPC.ClientConn
    conn, err := grpc.Dail(":9000", gRPC.WithInsecure)
	if err != nil {
		log.Fatal("could not connect: %s", err)
	}

    defer conn.Close()
    c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)

}