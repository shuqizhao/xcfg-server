package main

import (
	"xcfg-server/models"
	_ "xcfg-server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Role))
	orm.RegisterModel(new(models.Menu))
	orm.RegisterModel(new(models.RelUserRole))
	orm.RegisterModel(new(models.RelRoleMenu))
	orm.RegisterModel(new(models.Cfg))
	orm.RegisterModel(new(models.CfgHistory))

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(10.211.55.4:3306)/God?charset=utf8", 30, 30)
}

func main() {
	orm.Debug = true
	createTable()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

//自动建表
func createTable() {

	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}
