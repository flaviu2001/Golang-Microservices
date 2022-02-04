package repository

import (
	"Bleenco/common"
	"database/sql"
)

// Repository This repository contains the simple instructions that are put together in a service to form complex operations
type Repository interface {
	// UpsertPort This method inserts the port, or if there is an unlocs conflict there will be an update
	UpsertPort(unlocs string, name string, city string, country string, aliases []string, regions []string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string)

	// SelectPorts This method returns all the ports with the id between a lower and an upper bound. This is
	// to be used in a pagination mechanism
	SelectPorts(lowerBound int, upperBound int) []common.Port

	// BeginTransaction This method starts a transaction marking that from the point of view of the database
	// everything inside the transaction is considered atomic and won't overlap with common objects
	BeginTransaction(conn *sql.DB) *sql.Tx

	// EndTransaction This method ends the transaction started by BeginTransaction
	EndTransaction(connTx *sql.Tx)
}
