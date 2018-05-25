package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)


func init()  {

	/**
		<form method="post" ...>
  	<input type="hidden" name="_method" value="put" />
	 */
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method")!="" && ctx.Input.IsPost() {
			// 修改请求的method
			ctx.Request.Method = ctx.Input.Query("_method")
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

}
