package main

import (
  "fmt"
	_ "beeblog/docs"
	_ "beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
  user := beego.AppConfig.String("DB::DB_USER")
  password := beego.AppConfig.String("DB::DB_PASSWORD")
  host := beego.AppConfig.String("DB::DB_HOST")
  port := beego.AppConfig.String("DB::DB_PROT")
  db := beego.AppConfig.String("DB::DB_DATABASE")
  dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

