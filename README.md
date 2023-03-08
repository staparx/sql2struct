# Mysql表结构转换为Golang结构体

## 项目介绍
个人用于将mysql转换成结构体使用。

通过解析mysql的建表语句，将其转换成Golang中的结构体。

## 依赖
* 配置读取 - [viper](github.com/spf13/viper)
* SQL解析 - [tidb/parser](github.com/pingcap/tidb)
* ORM框架 - [gorm](gorm.io/gorm)


## 前置准备

将编译好的项目文件 sql2struct.go、config.yaml、model.tmpl 放入到 `$GOPATH/bin`路径下

config.yaml中配置好DSN的相关内容

## 运行方式
```shell
sql2struct -mode db -dbname database_name -path ./model
```
### 命令参数
```text
Usage of sql2struct:
  -dbname string
        数据库名称
  -mode string
        执行模式 sql:建表语句解析 db:数据库解析 (default "db")
  -path string
        model生成地址 (default "./model")
  -sql string
        建表语句
```
### TODO
* ✅ 数据库解析
* ❌ 建表语句SQL解析
