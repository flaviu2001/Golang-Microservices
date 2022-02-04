package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) getRegions(conn *sql.DB, unlocs string) []string {
	rows, err := conn.Query(selectRegions, unlocs)
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

func (p *RepositoryImpl) insertRegion(connTx *sql.Tx, portId int64, unlocs string, region string) {
	_, err := connTx.Exec(insertRegion, portId, unlocs, region)
	common.CheckError(err)
}

func (p *RepositoryImpl) removeRegions(connTx *sql.Tx, unlocs string) {
	_, err := connTx.Exec(removeRegions, unlocs)
	common.CheckError(err)
}
