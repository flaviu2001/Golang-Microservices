package main

import (
	"Bleenco/common"
	"Bleenco/port-domain-service/repository"
	"Bleenco/port-domain-service/service"
	pb "Bleenco/rpc"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCommunicatorServer
	service service.Service
}

func (s *server) Upsert(stream pb.Communicator_UpsertServer) error {
	for {
		rpcPort, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}
		if err != nil {
			return err
		}
		s.service.Upsert(common.RpcPortToJsonPort(rpcPort))
	}
}

func (s *server) Select(rpcPage *pb.RpcPage, stream pb.Communicator_SelectServer) error {
	ports := s.service.Select(int(rpcPage.Page))
	for _, port := range ports {
		if err := stream.Send(common.JsonPortToRpcPort(port)); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Here the grpc server is initialised according to the port specified in an environment variable (or a default value
	// in the case of its omission)
	var port = common.FromEnvVar(common.GrpcServerPort, common.DefaultPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunicatorServer(s, &server{service: &service.Impl{Repository: &repository.PostgresRepository{}}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
