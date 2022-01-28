package postgres

import (
	"Bleenco/port-domain-service/utils"
	"database/sql"
)

func (p *RepositoryImpl) GetRegions(unlocs string) []string {
	p.initConnection()
	defer p.closeConnection()

	rows, err := p.conn.Query(selectRegions, unlocs)
	utils.CheckError(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		utils.CheckError(err)
	}(rows)

	regions := make([]string, 0)

	for rows.Next() {
		var region string
		err := rows.Scan(&region)
		utils.CheckError(err)
		regions = append(regions, region)
	}

	return regions
}

func (p *RepositoryImpl) InsertRegion(portId int64, unlocs string, region string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(insertRegion, portId, unlocs, region)
	utils.CheckError(err)
}

func (p *RepositoryImpl) RemoveRegions(unlocs string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(removeRegions, unlocs)
	utils.CheckError(err)
}
