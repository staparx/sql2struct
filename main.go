package main

import (
	"flag"

	"github.com/staparx/sql2struct/database/mysql"

	"go.uber.org/zap"

	"github.com/pingcap/log"

	"github.com/staparx/sql2struct/config"
	"github.com/staparx/sql2struct/file"
	x_mode "github.com/staparx/sql2struct/mode"
	"github.com/staparx/sql2struct/parser"
	"github.com/staparx/sql2struct/util"
)

var (
	mode, dbName, path, sql string
)

func main() {
	var err error

	flag.StringVar(&mode, "mode", "db", "执行模式 sql：建表语句解析 db:数据库解析")
	flag.StringVar(&dbName, "dbname", "", "数据库名称")
	flag.StringVar(&sql, "sql", "", "建表语句")
	flag.StringVar(&path, "path", "./model", "model生成地址")
	// 解析命令行参数
	flag.Parse()
	//读取配置文件
	config.InitConfig()

	var param string
	switch mode {
	case x_mode.DbMode:
		if dbName == "" {
			log.Error("db模式下，'dbname'不能为空")
			return
		}
		param = dbName

	case x_mode.StrMode:
		if sql == "" {
			log.Error("sql模式下，'sql'不能为空")
			return
		}
		param = sql
		log.Error("sql模式暂未支持")
		return
	default:
		log.Error("-mode 不能为空。执行模式 sql：建表语句解析 db:数据库解析")
		return
	}

	//文件路径校验以及转换
	tmplPath, filePath, packageName, err := file.PathVerify(path)
	if err != nil {
		log.Error(err.Error())
		return
	}

	tableList := x_mode.NewModeService(mode).DbTable(param)

	for _, t := range tableList {
		var importPackagesMap = make(map[string]struct{}, 0)
		astNode, err := parser.ParseSql(t)
		if err != nil {
			log.Error(err.Error())
			return
		}
		result := parser.ExtractCol(astNode)

		table := &file.Table{
			TableName:   result.TableName,
			GoTableName: util.Case2Camel(result.TableName),
			PackageName: packageName,
			Fields:      make([]*file.Column, 0, 0),
		}
		for _, col := range result.ColNames {
			if col.GoType == mysql.GoTypeTime {
				importPackagesMap[`"time"`] = struct{}{}
			}
			switch col.Name {
			case "is_del":
				col.GoType = "soft_delete.DeletedAt"
				col.GormTag = "softDelete:flag,DeletedAtField:DeletedTime"
				importPackagesMap[`"gorm.io/plugin/soft_delete"`] = struct{}{}
			case "create_time":
				col.GormTag = "autoCreateTime"
			case "modify_time":
				col.GormTag = "autoUpdateTime"
			case "deleted_time":
				col.GormTag = "deletedAt"
			}

			column := &file.Column{
				GoColumnName:  util.Case2Camel(col.Name),
				GoColumnType:  col.GoType,
				ColumnName:    col.Name,
				ColumnType:    col.Type,
				ColumnComment: col.Comment,
				GormTag:       col.GormTag,
			}
			table.Fields = append(table.Fields, column)
		}
		for name, _ := range importPackagesMap {
			table.ImportPackages = append(table.ImportPackages, name)
		}

		err = file.WriteToFileAndCreate(tmplPath, filePath, table)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("table to model successful", zap.String("table_name", table.TableName))
	}
	return
}
