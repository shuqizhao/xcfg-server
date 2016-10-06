package controllers

import (
	"encoding/json"
	"xcfg-server/models"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

//登录
func (c *AuthController) Login() {
	data := 0
	uservm := models.UserLoginViewModel{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &uservm)
	if user := models.Login(uservm.Name, uservm.Pwd); user.Id > 0 {
		data = 1
		c.Ctx.SetCookie("adAuthCookie", "true")
		c.Ctx.SetCookie("loginUser", user.Name)
		c.SetSession("userId", user.Id)
	}
	jsonResult := models.JsonResult{Code: 200, Data: data}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *AuthController) GetMenus() {
	un := c.Ctx.GetCookie("loginUser")
	if un == "admin" {
		menus := models.GetMenus()
		jsonResult := models.JsonResult{200, menus, ""}
		c.Data["json"] = &jsonResult
	} else {
		id := c.GetSession("userId").(int)
		roles := models.GetRolesByUserId(id)
		menuIds := []int{}
		for _, v := range roles {
			rms := models.GetMenusByRoleId(v.RoleId)
			for _, rm := range rms {
				menuIds = append(menuIds, rm.MenuId)
			}
		}
		menus := models.GetMenusByMenuIds(menuIds)
		jsonResult := models.JsonResult{200, menus, ""}
		c.Data["json"] = &jsonResult
	}
	c.ServeJSON()
}
