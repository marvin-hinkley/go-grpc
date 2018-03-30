package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"../person"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := person.Server{}
	grpcServer := grpc.NewServer()

	person.RegisterPersonServiceServer(grpcServer, &s)
	grpcServer.Serve(lis)
}
