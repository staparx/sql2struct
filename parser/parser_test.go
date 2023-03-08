package parser

import (
	"fmt"
	"testing"
)

func TestParseSql(t *testing.T) {
	sql := "CREATE TABLE `user`( `id` bigint(64) NOT NULL AUTO_INCREMENT, `name` varchar(100) DEFAULT NULL COMMENT '名称', PRIMARY KEY (`id`) USING BTREE) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4;"
	node, err := ParseSql(sql)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(node)
}
