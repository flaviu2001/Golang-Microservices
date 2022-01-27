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
