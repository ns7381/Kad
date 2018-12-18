package database

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var ORM *xorm.Engine

func init() {
	var err error
	ORM, err = xorm.NewEngine("mysql", "root:123456aB@(iiii)/service_ml?charset=utf8")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ORM.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	ORM.ShowSQL(true)
}