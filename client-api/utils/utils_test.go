package utils

import (
	"Bleenco/rpc"
	"testing"
)

func TestRpcPortToJsonPort(t *testing.T) {
	rpcPort := &rpc.RpcPort{
		Name:        "a",
		City:        "b",
		Country:     "c",
		Alias:       []string{"d"},
		Regions:     []string{"e"},
		Coordinates: []float32{1, 2},
		Province:    "f",
		Timezone:    "g",
		Unlocs:      []string{"h"},
		Code:        "i",
	}

	port := RpcPortToJsonPort(rpcPort)
	if !(port.Name == "a" && port.City == "b" && port.Country == "c" && len(port.Alias) == 1 &&
		port.Alias[0] == "d" && len(port.Regions) == 1 && port.Regions[0] == "e" && len(port.Coordinates) == 2 &&
		port.Coordinates[0] == 1 && port.Coordinates[1] == 2 && port.Province == "f" && port.Timezone == "g" &&
		len(port.Unlocs) == 1 && port.Unlocs[0] == "h" && port.Code == "i") {
		t.Fatalf("Wrong port")
	}
}

func TestJsonPortToRpcPort(t *testing.T) {
	jsonPort := Port{
		Name:        "a",
		City:        "b",
		Country:     "c",
		Alias:       []string{"d"},
		Regions:     []string{"e"},
		Coordinates: []float32{1, 2},
		Province:    "f",
		Timezone:    "g",
		Unlocs:      []string{"h"},
		Code:        "i",
	}

	port := JsonPortToRpcPort(jsonPort)
	if !(port.Name == "a" && port.City == "b" && port.Country == "c" && len(port.Alias) == 1 &&
		port.Alias[0] == "d" && len(port.Regions) == 1 && port.Regions[0] == "e" && len(port.Coordinates) == 2 &&
		port.Coordinates[0] == 1 && port.Coordinates[1] == 2 && port.Province == "f" && port.Timezone == "g" &&
		len(port.Unlocs) == 1 && port.Unlocs[0] == "h" && port.Code == "i") {
		t.Fatalf("Wrong port")
	}
}
