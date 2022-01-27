package main

import (
	"Bleenco/common"
	pb "Bleenco/rpc"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var c pb.CommunicatorClient

func main() {
	// The port domain service address is received from environment variables, otherwise defaulted to what is specified
	// in the constants file
	var addr = common.FromEnvVar(common.GrpcServerAddr, common.DefaultAddress)
	var port = common.FromEnvVar(common.GrpcServerPort, common.DefaultPort)
	fullAddr := fmt.Sprintf("%s:%s", addr, port)
	conn, err := grpc.Dial(fullAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		common.CheckError(err)
	}(conn)

	c = pb.NewCommunicatorClient(conn)

	// The REST Api is created using mux and two endpoints are provided.
	router := mux.NewRouter().StrictSlash(true)

	// All the handlers are defined in this method.
	registerRoutes(router)

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
