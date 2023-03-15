package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TableCreate struct {
	Table       string
	CreateTable string
}

type DBMysql struct {
	dbName string
}

func NewDBMysql(dbName string) *DBMysql {
	return &DBMysql{
		dbName: dbName,
	}
}

func (d *DBMysql) DbTable() []string {
	tableList, err := d.showCreateTables()
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	resp := make([]string, len(tableList), len(tableList))

	for index, table := range tableList {
		resp[index] = table.CreateTable
	}

	return resp
}

// ShowCreateTables 获取数据库所有表的建表语句
func (d *DBMysql) showCreateTables() ([]*TableCreate, error) {
	var result = make([]*TableCreate, 0, 0)
	dsn := d.dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// GORM v2 使用以下的方式关闭 DB
	defer func() {
		if sqlDB, err := db.DB(); err == nil && sqlDB != nil {
			_ = sqlDB.Close()
		}
	}()

	var tables []string
	err = db.Raw("SHOW TABLES;").Scan(&tables).Error
	if err != nil {
		return nil, err
	}

	for _, table := range tables {
		var t []map[string]interface{}
		sql := fmt.Sprintf("SHOW CREATE TABLE %s;", table)
		err = db.Raw(sql).Scan(&t).Error
		if err != nil {
			return nil, err
		}
		for _, v := range t {
			info := &TableCreate{}
			for key, value := range v {
				valueStr, ok := value.(string)
				if !ok {
					return nil, err
				}
				if key == "Table" {
					info.Table = valueStr
				} else {
					info.CreateTable = valueStr
				}
			}
			result = append(result, info)
		}
	}
	return result, nil
}

func (d *DBMysql) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Host,
		d.dbName)
}

type DBSqlStr struct {
	sql string
}

func NewDBSqlStr(sql string) *DBSqlStr {
	return &DBSqlStr{
		sql: sql,
	}
}
func (d *DBSqlStr) DbTable() []string {
	return []string{d.sql}
}
