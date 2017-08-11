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
	beego.AddViewPath("myviewpath")

	beego.Run()
}
