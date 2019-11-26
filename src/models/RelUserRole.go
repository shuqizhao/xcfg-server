package models

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RelUserRole struct {
	Id       int
	UserId   int
	UserName string
	RoleId   int
	RoleName string
}

func GetRolesByUserId(id int) []*RelUserRole {
	o := orm.NewOrm()
	qs := o.QueryTable("RelUserRole").Filter("user_id", id)
	var relUserRoles []*RelUserRole
	qs.All(&relUserRoles)
	return relUserRoles
}

func AddRolesByUserId(id string, roles string) bool {
	roleids := strings.Split(roles, ",")
	o := orm.NewOrm()
	qs := o.QueryTable("RelUserRole").Filter("user_id", id)
	qs.Delete()
	for _, roleId := range roleids {
		if roleId == "" {
			continue
		}
		ur := new(RelUserRole)
		userIdInt, _ := strconv.Atoi(id)
		ur.UserId = userIdInt
		roleIdInt, _ := strconv.Atoi(roleId)
		ur.RoleId = roleIdInt
		ur.RoleName = GetRole(roleIdInt).Name
		o.Insert(ur)
	}
	return true
}
