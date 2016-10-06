package models

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RelRoleMenu struct {
	Id       int
	RoleId   int
	RoleName string
	MenuId   int
	MenuName string
}

func GetMenusByRoleId(id int) []*RelRoleMenu {
	o := orm.NewOrm()
	qs := o.QueryTable("RelRoleMenu").Filter("role_id", id)
	var relRoleMenus []*RelRoleMenu
	qs.All(&relRoleMenus)
	return relRoleMenus
}

func AddMenusByRoleId(id string, menus string) bool {
	menuids := strings.Split(menus, ",")
	o := orm.NewOrm()
	qs := o.QueryTable("RelRoleMenu").Filter("role_id", id)
	qs.Delete()
	for _, menuId := range menuids {
		if menuId == "" {
			continue
		}
		rm := new(RelRoleMenu)
		roleIdInt, _ := strconv.Atoi(id)
		rm.RoleId = roleIdInt
		menuIdInt, _ := strconv.Atoi(menuId)
		rm.MenuId = menuIdInt
		rm.MenuName = GetMenu(menuIdInt).Name
		o.Insert(rm)
	}
	return true
}
