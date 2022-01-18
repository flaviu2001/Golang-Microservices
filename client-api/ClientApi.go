package main

import (
	"Bleenco/common"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func GetEntries(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	entriesChannel, errorChannel := GetPorts()
	entriesOpen := true
	errorOpen := true
	running := true
	var entry common.Entry
	var err error
	for running {
		select {
		case entry, entriesOpen = <-entriesChannel:
			if entriesOpen {
				post, posterr := json.Marshal(entry.Port)
				common.CheckError(posterr)

				resp, httperr := http.Post("http://localhost:8080/upsert", "application/json", bytes.NewBuffer(post))
				common.CheckError(httperr)
				err := resp.Body.Close()
				common.CheckError(err)
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

	var response = common.JsonStatusResponse{Status: "success"}

	err = json.NewEncoder(w).Encode(response)
	common.CheckError(err)
}

func handleSelect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := mux.Vars(r)["page"]
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/select/%s", page))
	common.CheckError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		common.CheckError(err)
	}(resp.Body)
	var ports interface{}
	err = json.NewDecoder(resp.Body).Decode(&ports)
	common.CheckError(err)
	err = json.NewEncoder(w).Encode(common.JsonPortsResponseNoTypeCast{Status: "success", Ports: ports})
	common.CheckError(err)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/parse", GetEntries).Methods("GET")
	router.HandleFunc("/select/{page}", handleSelect).Methods("GET")

	fmt.Println("Server at 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
