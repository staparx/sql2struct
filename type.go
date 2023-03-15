package main

type GoType string

const (
	GoTypeInt64 GoType = "int64"
	GoTypeInt32 GoType = "int32"
	GoTypeInt   GoType = "int"

	GoTypeFloat32 GoType = "float32"
	GoTypeFloat64 GoType = "float64"

	GoTypeTime GoType = "time.Time"

	GoTypeString GoType = "string"
)

var GoTypeMap = map[string]GoType{
	"tinyint":    GoTypeInt64,
	"smallint":   GoTypeInt64,
	"mediumint":  GoTypeInt64,
	"int":        GoTypeInt64,
	"bigint":     GoTypeInt64,
	"float":      GoTypeFloat64,
	"double":     GoTypeFloat64,
	"decimal":    GoTypeFloat64,
	"date":       GoTypeTime,
	"time":       GoTypeTime,
	"year":       GoTypeTime,
	"datetime":   GoTypeTime,
	"timestamp":  GoTypeTime,
	"char":       GoTypeString,
	"varchar":    GoTypeString,
	"tinyblob":   GoTypeString,
	"tinytext":   GoTypeString,
	"blob":       GoTypeString,
	"text":       GoTypeString,
	"mediumblob": GoTypeString,
	"mediumtext": GoTypeString,
	"longblob":   GoTypeString,
	"longtext":   GoTypeString,
}

type TagType struct {
	Prefix string
	Suffix string
}

var TagTypeMap = map[string]TagType{
	"json": {"", ""},
	"gorm": {"column:", ""},
}
