package main

import (
	_ "UnionGo/routers"
	"github.com/astaxie/beego"
	"UnionGo/controllers"
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var Cfg = beego.AppConfig

func init() {

	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"
	fmt.Println(dbLink)
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)


}

func main() {
	orm.Debug = true
	beego.SessionOn=true
	beego.AutoRouter(&controllers.MainController{})
	beego.Run()

}
