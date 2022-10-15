package main

import (
	"log"
	"net"
	"simple_bank/gapi"
	"simple_bank/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("gRPC server at", "0.0.0.0:9090")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}