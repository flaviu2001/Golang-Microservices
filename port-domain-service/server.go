package main

import (
	"Bleenco/port-domain-service/service"
	"Bleenco/port-domain-service/utils"
	pb "Bleenco/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
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

		s.service.Upsert(utils.RpcPortToJsonPort(rpcPort))
	}
}

func (s *server) Select(rpcPage *pb.RpcPage, stream pb.Communicator_SelectServer) error {
	ports := s.service.Select(int(rpcPage.Page))
	for _, port := range ports {
		if err := stream.Send(utils.JsonPortToRpcPort(port)); err != nil {
			return err
		}
	}
	return nil
}
