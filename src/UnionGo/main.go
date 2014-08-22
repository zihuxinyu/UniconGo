package main

import (
	_ "UnionGo/routers"
	"github.com/astaxie/beego"
	"UnionGo/controllers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"reflect"
	"fmt"
)

func init() {


	dbUser, _ := beego.GetConfig("string", "db_user")
	dbPass, _ := beego.GetConfig("string", "db_pass")
	dbHost, _ := beego.GetConfig("string", "db_host")
	dbPort, _ := beego.GetConfig("string", "db_port")
	dbName, _ := beego.GetConfig("string", "db_name")
//	maxIdleConn, _ := beego.AppConfig.Int("db_max_idle_conn")
//	maxOpenConn, _ :=  beego.AppConfig.Int( "db_max_open_conn")
	maxIdleConn, _ := beego.GetConfig("int","db_max_idle_conn")
	maxOpenConn, _ := beego.GetConfig("int", "db_max_open_conn")
	beego.Debug("dd",reflect.TypeOf(maxIdleConn),reflect.TypeOf(maxOpenConn))
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn.(int), maxOpenConn.(int))


}

func main() {

	orm.Debug = true
	beego.SessionOn = true
	beego.AutoRouter(&controllers.MainController{})
	beego.Run()

}
