package engine

import (
	"github.com/SAP/go-ase"
	"github.com/SAP/go-dblib/dsn"
	"github.com/realzuojianxiang/really_go_orm"
)

func Ase(username string, password string, host string, port string, database string, timeout int) (*really_go_orm.Engine, error) {
	info, _ := ase.NewInfo()
	info.Username = username
	info.Password = password
	info.Host = host
	info.Port = port
	info.Database = database
	info.PacketReadTimeout = timeout
	return really_go_orm.NewEngine("ase", dsn.FormatSimple(info))

}
