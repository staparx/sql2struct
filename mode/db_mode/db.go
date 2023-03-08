package db_mode

import (
	"github.com/pingcap/log"
	"github.com/staparx/sql2struct/database/mysql"
)

type dbMode struct {
}

func NewDbMode() *dbMode {
	return &dbMode{}
}

func (m *dbMode) DbTable(param string) []string {
	db := mysql.NewDBMysql(param)
	tableList, err := db.ShowCreateTables()
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	resp := make([]string, len(tableList), len(tableList))

	for index, table := range tableList {
		resp[index] = table.CreateTable
	}

	return resp
}
