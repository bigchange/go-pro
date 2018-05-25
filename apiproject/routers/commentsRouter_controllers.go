package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apiproject/controllers:CMSController"] = append(beego.GlobalControllerRouter["apiproject/controllers:CMSController"],
		beego.ControllerComments{
			Method: "AllBlock",
			Router: `/all/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:CMSController"] = append(beego.GlobalControllerRouter["apiproject/controllers:CMSController"],
		beego.ControllerComments{
			Method: "StaticBlock",
			Router: `/staticblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
