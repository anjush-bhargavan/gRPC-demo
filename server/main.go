package main

import (
	"log"
	"net"
	pb "github.com/anjush-bhargavan/gRPC-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080" //defining port
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port) //creates a listener
	if err != nil {
		log.Fatalf("Failed to start the server %v",err)
	}

	grpcServer := grpc.NewServer()  //creates grpc server
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) //register helloSErver structy with gRPC server

	log.Printf("server started at %v",lis.Addr())
	
	if err := grpcServer.Serve(lis); err != nil { //start serving incoming gRPC requestes on the listener
		log.Fatalf("Failed to start: %v",err)
	}
}
