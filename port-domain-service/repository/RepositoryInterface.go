package repository

import "Bleenco/common"

type Repository interface {
	RemoveAliases(unlocs string)
	RemoveRegions(unlocs string)
	GetNewPortId() int64
	UpsertPort(portId int64, unlocs string, name string, city string, country string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string)
	FindPortId(unlocs string) int64
	InsertAlias(portId int64, unlocs string, alias string)
	InsertRegion(portId int64, unlocs string, region string)
	GetAliases(unlocs string) []string
	GetRegions(unlocs string) []string
	SelectPorts(lowerBound int, upperBound int) []common.Port
}
