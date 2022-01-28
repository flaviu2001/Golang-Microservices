package postgres

import (
	"Bleenco/port-domain-service/config"
	"Bleenco/port-domain-service/utils"
	"database/sql"
	"fmt"
)

type RepositoryImpl struct {
	// This flag specifies whether the tables are ensured to be created and the operations are safe to run
	databaseInitialized bool
	// This is the variable that holds a connection to the database
	conn *sql.DB
}

// initDatabase This method attempts to create the necessary tables and marks a flag upon success.
func (p *RepositoryImpl) initDatabase() {
	if p.databaseInitialized {
		return
	}

	_, err := p.conn.Exec(createPortTable)
	utils.CheckError(err)

	_, err = p.conn.Exec(createAliasTable)
	utils.CheckError(err)

	_, err = p.conn.Exec(createRegionTable)
	utils.CheckError(err)

	p.databaseInitialized = true
}

// initConnection This method initialises the connection to a postgres database
func (p *RepositoryImpl) initConnection() {
	cfg := config.NewConfig()
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		utils.CheckError(err)
	}

	p.conn = db
	p.initDatabase()
}

// closeConnection This method will be called when there is no more use for the connection.
func (p *RepositoryImpl) closeConnection() {
	err := p.conn.Close()
	utils.CheckError(err)
}
