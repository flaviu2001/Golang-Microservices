package config

import (
	"Bleenco/common"
	"Bleenco/port-domain-service/constants"
)

type Config struct {
	GrpcServerPort string
	DbHost         string
	DbPort         string
	DbUser         string
	DbPass         string
	DbName         string
}

func NewConfig() Config {
	return Config{
		GrpcServerPort: common.FromEnvVar(common.GrpcServerPort, common.DefaultPort),
		DbHost:         common.FromEnvVar(constants.EnvDbHost, constants.DbHost),
		DbPort:         common.FromEnvVar(constants.EnvDbPort, constants.DbPort),
		DbUser:         common.FromEnvVar(constants.EnvDbUser, constants.DbUser),
		DbPass:         common.FromEnvVar(constants.EnvDbPass, constants.DbPassword),
		DbName:         common.FromEnvVar(constants.EnvDbName, constants.DbName),
	}
}
