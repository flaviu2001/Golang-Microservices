package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) UpsertPort(unlocs string, name string, city string, country string, aliases []string, regions []string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string) {
	conn := p.initConnection()
	defer p.closeConnection(conn)

	connTx := p.BeginTransaction(conn)
	defer p.EndTransaction(connTx)

	p.removeAliases(connTx, unlocs)
	p.removeRegions(connTx, unlocs)

	_, err := connTx.Exec(upsertPortStatement, unlocs, name, city, country, coord1, coord2, province, timezone, code)
	common.CheckError(err)

	portId := p.findPortId(connTx, unlocs)

	for _, alias := range aliases {
		p.insertAlias(connTx, portId, unlocs, alias)
	}

	for _, region := range regions {
		p.insertRegion(connTx, portId, unlocs, region)
	}
}

func (p *RepositoryImpl) findPortId(connTx *sql.Tx, unlocs string) int64 {
	rows, err := connTx.Query(selectPortId, unlocs)
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

func (p *RepositoryImpl) SelectPorts(lowerBound int, upperBound int) []common.Port {
	conn := p.initConnection()
	defer p.closeConnection(conn)

	rows, err := conn.Query(paginatedSelectPort, lowerBound, upperBound)
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
			Alias:       p.getAliases(conn, unlocs),
			Regions:     p.getRegions(conn, unlocs),
			Coordinates: coordinates,
			Province:    province,
			Timezone:    timezone,
			Unlocs:      []string{unlocs},
			Code:        dereferencedCode,
		})
	}
	return ports
}
