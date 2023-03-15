package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Table struct {
	TableName      string    // table name
	GoTableName    string    // go struct name
	PackageName    string    // package name
	Fields         []*Column // columns
	ImportPackages []string  // imports
}
type Column struct {
	GoColumnName  string
	GoColumnType  GoType
	ColumnName    string // column_name
	ColumnType    string // column_type
	ColumnComment string // column_comment
	Tag           []*Tag // gorm_tag
}

type Tag struct {
	TagKey   string
	TagValue string
}

func (t *Table) WriteToFileAndCreate(tmplPath, filePath string) error {
	pattern := filepath.Join(tmplPath, "model.tmpl")
	templates := template.Must(template.ParseGlob(pattern))

	if !strings.HasSuffix(filePath, "/") {
		filePath += "/"
	}
	filePath += fmt.Sprintf("%s.go", t.TableName)

	file, err := os.Create(filePath)
	if err != nil {

		return err
	}
	defer file.Close()

	err = templates.Execute(file, t)
	if err != nil {
		return err
	}
	return nil
}

func PathVerify(filePath string) (string, string, string, error) {
	var err error
	tmplPath := cfg.Path.Template
	if tmplPath == "" {
		//查询模版文件是否存在
		gopath := os.Getenv("GOPATH")
		tmplPath = gopath + "/bin"
		_, err = os.Stat(tmplPath)
		if err != nil {
			return "", "", "", err
		}
	}
	log.Info("模版文件查询成功", zap.String("tmpl_path", tmplPath))

	//查询最终文件路径是否存在，不存在则创建
	//获取最终文件的绝对路径
	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return "", "", "", err
	}

	//获取最终文件的包名
	list := strings.Split(filePath, "/")
	packageName := list[len(list)-1]

	//查询最终文件路径是否存在
	_, err = os.Stat(filePath)

	if err == nil {
		return tmplPath, filePath, packageName, err
	}
	//不存在则进行创建
	if os.IsNotExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return "", "", "", err
		}
	}
	return tmplPath, filePath, packageName, nil
}
