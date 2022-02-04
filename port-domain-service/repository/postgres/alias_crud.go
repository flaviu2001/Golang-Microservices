package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) GetAliases(unlocs string) []string {
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

func (p *RepositoryImpl) InsertAlias(portId int64, unlocs string, alias string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(insertAlias, portId, unlocs, alias)
	common.CheckError(err)
}

func (p *RepositoryImpl) RemoveAliases(unlocs string) {
	p.initConnection()
	defer p.closeConnection()

	_, err := p.conn.Exec(removeAliases, unlocs)
	common.CheckError(err)
}
