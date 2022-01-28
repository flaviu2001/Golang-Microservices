package config

import (
	"Bleenco/client-api/constants"
	"Bleenco/client-api/utils"
)

type Config struct {
	GrpcServerAddr string
	GrpcServerPort string
}

func NewConfig() Config {
	return Config{
		GrpcServerAddr: utils.FromEnvVar(constants.GrpcServerAddr, constants.DefaultAddress),
		GrpcServerPort: utils.FromEnvVar(constants.GrpcServerPort, constants.DefaultPort),
	}
}
