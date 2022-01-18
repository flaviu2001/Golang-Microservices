package main

import (
	"Bleenco/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func GetEntries(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	entriesChannel, errorChannel := utils.GetPorts()
	entriesOpen := true
	errorOpen := true
	running := true
	var entry utils.Entry
	var err error
	for running {
		select {
		case entry, entriesOpen = <-entriesChannel:
			if entriesOpen {
				post, posterr := json.Marshal(entry.Port)
				utils.CheckError(posterr)

				resp, httperr := http.Post("http://localhost:8080/upsert", "application/json", bytes.NewBuffer(post))
				utils.CheckError(httperr)
				err := resp.Body.Close()
				utils.CheckError(err)
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

	var response = utils.JsonStatusResponse{Status: "success"}

	err = json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/entries", GetEntries).Methods("GET")

	fmt.Println("Server at 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
