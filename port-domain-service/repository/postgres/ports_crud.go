package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) GetNewPortId() int64 {
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

func (p *RepositoryImpl) UpsertPort(portId int64, unlocs string, name string, city string, country string, coord1 interface{}, coord2 interface{}, province string, timezone string, code string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(upsertPortStatement, portId, unlocs, name, city, country, coord1, coord2, province, timezone, code)
	common.CheckError(err)
}

func (p *RepositoryImpl) FindPortId(unlocs string) int64 {
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

func (p *RepositoryImpl) SelectPorts(lowerBound int, upperBound int) []common.Port {
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
