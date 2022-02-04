package postgres

import (
	"Bleenco/common"
	"database/sql"
)

func (p *RepositoryImpl) getAliases(conn *sql.DB, unlocs string) []string {
	rows, err := conn.Query(selectAliases, unlocs)
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

func (p *RepositoryImpl) insertAlias(connTx *sql.Tx, portId int64, unlocs string, alias string) {
	_, err := connTx.Exec(insertAlias, portId, unlocs, alias)
	common.CheckError(err)
}

func (p *RepositoryImpl) removeAliases(connTx *sql.Tx, unlocs string) {
	_, err := connTx.Exec(removeAliases, unlocs)
	common.CheckError(err)
}
