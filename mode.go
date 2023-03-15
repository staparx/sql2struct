package main

import "errors"

type ModeType string

const (
	SqlMode ModeType = "sql"
	DbMode  ModeType = "db"
)

type TableModeInterface interface {
	DbTable() []string
}

func NewModeService(mode ModeType, param string) (resp TableModeInterface, err error) {
	switch mode {

	case SqlMode:
		resp = NewDBSqlStr(param)

	case DbMode:

		resp = NewDBMysql(param)

	default:
		return nil, errors.New("模式错误，请选择正确的模式")
	}
	return resp, nil
}
