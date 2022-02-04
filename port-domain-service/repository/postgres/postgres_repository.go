package postgres

import (
	"Bleenco/common"
	"Bleenco/port-domain-service/config"
	"database/sql"
	"fmt"
)

type RepositoryImpl struct {
	// This flag specifies whether the tables are ensured to be created and the operations are safe to run
	databaseInitialized bool
}

// initDatabase This method attempts to create the necessary tables and marks a flag upon success.
func (p *RepositoryImpl) initDatabase(conn *sql.DB) {
	if p.databaseInitialized {
		return
	}

	_, err := conn.Exec(createPortTable)
	common.CheckError(err)

	_, err = conn.Exec(createAliasTable)
	common.CheckError(err)

	_, err = conn.Exec(createRegionTable)
	common.CheckError(err)

	p.databaseInitialized = true
}

// initConnection This method initialises the connection to a postgres database
func (p *RepositoryImpl) initConnection() *sql.DB {
	cfg := config.NewConfig()
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		common.CheckError(err)
	}

	p.initDatabase(db)

	return db
}

// closeConnection This method will be called when there is no more use for the connection.
func (p *RepositoryImpl) closeConnection(conn *sql.DB) {
	err := conn.Close()
	common.CheckError(err)
}

func (p *RepositoryImpl) BeginTransaction(conn *sql.DB) *sql.Tx {
	begin, err := conn.Begin()
	if err != nil {
		common.CheckError(err)
	}
	return begin
}

func (p *RepositoryImpl) EndTransaction(transaction *sql.Tx) {
	err := transaction.Commit()
	if err != nil {
		common.CheckError(err)
	}
}
