package main

import (
	"Bleenco/common"
	pb "Bleenco/rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var c pb.CommunicatorClient

func handleParser(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go func() {
		entriesChannel, errorChannel := GetPorts()
		entriesOpen := true
		errorOpen := true
		running := true
		var entry common.Entry
		stream, err := c.Upsert(context.Background())
		for running {
			select {
			case entry, entriesOpen = <-entriesChannel:
				if entriesOpen {
					if err := stream.Send(common.JsonPortToRpcPort(entry.Port)); err != nil {
						common.CheckError(err)
					}
				} else {
					entriesChannel = nil
				}
			case err, errorOpen = <-errorChannel:
				if errorOpen {
					_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
				} else {
					errorChannel = nil
				}
			default:
				if entriesChannel == nil && errorChannel == nil {
					running = false
				}
			}
		}
		_, err = stream.CloseAndRecv()
		if err != nil {
			common.CheckError(err)
		}
	}()
	var response = common.JsonStatusResponse{Status: "started"}
	err := json.NewEncoder(w).Encode(response)
	common.CheckError(err)
}

func handleSelect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := mux.Vars(r)["page"]
	intPage, err := strconv.Atoi(page)
	common.CheckError(err)
	stream, err := c.Select(context.Background(), &pb.RpcPage{Page: int32(intPage)})
	var ports = make([]common.Port, 0)
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			break
		}
		common.CheckError(err)
		ports = append(ports, common.RpcPortToJsonPort(port))
	}

	err = json.NewEncoder(w).Encode(common.JsonPortsResponse{Status: "success", Ports: ports})
	common.CheckError(err)
}

func main() {
	var addr string
	var port string
	if addr = os.Getenv(common.GrpcServerAddr); addr == "" {
		addr = common.DefaultAddress
	}
	if port = os.Getenv(common.GrpcServerPort); port == "" {
		port = common.DefaultPort
	}
	addr = fmt.Sprintf("%s:%s", addr, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		common.CheckError(err)
	}(conn)
	c = pb.NewCommunicatorClient(conn)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/parse", handleParser).Methods("GET")
	router.HandleFunc("/select/{page}", handleSelect).Methods("GET")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
