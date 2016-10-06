package controllers

import "github.com/astaxie/beego"

//main controller
type MainController struct {
	beego.Controller
}

//index
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
