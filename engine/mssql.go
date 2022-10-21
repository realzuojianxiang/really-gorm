package engine

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/realzuojianxiang/really_go_orm"
)

func SqlServer(username string, password string, host string, port string, database string, timeout int) (*really_go_orm.Engine, error) {
	return really_go_orm.NewEngine("mssql", fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=%d",
		username,
		password,
		host,
		port,
		database,
		timeout,
	))
}
