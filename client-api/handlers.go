package main

import (
	"Bleenco/client-api/constants"
	"Bleenco/client-api/parser"
	"Bleenco/common"
	pb "Bleenco/rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
)

// handleParser this method starts parsing the json file and each Port, one by one, will be fed to the
// port domain service to be persisted. All this happens in a background thread so that the user will
// not have to wait for the call to finish while all this is happening.
func handleParser(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go func() {
		entriesChannel, errorChannel := parser.GetPorts(constants.PortsJsonFilename)

		// entriesOpen and errorOpen mark whether the channels are still open or closed.
		entriesOpen := true
		errorOpen := true
		running := true

		// Entry read from the channel
		var entry common.Entry

		// Call to the server to upsert the entries
		stream, err := c.Upsert(context.Background())
		for running {
			select {
			// There is an entry in the channel
			case entry, entriesOpen = <-entriesChannel:
				if entriesOpen {
					// Send the entry through the stream to the port domain service
					if err := stream.Send(common.JsonPortToRpcPort(entry.Port)); err != nil {
						common.CheckError(err)
					}
				} else {
					// Mark the entriesChannel as nil so that no further reads will succeed or even be attempted
					entriesChannel = nil
				}
			case err, errorOpen = <-errorChannel:
				if errorOpen {
					_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
				} else {
					// Similar to entriesChannel
					errorChannel = nil
				}
			default:
				// If both channels are found to be closed the loop will finish and the method will be allowed to exit.
				if entriesChannel == nil && errorChannel == nil {
					running = false
				}
			}
		}

		// Close the stream to let the server know no more upserts will happen
		_, err = stream.CloseAndRecv()

		if err != nil {
			common.CheckError(err)
		}
	}()
	// Return a simple response to the user
	var response = common.JsonStatusResponse{Status: "started"}
	err := json.NewEncoder(w).Encode(response)
	common.CheckError(err)
}

// handleSelect this method returns the persisted Ports from the port domain service using pagination. You specify the
// page and it will retrieve the respective 100 ports in a json.
func handleSelect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := mux.Vars(r)["page"]
	intPage, err := strconv.Atoi(page)
	common.CheckError(err)

	// call to the port domain server to receive the required ports
	stream, err := c.Select(context.Background(), &pb.RpcPage{Page: int32(intPage)})
	var ports = make([]common.Port, 0)

	for {
		port, err := stream.Recv()
		if err == io.EOF {
			break
		}

		common.CheckError(err)

		// Build a list from all the ports and return it
		ports = append(ports, common.RpcPortToJsonPort(port))
	}

	err = json.NewEncoder(w).Encode(common.JsonPortsResponse{Status: "success", Ports: ports})
	common.CheckError(err)
}
