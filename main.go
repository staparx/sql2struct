package main

import (
	"flag"
	"go.uber.org/zap"
)

var (
	mode              string
	dbName, path, sql string
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
	InitConfig()

	var param string

	modeType := ModeType(mode)

	switch modeType {

	case DbMode:
		if dbName == "" {
			log.Error("db模式下，'dbname'不能为空")
			return
		}
		param = dbName

	case SqlMode:
		if sql == "" {
			log.Error("sql模式下，'sql'不能为空")
			return
		}
		param = sql

	default:
		log.Error("-mode 不能为空。执行模式 sql：建表语句解析 db:数据库解析")
		return
	}

	//文件路径校验以及转换
	tmplPath, filePath, packageName, err := PathVerify(path)
	if err != nil {
		log.Error(err.Error())
		return
	}

	modeService, err := NewModeService(modeType, param)
	if err != nil {
		log.Error(err.Error())
		return
	}

	tableList := modeService.DbTable()

	for _, t := range tableList {
		fileStruct, err := ParseSQLToFileStruct(t, packageName)
		if err != nil {
			log.Error(err.Error())
			return
		}

		err = fileStruct.WriteToFileAndCreate(tmplPath, filePath)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("表结构体转换成功", zap.String("表名", fileStruct.TableName), zap.String("文件路径", filePath))
	}
	return
}
