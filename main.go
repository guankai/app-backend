package main

import (
	_ "app-backend/routers"

	"github.com/astaxie/beego"
	"app-backend/models/db"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	db.InitDatabase()
	orm.RegisterDataBase("default", "mysql", "root:.@127.0.0.1/app-backend?charset=utf8", 30)
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
