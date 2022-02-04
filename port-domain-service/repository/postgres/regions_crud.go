package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) GetRegions(unlocs string) []string {
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

func (p *RepositoryImpl) InsertRegion(portId int64, unlocs string, region string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(insertRegion, portId, unlocs, region)
	common.CheckError(err)
}

func (p *RepositoryImpl) RemoveRegions(unlocs string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(removeRegions, unlocs)
	common.CheckError(err)
}
