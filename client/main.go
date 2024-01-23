package main

import (
	"log"
	pb "github.com/anjush-bhargavan/gRPC-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080" 
)

func main() {
	conn,err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials())) //creating a conncection to gRPC server running on 8080
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn) // create new instance of gRPC client, for the greetservice using the connection

	names := &pb.NamesList{
		Names: []string{"John","Bob","Jesse"},
	}

	callSayHello(client)
	callSayHelloServerStream(client,names)
	callSayHelloClientStream(client,names)
	CallHelloBidirectionalStream(client,names)
	
}