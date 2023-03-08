package database

type TableDescribe struct {
	Field      string `json:"field"`
	FieldType  string `json:"type"`
	Collation  string `json:"collation"`
	Null       string `json:"null"`
	Key        string `json:"key"`
	Default    string `json:"default"`
	Extra      string `json:"extra"`
	Privileges string `json:"privileges"`
	Comment    string `json:"comment"`
}

type TableCreate struct {
	Table       string
	CreateTable string
}

type DBTableDescribeInterface interface {
	// GetTableDescribeList 获取表字段结构
	GetTableDescribeList() (map[string][]*TableDescribe, error)
	// ShowCreateTables 获取建表语句
	ShowCreateTables() ([]*TableCreate, error)
}
