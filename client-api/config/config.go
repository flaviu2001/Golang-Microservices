package config

import "Bleenco/common"

type Config struct {
	GrpcServerAddr string
	GrpcServerPort string
}

func NewConfig() Config {
	return Config{
		GrpcServerAddr: common.FromEnvVar(common.GrpcServerAddr, common.DefaultAddress),
		GrpcServerPort: common.FromEnvVar(common.GrpcServerPort, common.DefaultPort),
	}
}
