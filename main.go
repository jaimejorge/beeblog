package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"

	"beeblog/configs"
	_ "beeblog/docs"
	_ "beeblog/routers"
)

func init() {
	user := beego.AppConfig.String("DB::DB_USER")
	password := beego.AppConfig.String("DB::DB_PASSWORD")
	host := beego.AppConfig.String("DB::DB_HOST")
	port := beego.AppConfig.String("DB::DB_PORT")
	db := beego.AppConfig.String("DB::DB_DATABASE")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)

	orm.RegisterDataBase("default", "mysql", dataSource)

	configs.Init()
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
