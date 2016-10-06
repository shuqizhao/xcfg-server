package controllers

import "github.com/astaxie/beego"
import "xcfg-server/models"

type MenuController struct {
	beego.Controller
}

func (c *RoleController) GetMenus() {
	dtp := &models.DataTableParameter{}
	dtp.Unmarshal(c)
	menus, count := models.GetMenusByPage(dtp.DisplayStart, dtp.DisplayLength)
	jsonResult := models.DataTableResult{dtp.Echo, count, count, menus}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}
