package postgres

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
		"id SERIAL PRIMARY KEY," +
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

	selectPortId = "SELECT id FROM ports WHERE unlocs = $1"

	upsertPortStatement = "INSERT INTO ports (unlocs, name, city, country, coord1, coord2, province, timezone, code) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) " +
		"ON CONFLICT (unlocs) DO UPDATE " +
		"SET name = $2, city = $3, country = $4, coord1 = $5, coord2 = $6, province = $7, timezone = $8, code = $9"
	insertAlias         = "INSERT INTO aliases(port_id, unlocs, alias) values ($1, $2, $3)"
	insertRegion        = "INSERT INTO regions(port_id, unlocs, region) values ($1, $2, $3)"
	paginatedSelectPort = "SELECT * FROM ports WHERE id BETWEEN $1 AND $2"
	selectAliases       = "SELECT alias FROM aliases WHERE unlocs = $1"
	selectRegions       = "SELECT region FROM regions WHERE unlocs = $1"
)
