package utils

import (
	"Bleenco/common"
	"sync"
)

type JsonStatusResponse struct {
	Status string `json:"status"`
}

type JsonPortsResponse struct {
	Status string        `json:"status"`
	Ports  []common.Port `json:"ports"`
}

type Entry struct {
	PortName string
	Port     common.Port
}

type ParserState struct {
	CurrentlyParsing bool
	Mutex            sync.Mutex
}
