package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	mysqlx "github.com/staparx/sql2struct/database/mysql"
)

type Table struct {
	TableName      string    // table name
	GoTableName    string    // go struct name
	PackageName    string    // package name
	Fields         []*Column // columns
	ImportPackages []string
}
type Column struct {
	GoColumnName  string
	GoColumnType  mysqlx.GoType
	ColumnName    string // column_name
	ColumnType    string // column_type
	ColumnComment string // column_comment
	GormTag       string
}

func WriteToFileAndCreate(tmplPath, filePath string, body *Table) error {
	pattern := filepath.Join(tmplPath, "model.tmpl")
	templates := template.Must(template.ParseGlob(pattern))

	if !strings.HasSuffix(filePath, "/") {
		filePath += "/"
	}
	filePath += fmt.Sprintf("%s.go", body.TableName)

	file, err := os.Create(filePath)
	if err != nil {

		return err
	}
	defer file.Close()

	err = templates.Execute(file, body)
	if err != nil {
		return err
	}
	return nil
}
