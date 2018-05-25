package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

// CMS API
type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}


// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

	this.Ctx.WriteString(this.AppController)
	this.Data["json"] = map[string]interface{}{"name": "jerry"}

	this.XSRFExpire = 7200
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	this.ServeJSON()

	this.Finish()

	this.StopRun()

}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {

}


