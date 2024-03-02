package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/shuqizhao/xcfg-server/src/models"

	"github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) GetRoles() {
	dtp := &models.DataTableParameter{}
	dtp.Unmarshal(c)
	roles, count := models.GetRolesByPage(dtp.DisplayStart, dtp.DisplayLength)

	jsonResult := models.DataTableResult{dtp.Echo, count, count, roles}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *RoleController) GetRole() {
	id, _ := c.GetInt("id")
	role := models.GetRole(id)
	menus := ""

	mrs := models.GetMenusByRoleId(role.Id)
	for i, mr := range mrs {
		if i == len(mrs)-1 {
			menus += mr.MenuName
		} else {
			menus += mr.MenuName + "<p/>"
		}
	}
	rrvm := models.RoleMenuViewModel{strconv.Itoa(role.Id), role.Name, menus, role.IsEnabled}
	jsonResult := models.JsonResult{Code: 200, Data: rrvm}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *RoleController) AddRole() {
	var role models.Role
	json.Unmarshal(c.Ctx.Input.RequestBody, &role)
	if role.Name == "" {
		c.ParseForm(&role)
	}
	models.AddRole(role.Name)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *RoleController) IsExists() {
	var role models.Role
	json.Unmarshal(c.Ctx.Input.RequestBody, &role)
	if role.Name == "" {
		c.ParseForm(&role)
	}
	data := 0
	if models.IsExistsRoleName(role.Name, role.Id) {
		data = 1
	}
	jsonResult := models.JsonResult{Code: 200, Data: data}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *RoleController) GetRoleMenus() {
	id, _ := c.GetInt("id")
	mrs := models.GetMenusByRoleId(id)
	right := []models.IdNameViewModel{}
	rightMap := make(map[int]string)
	for _, v := range mrs {
		right = append(right, models.IdNameViewModel{v.MenuId, v.MenuName})
		rightMap[v.MenuId] = v.MenuName
	}
	menus := models.GetResources()
	left := []models.IdNameViewModel{}
	for _, v := range menus {
		_, ok := rightMap[v.Id]
		if !ok {
			left = append(left, models.IdNameViewModel{v.Id, v.Name})
		}
	}
	s2s := models.SelectToSelectViewModel{left, right}
	c.Data["json"] = &s2s
	c.ServeJSON()
}

func (c *RoleController) UpdateRole() {
	var rm models.RoleMenuViewModel
	json.Unmarshal(c.Ctx.Input.RequestBody, &rm)
	if rm.Id == "" {
		c.ParseForm(&rm)
	}
	models.AddMenusByRoleId(rm.Id, rm.Menus)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}
