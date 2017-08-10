// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	// namespace
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/userInfo/:id",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 基础路由
	beego.Get("/",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Post("/alice",func(ctx *context.Context){
		ctx.Output.Body([]byte("bob"))
	})

	beego.Any("/foo",func(ctx *context.Context){
		ctx.Output.Body([]byte("bar"))
	})

	// 固定路由
	beego.Router("/api", &controllers.UserController{})

	// 正则路由
	beego.Router("/api/:id", &controllers.ObjectController{})

	// 自定义方法及 RESTful 规则
	beego.Router("/api/list",&controllers.ObjectController{},"*:ListFood")
	beego.Router("/api/create",&controllers.ObjectController{},"get,post:CreateFood")
	beego.Router("/api/update",&controllers.ObjectController{},"put:UpdateFood;get:CreateFood")
	beego.Router("/api/delete",&controllers.ObjectController{},"delete:DeleteFood")


	// 自动匹配 (/:controller/:method )
	beego.AutoRouter(&controllers.AutoController{})

	// this.Ctx.Input.Param(":ext") 获取后缀名

	// 注解路由
	// 只会在 dev 模式下进行生成，生成的路由放在 “/routers/commentsRouter.go” 文件中
	beego.Include(&controllers.CMSController{})

	// 其实效果和自己通过 Router 函数注册是一样的
	beego.Router("/staticblock/:key", &controllers.CMSController{}, "get:StaticBlock")
	beego.Router("/all/:key", &controllers.CMSController{}, "get:AllBlock")


	/**
	 * GET /v1/notallowed
	 * GET /v1/version
   * GET /v1/changepassword
   * POST /v1/changepassword
	 * GET /v1/shop/123
	 * GET /v1/cms/ 对应 MainController、CMSController、BlockController 中得注解路由
	 *
	 */
	//初始化 namespace
	ns1 :=
		beego.NewNamespace("/v1",
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.Domain() == "api.beego.me" {
					return true
				}
				return false
			}),
			// beego.NSBefore(auth),
			beego.NSGet("/notallowed", func(ctx *context.Context) {
				ctx.Output.Body([]byte("notAllowed"))
			}),
			beego.NSRouter("/version", &controllers.ObjectController{}, "get:ShowAPIVersion"),
			beego.NSRouter("/changepassword", &controllers.ObjectController{}),
			beego.NSNamespace("/shop",
				// beego.NSBefore(sentry),
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("notAllowed"))
				}),
			),
			beego.NSNamespace("/cms",
				beego.NSInclude(
					&controllers.ObjectController{},
					&controllers.CMSController{},
					&controllers.ObjectController{},
				),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns1)




}
