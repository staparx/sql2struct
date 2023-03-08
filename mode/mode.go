package mode

import (
	"github.com/staparx/sql2struct/mode/db_mode"
	"github.com/staparx/sql2struct/mode/str_mode"
)

const (
	StrMode = "sql"
	DbMode  = "db"
)

type TableModeInterface interface {
	DbTable(param string) []string
}

func NewModeService(mode string) (resp TableModeInterface) {
	switch mode {
	case StrMode:
		resp = str_mode.NewStrMode()
	case DbMode:
		resp = db_mode.NewDbMode()
	default:
		return nil
	}
	return resp
}
