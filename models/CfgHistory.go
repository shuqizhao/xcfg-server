package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
	"time"
)

type CfgHistory struct {
	Id           int
	CfgId        int
	AppName      string
	CfgName      string
	MajorVersion int
	MinorVersion int
	CreateTime   string
	CreateBy     int
	UpdateTime   string
	UpdateBy     int
	CfgFile      string `orm:"type(text)"`
	Environment  string
}

func GetCfgHistorysByPage(offset int64, length int64, cfg CfgInputViewModel, cfgId string) ([]*CfgHistory, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("CfgHistory")
	if cfg.CfgName != "" {
		qs = qs.Filter("cfg_name__icontains", cfg.CfgName)
	}
	if cfg.AppName != "" {
		qs = qs.Filter("app_name__icontains", cfg.AppName)
	}
	if cfg.CreateTime != "" {
		qs = qs.Filter("create_time__gte", cfg.CreateTime)
	}
	if cfg.CreateTime1 != "" {
		qs = qs.Filter("create_time__lte", cfg.CreateTime1)
	}

	if cfg.UpdateTime != "" {
		qs = qs.Filter("update_time__gte", cfg.UpdateTime)
	}
	if cfg.UpdateTime1 != "" {
		qs = qs.Filter("update_time__lte", cfg.UpdateTime1)
	}
	qs = qs.Filter("cfg_id",cfgId)

	qs = qs.OrderBy("-id")

	var cfgs []*CfgHistory
	qs.Limit(length, offset).All(&cfgs)
	count, _ := qs.Count()
	return cfgs, count
}


func AddCfgHistory(cfgId int) int64 {
	o := orm.NewOrm()
	newcfg := Cfg{Id: cfgId}
	if o.Read(&newcfg) == nil {
		var cfgHistroy CfgHistory
		cfgHistroy.CfgId = cfgId
		cfgHistroy.CfgName = newcfg.CfgName
		cfgHistroy.AppName = newcfg.AppName
		cfgHistroy.CfgFile = newcfg.CfgFile
		cfgHistroy.MajorVersion = newcfg.MajorVersion
		cfgHistroy.MinorVersion = newcfg.MinorVersion
		cfgHistroy.Environment = newcfg.Environment
		cfgHistroy.CreateTime = newcfg.CreateTime
		cfgHistroy.UpdateTime = newcfg.UpdateTime
		id, err := o.Insert(&cfgHistroy)
		if err == nil {
			return id
		} else {
			return 0
		}
	}
	return 0
}

func RollBack(cfgHistoryId string) int64 {
	id, _ := strconv.Atoi(cfgHistoryId)
	o := orm.NewOrm()
	cfgHistory := CfgHistory{Id: id}
	if o.Read(&cfgHistory) == nil {
		idStr := strconv.Itoa(cfgHistory.CfgId)
		cfgUpdateViewModel:=CfgUpdateViewModel{}
		cfgUpdateViewModel.Id = idStr
		cfgUpdateViewModel.CfgFile = cfgHistory.CfgFile
		UpdateCfg(cfgUpdateViewModel)
	}
	return 0
}

func GetCfgHistoryFile(id int) string {
	o := orm.NewOrm()
	cfg := CfgHistory{Id: id}
	err := o.Read(&cfg)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(cfg.Id, cfg.CfgName)
		GetRedisClient().Put(strconv.Itoa(cfg.Id), cfg.CfgFile, time.Hour*24*360)
	}
	return cfg.CfgFile

}