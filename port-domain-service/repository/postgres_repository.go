package repository

import (
	"Bleenco/common"
	"database/sql"
	"fmt"
)

const (
	createAliasTable = "CREATE TABLE IF NOT EXISTS aliases (" +
		"id SERIAL PRIMARY KEY," +
		"port_id BIGINT REFERENCES ports(id)," +
		"unlocs VARCHAR," +
		"alias VARCHAR)"
	createRegionTable = "CREATE TABLE IF NOT EXISTS regions (" +
		"id SERIAL PRIMARY KEY," +
		"port_id BIGINT REFERENCES ports(id)," +
		"unlocs VARCHAR," +
		"region VARCHAR)"
	createPortTable = "CREATE TABLE IF NOT EXISTS ports (" +
		"id BIGINT PRIMARY KEY," +
		"unlocs VARCHAR UNIQUE," +
		"name VARCHAR," +
		"city VARCHAR," +
		"country VARCHAR," +
		"coord1 REAL," +
		"coord2 REAL," +
		"province VARCHAR," +
		"timezone VARCHAR," +
		"code VARCHAR)"

	removeAliases = "DELETE FROM aliases WHERE unlocs = $1"
	removeRegions = "DELETE FROM regions WHERE unlocs = $1"

	selectHighestId = "SELECT GREATEST(0, max(id)) from (select id from ports order by id desc limit 1) t"
	selectPortId    = "SELECT id FROM ports WHERE unlocs = $1"

	upsertPortStatement = "INSERT INTO ports " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) " +
		"ON CONFLICT (unlocs) DO UPDATE " +
		"SET name = $3, city = $4, country = $5, coord1 = $6, coord2 = $7, province = $8, timezone = $9, code = $10"
	insertAlias         = "INSERT INTO aliases(port_id, unlocs, alias) values ($1, $2, $3)"
	insertRegion        = "INSERT INTO regions(port_id, unlocs, region) values ($1, $2, $3)"
	paginatedSelectPort = "SELECT * FROM ports WHERE id BETWEEN $1 AND $2"
	selectAliases       = "SELECT alias FROM aliases WHERE unlocs = $1"
	selectRegions       = "SELECT region FROM regions WHERE unlocs = $1"
)

type PostgresRepository struct {
	// This flag specifies whether the tables are ensured to be created and the operations are safe to run
	databaseInitialized bool
	// This is the variable that holds a connection to the database
	conn *sql.DB
}

// initDatabase This method attempts to create the necessary tables and marks a flag upon success.
func (p *PostgresRepository) initDatabase() {
	if p.databaseInitialized {
		return
	}

	_, err := p.conn.Exec(createPortTable)
	common.CheckError(err)

	_, err = p.conn.Exec(createAliasTable)
	common.CheckError(err)

	_, err = p.conn.Exec(createRegionTable)
	common.CheckError(err)

	p.databaseInitialized = true
}

// initConnection This method initialises the connection to a postgres database
func (p *PostgresRepository) initConnection() {
	host := common.FromEnvVar(common.EnvDbHost, common.DbHost)
	port := common.FromEnvVar(common.EnvDbPort, common.DbPort)
	user := common.FromEnvVar(common.EnvDbUser, common.DbUser)
	password := common.FromEnvVar(common.EnvDbPass, common.DbPassword)
	dbname := common.FromEnvVar(common.EnvDbName, common.DbName)
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		common.CheckError(err)
	}

	p.conn = db
	p.initDatabase()
}

// closeConnection This method will be called when there is no more use for the connection.
func (p *PostgresRepository) closeConnection() {
	err := p.conn.Close()
	common.CheckError(err)
}

// The following methods have their documentation provided in the interface

func (p *PostgresRepository) RemoveAliases(unlocs string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(removeAliases, unlocs)
	common.CheckError(err)
}

func (p *PostgresRepository) RemoveRegions(unlocs string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(removeRegions, unlocs)
	common.CheckError(err)
}

func (p *PostgresRepository) GetNewPortId() int64 {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(selectHighestId)
	common.CheckError(err)
	var portId int64

	for rows.Next() {
		err = rows.Scan(&portId)
		common.CheckError(err)
	}

	err = rows.Close()
	common.CheckError(err)

	return portId + 1
}

func (p *PostgresRepository) UpsertPort(portId int64, unlocs string, name string, city string, country string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(upsertPortStatement, portId, unlocs, name, city, country, coord1, coord2, province, timezone, code)
	common.CheckError(err)
}

func (p *PostgresRepository) FindPortId(unlocs string) int64 {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(selectPortId, unlocs)
	common.CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		common.CheckError(err)
	}(rows)

	var portId int64

	for rows.Next() {
		err = rows.Scan(&portId)
		common.CheckError(err)
	}

	return portId
}

func (p *PostgresRepository) InsertAlias(portId int64, unlocs string, alias string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(insertAlias, portId, unlocs, alias)
	common.CheckError(err)
}

func (p *PostgresRepository) InsertRegion(portId int64, unlocs string, region string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(insertRegion, portId, unlocs, region)
	common.CheckError(err)
}

func (p *PostgresRepository) GetAliases(unlocs string) []string {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(selectAliases, unlocs)
	common.CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		common.CheckError(err)
	}(rows)

	aliases := make([]string, 0)

	for rows.Next() {
		var alias string
		err := rows.Scan(&alias)
		common.CheckError(err)
		aliases = append(aliases, alias)
	}

	return aliases
}

func (p *PostgresRepository) GetRegions(unlocs string) []string {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(selectRegions, unlocs)
	common.CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		common.CheckError(err)
	}(rows)

	regions := make([]string, 0)

	for rows.Next() {
		var region string
		err := rows.Scan(&region)
		common.CheckError(err)
		regions = append(regions, region)
	}

	return regions
}

func (p *PostgresRepository) SelectPorts(lowerBound int, upperBound int) []common.Port {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(paginatedSelectPort, lowerBound, upperBound)
	common.CheckError(err)

	ports := make([]common.Port, 0)

	for rows.Next() {
		var id int64
		var unlocs string
		var name string
		var city string
		var country string
		var coord1 *float32
		var coord2 *float32
		var province string
		var timezone string
		var code *string

		err := rows.Scan(&id, &unlocs, &name, &city, &country, &coord1, &coord2, &province, &timezone, &code)
		common.CheckError(err)

		var coordinates []float32
		if coord1 == nil {
			coordinates = []float32{}
		} else {
			coordinates = []float32{*coord1, *coord2}
		}

		var dereferencedCode = ""
		if code != nil {
			dereferencedCode = *code
		}

		ports = append(ports, common.Port{
			Name:        name,
			City:        city,
			Country:     country,
			Alias:       p.GetAliases(unlocs),
			Regions:     p.GetRegions(unlocs),
			Coordinates: coordinates,
			Province:    province,
			Timezone:    timezone,
			Unlocs:      []string{unlocs},
			Code:        dereferencedCode,
		})
	}
	return ports
}
