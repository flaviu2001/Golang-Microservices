package config

import (
	"Bleenco/port-domain-service/constants"
	"Bleenco/port-domain-service/utils"
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
		GrpcServerPort: utils.FromEnvVar(constants.GrpcServerPort, constants.DefaultPort),
		DbHost:         utils.FromEnvVar(constants.EnvDbHost, constants.DbHost),
		DbPort:         utils.FromEnvVar(constants.EnvDbPort, constants.DbPort),
		DbUser:         utils.FromEnvVar(constants.EnvDbUser, constants.DbUser),
		DbPass:         utils.FromEnvVar(constants.EnvDbPass, constants.DbPassword),
		DbName:         utils.FromEnvVar(constants.EnvDbName, constants.DbName),
	}
}
