package main

import (
	"Bleenco/port-domain-service/config"
	"Bleenco/port-domain-service/repository/postgres"
	"Bleenco/port-domain-service/service"
	pb "Bleenco/rpc"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Here the grpc server is initialised according to the port specified in an environment variable (or a default value
	// in the case of its omission)
	cfg := config.NewConfig()
	var port = cfg.GrpcServerPort
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCommunicatorServer(s, &server{service: &service.Impl{Repository: &postgres.RepositoryImpl{}}})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
