package file

import (
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	templateFile := "../"
	filePath := "./demo/test"
	var table = &Table{
		TableName:   "user",
		GoTableName: "User",
		PackageName: "file",
		Fields: []*Column{
			{
				GoColumnName:  "Name",
				GoColumnType:  "string",
				ColumnName:    "name",
				ColumnComment: "姓名",
				ColumnType:    "varchar(255)",
			},
			{
				GoColumnName:  "Age",
				GoColumnType:  "int64",
				ColumnName:    "age",
				ColumnComment: "年龄",
				ColumnType:    "int",
			},
		},
	}
	err := WriteToFileAndCreate(templateFile, filePath, table)
	if err != nil {
		t.Errorf("The file was not created successfully, err:%s", err.Error())
		return
	}
	return
}

func TestLookPath(t *testing.T) {
	path, err := filepath.Abs("./")
	if err != nil {
		t.Errorf("abs error:%s", err)
		return
	}

	t.Logf("success: %s", path)
}
