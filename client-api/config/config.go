package config

import (
	"Bleenco/client-api/constants"
	"Bleenco/common"
)

type Config struct {
	GrpcServerAddr string
	GrpcServerPort string
}

func NewConfig() Config {
	return Config{
		GrpcServerAddr: common.FromEnvVar(constants.GrpcServerAddr, constants.DefaultAddress),
		GrpcServerPort: common.FromEnvVar(common.GrpcServerPort, common.DefaultPort),
	}
}
