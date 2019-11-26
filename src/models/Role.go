package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int
	Name       string
	IsDeleted  int
	IsEnabled  int
	CreateTime string
	CreateBy   int
	UpdateTime string
	UpdateBy   int
}

type RoleMenuViewModel struct {
	Id        string
	Name      string
	Menus     string
	IsEnabled int
}

func AddRole(name string) int64 {
	o := orm.NewOrm()
	var role Role
	role.Name = name
	role.IsDeleted = 0
	role.IsEnabled = 1
	role.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	role.UpdateTime = role.CreateTime
	id, err := o.Insert(&role)
	if err == nil {
		return id
	} else {
		return 0
	}
}

func GetRolesByPage(offset int64, length int64) ([]*Role, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("Role").Filter("is_deleted", "0").Filter("is_enabled", "1")

	var roles []*Role
	qs.Limit(length, offset).All(&roles)
	count, _ := qs.Count()
	return roles, count
}

func GetRoles() []*Role {
	o := orm.NewOrm()
	qs := o.QueryTable("Role").Filter("is_deleted", "0").Filter("is_enabled", "1")

	var roles []*Role
	qs.All(&roles)
	return roles
}

func IsExistsRoleName(rn string, id int) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("Role").Exclude("id", id).Filter("is_deleted", "0").Filter("is_enabled", "1").Filter("name__iexact", rn)
	return qs.Exist()
}

func GetRole(id int) Role {
	o := orm.NewOrm()
	role := Role{Id: id}
	o.Read(&role)
	return role
}
