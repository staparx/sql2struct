package mysql

import (
	"fmt"

	"github.com/staparx/sql2struct/config"

	"github.com/staparx/sql2struct/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMysql struct {
	dbName string
}

func NewDBMysql(dbName string) *DBMysql {
	return &DBMysql{
		dbName: dbName,
	}
}

func (d *DBMysql) GetTableDescribeList() (map[string][]*database.TableDescribe, error) {
	result := make(map[string][]*database.TableDescribe)
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
		var tableDescribeList = make([]*database.TableDescribe, 0, 0)
		sql := fmt.Sprintf("SHOW FULL COLUMNS FROM %s;", table)
		err = db.Raw(sql).Scan(&tableDescribeList).Error
		if err != nil {
			return nil, err
		}
		result[table] = tableDescribeList
	}

	return result, nil
}

// ShowCreateTables 获取数据库所有表的建表语句
func (d *DBMysql) ShowCreateTables() ([]*database.TableCreate, error) {
	var result = make([]*database.TableCreate, 0, 0)
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
			info := &database.TableCreate{}
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
		config.Cfg.Db.Username,
		config.Cfg.Db.Password,
		config.Cfg.Db.Host,
		d.dbName)
}
