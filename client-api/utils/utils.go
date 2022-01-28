package utils

import (
	"Bleenco/rpc"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func FromEnvVar(envName string, defaultValue string) string {
	var value string
	if value = os.Getenv(envName); value == "" {
		value = defaultValue
	}
	return value
}

func RpcPortToJsonPort(rpcPort *rpc.RpcPort) Port {
	return Port{
		Name:        rpcPort.Name,
		City:        rpcPort.City,
		Country:     rpcPort.Country,
		Alias:       rpcPort.Alias,
		Regions:     rpcPort.Regions,
		Coordinates: rpcPort.Coordinates,
		Province:    rpcPort.Province,
		Timezone:    rpcPort.Timezone,
		Unlocs:      rpcPort.Unlocs,
		Code:        rpcPort.Code,
	}
}

func JsonPortToRpcPort(jsonPort Port) *rpc.RpcPort {
	return &rpc.RpcPort{
		Name:        jsonPort.Name,
		City:        jsonPort.City,
		Country:     jsonPort.Country,
		Alias:       jsonPort.Alias,
		Regions:     jsonPort.Regions,
		Coordinates: jsonPort.Coordinates,
		Province:    jsonPort.Province,
		Timezone:    jsonPort.Timezone,
		Unlocs:      jsonPort.Unlocs,
		Code:        jsonPort.Code,
	}
}
