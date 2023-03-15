package main

import (
	"errors"
	"fmt"
	"github.com/xwb1989/sqlparser"
)

func ParseSQLToFileStruct(sql string, packageName string) (*Table, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}
	switch ddl := stmt.(type) {
	default:
		return nil, errors.New("sql is not DDL")

	case *sqlparser.DDL:
		//获取表名
		tableName := ddl.NewName.Name.String()
		//驼峰转换成Go的类型名
		goTableName, err := ToCase(cfg.Case.Old, cfg.Case.New, tableName)
		if err != nil {
			return nil, err
		}

		//构建模版文件需要的结构体
		table := &Table{
			TableName:      tableName,
			GoTableName:    goTableName,
			PackageName:    packageName,
			Fields:         make([]*Column, 0, len(ddl.TableSpec.Columns)),
			ImportPackages: make([]string, 0, 0),
		}

		//字段获取
		for _, column := range ddl.TableSpec.Columns {
			columnName := column.Name.String()
			columnType := column.Type.Type

			goColumnName, err := ToCase(cfg.Case.Old, cfg.Case.New, columnName)
			if err != nil {
				return nil, err
			}

			//go对应的类型
			goColumnType, ok := cfg.TypeMap[columnType]
			if !ok {
				goColumnType = GoTypeString
			}

			if goColumnType == GoTypeTime {
				table.ImportPackages = append(table.ImportPackages, `"time"`)
			}

			fileColumn := &Column{
				GoColumnName: goColumnName,
				GoColumnType: goColumnType,
				ColumnName:   columnName,
				ColumnType:   columnType,
				Tag:          make([]*Tag, 0, len(cfg.Tag)),
			}

			//获取字段的备注
			if column.Type.Comment != nil {
				fileColumn.ColumnComment = BytesToStr(column.Type.Comment.Val)
			}

			//组装反射的字段
			for _, t := range cfg.Tag {
				tagType := TagTypeMap[t]
				tagValue := fmt.Sprintf("%s%s%s", tagType.Prefix, columnName, tagType.Suffix)

				tag := &Tag{
					TagKey:   t,
					TagValue: tagValue,
				}
				fileColumn.Tag = append(fileColumn.Tag, tag)
			}

			table.Fields = append(table.Fields, fileColumn)
		}

		return table, nil
	}
}
