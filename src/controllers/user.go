package controllers

import (
	"encoding/json"
	"strconv"
	"models"

	"github.com/astaxie/beego"
)

//user controller
type UserController struct {
	beego.Controller
}

func (c *UserController) GetUsers() {
	dtp := &models.DataTableParameter{}
	dtp.Unmarshal(c)
	usrIvm := &models.UserInputViewModel{}
	json.Unmarshal([]byte(dtp.Search), usrIvm)
	users, count := models.GetUsersByPage(dtp.DisplayStart, dtp.DisplayLength, *usrIvm)
	for _, usr := range users {
		usr.Pwd = "******"
	}
	jsonResult := models.DataTableResult{dtp.Echo, count, count, users}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) EnableUser() {
	var ids models.Ids
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	if len(ids.Ids)==0{
		c.ParseForm(&ids)
	}
	models.EnableUser(ids.Ids)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) DisableUser() {
	var ids models.Ids
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	if len(ids.Ids)==0{
		c.ParseForm(&ids)
	}
	models.DisableUser(ids.Ids)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) AddUser() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if user.Name==""{
		c.ParseForm(&user)
	}
	models.AddUser(user.Name, user.Pwd)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) UpdateUser() {
	var ur models.UserRoleViewModel
	json.Unmarshal(c.Ctx.Input.RequestBody, &ur)
	if ur.Id==""{
		c.ParseForm(&ur)
	}
	models.AddRolesByUserId(ur.Id, ur.Roles)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) IsExists() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if user.Name==""{
		c.ParseForm(&user)
	}
	data := 0
	if models.IsExistsUserName(user.Name, user.Id) {
		data = 1
	}
	jsonResult := models.JsonResult{Code: 200, Data: data}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}
func (c *UserController) GetUser() {
	id, _ := c.GetInt("id")
	user := models.GetUser(id)
	user.Pwd = "******"
	roles := ""

	urs := models.GetRolesByUserId(user.Id)
	for i, ur := range urs {
		if i == len(urs)-1 {
			roles += ur.RoleName
		} else {
			roles += ur.RoleName + ","
		}
	}
	urvm := models.UserRoleViewModel{strconv.Itoa(user.Id), user.Name, roles, user.IsEnabled}
	jsonResult := models.JsonResult{Code: 200, Data: urvm}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *UserController) GetUserRoles() {
	id, _ := c.GetInt("id")
	urs := models.GetRolesByUserId(id)
	right := []models.IdNameViewModel{}
	rightMap := make(map[int]string)
	for _, v := range urs {
		right = append(right, models.IdNameViewModel{v.RoleId, v.RoleName})
		rightMap[v.RoleId] = v.RoleName
	}
	roles := models.GetRoles()
	left := []models.IdNameViewModel{}
	for _, v := range roles {
		_, ok := rightMap[v.Id]
		if !ok {
			left = append(left, models.IdNameViewModel{v.Id, v.Name})
		}
	}
	s2s := models.SelectToSelectViewModel{left, right}
	c.Data["json"] = &s2s
	c.ServeJSON()
}
