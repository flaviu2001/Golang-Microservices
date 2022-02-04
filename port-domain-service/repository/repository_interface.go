package repository

import (
	"Bleenco/common"
)

// Repository This repository contains the simple instructions that are put together in a service to form complex operations
type Repository interface {
	// RemoveAliases This method removes all aliases with the mentioned unlocs
	RemoveAliases(unlocs string)

	// RemoveRegions This method removes all regions with the mentioned unlocs
	RemoveRegions(unlocs string)

	// GetNewPortId This method returns the maximum id of all ports plus one
	GetNewPortId() int64

	// UpsertPort This method inserts the port, or if there is an unlocs conflict there will be an update
	UpsertPort(portId int64, unlocs string, name string, city string, country string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string)

	// FindPortId This method returns the id of the port with the mentioned unlocs
	FindPortId(unlocs string) int64

	// InsertAlias This method inserts an alias into the repository
	InsertAlias(portId int64, unlocs string, alias string)

	// InsertRegion This method inserts a region into the repository
	InsertRegion(portId int64, unlocs string, region string)

	// GetAliases This method returns all the aliases of a given port
	GetAliases(unlocs string) []string

	// GetRegions This method returns all the regions of a given port
	GetRegions(unlocs string) []string

	// SelectPorts This method returns all the ports with the id between a lower and an upper bound. This is
	// to be used in a pagination mechanism
	SelectPorts(lowerBound int, upperBound int) []common.Port
}
