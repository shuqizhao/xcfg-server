package models

import "github.com/astaxie/beego/orm"

type Menu struct {
	Id         int
	Name       string
	Url        string
	ParentId   int
	IsDeleted  int
	IsEnabled  int
	CreateTime string
	CreateBy   int
	UpdateTime string
	UpdateBy   int
}

func GetMenusByPage(offset int64, length int64) ([]*Menu, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("Menu").Filter("is_deleted", "0").Filter("is_enabled", "1")

	var menus []*Menu
	qs.Limit(length, offset).All(&menus)
	count, _ := qs.Count()
	return menus, count
}

func GetMenusByMenuIds(ids []int) []*Menu {
	o := orm.NewOrm()

	qs := o.QueryTable("Menu").Filter("is_deleted", "0").Filter("is_enabled", "1").Filter("id__in", ids)
	var menus []*Menu
	qs.All(&menus)
	return menus
}

func GetMenu(id int) Menu {
	o := orm.NewOrm()
	menu := Menu{Id: id}
	o.Read(&menu)
	return menu
}

func GetMenus() []*Menu {
	o := orm.NewOrm()
	qs := o.QueryTable("Menu").Filter("is_deleted", "0").Filter("is_enabled", "1")

	var menus []*Menu
	qs.All(&menus)
	return menus
}
