package main

import (
	_ "apiproject/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SessionConfig{SessionOn:true}
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.AddViewPath("myviewpath")

	// beego.BConfig.AppName="beepkg"

	// beego.AppConfig.String("mysqluser")
	// beego.AppConfig.String("mysqlpass")
	// beego.AppConfig.String("mysqlurls")
	// beego.AppConfig.String("mysqldb")

	beego.Run()
}
