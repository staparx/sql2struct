package mysql

import "github.com/pingcap/tidb/parser/mysql"

type GoType string

const (
	GoTypeTime    GoType = "time.Time"
	GoTypeString  GoType = "string"
	GoTypeInt64   GoType = "int64"
	GoTypeFloat32 GoType = "float32"
	GoTypeFloat64 GoType = "float64"
)

var TypeMap = map[byte]GoType{
	mysql.TypeUnspecified: GoTypeString,
	mysql.TypeTiny:        GoTypeInt64,
	mysql.TypeShort:       GoTypeInt64,
	mysql.TypeLong:        GoTypeInt64,
	mysql.TypeFloat:       GoTypeFloat64,
	mysql.TypeDouble:      GoTypeFloat32,
	mysql.TypeNull:        GoTypeString,
	mysql.TypeTimestamp:   GoTypeTime,
	mysql.TypeLonglong:    GoTypeInt64,
	mysql.TypeInt24:       GoTypeInt64,
	mysql.TypeDate:        GoTypeTime,
	mysql.TypeDuration:    GoTypeTime,
	mysql.TypeDatetime:    GoTypeTime,
	mysql.TypeYear:        GoTypeTime,
	mysql.TypeNewDate:     GoTypeTime,
	mysql.TypeVarchar:     GoTypeString,
	mysql.TypeBit:         GoTypeString,

	mysql.TypeJSON:       GoTypeString,
	mysql.TypeNewDecimal: GoTypeFloat64,
	mysql.TypeEnum:       GoTypeString,
	mysql.TypeSet:        GoTypeString,
	mysql.TypeTinyBlob:   GoTypeString,
	mysql.TypeMediumBlob: GoTypeString,
	mysql.TypeLongBlob:   GoTypeString,
	mysql.TypeBlob:       GoTypeString,
	mysql.TypeVarString:  GoTypeString,
	mysql.TypeString:     GoTypeString,
	mysql.TypeGeometry:   GoTypeString,
}
