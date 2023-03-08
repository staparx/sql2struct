package parser

import (
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/test_driver"
	_ "github.com/pingcap/tidb/parser/test_driver"
	mysqlx "github.com/staparx/sql2struct/database/mysql"
)

func ParseSql(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

func ExtractCol(rootNode *ast.StmtNode) *ColX {
	v := &ColX{}
	(*rootNode).Accept(v)
	return v
}

type ColX struct {
	TableName string
	ColNames  []*Column
}

type Column struct {
	Name    string
	Type    string
	Comment string
	GoType  mysqlx.GoType
	GormTag string
}

func (v *ColX) Enter(in ast.Node) (ast.Node, bool) {
	if stmt, ok := in.(*ast.CreateTableStmt); ok {
		v.TableName = stmt.Table.Name.String()
		for _, col := range stmt.Cols {
			//类型转换为go的类型
			var goType mysqlx.GoType
			if col.Tp != nil {
				if t, ok := mysqlx.TypeMap[col.Tp.Tp]; ok {
					goType = t
				}
			}
			column := &Column{
				Name:    col.Name.String(),
				Type:    col.Tp.String(),
				Comment: "",
				GoType:  goType,
				GormTag: "",
			}

			for _, opt := range col.Options {
				switch opt.Tp {
				case ast.ColumnOptionComment:
					valueExpr, ok := opt.Expr.(*test_driver.ValueExpr)
					if ok {
						column.Comment = valueExpr.GetString()
					}
				}
			}
			v.ColNames = append(v.ColNames, column)
		}
	}
	return in, false
}

func (v *ColX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
