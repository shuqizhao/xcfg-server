package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Cfg struct {
	Id           int
	AppName      string
	CfgName      string
	MajorVersion int
	MinorVersion int
	CfgFile      string `orm:"type(text)"`
	Environment  string
	CreateTime   string
	CreateBy     int
	UpdateTime   string
	UpdateBy     int
}

type CfgInputViewModel struct {
	Id           int
	AppName      string
	CfgName      string
	MajorVersion int
	MinorVersion int
	CfgFile      string
	Environment  string
	CreateTime   string
	CreateTime1  string
	CreateBy     int
	UpdateTime   string
	UpdateTime1  string
	UpdateBy     int
}

type CfgUpdateViewModel struct {
	Id      string
	CfgFile string
}

func IsExistsCfg(cfgName string, appName string, env string) bool {
	fmt.Println(env, cfgName)
	o := orm.NewOrm()
	qs := o.QueryTable("cfg").Filter("CfgName", cfgName).Filter("AppName", appName).Filter("Environment", env)
	return qs.Exist()
}

func AddCfg(cfgName string, appName string, cfgFile string, env string) int64 {
	el, err := LoadByXml(cfgFile)
	if err != nil {
		fmt.Println("err", err)
	}
	el.RemoveAttr("majorVersion")
	el.RemoveAttr("minorVersion")
	el.AddAttr("majorVersion", "1")
	el.AddAttr("minorVersion", "1")

	o := orm.NewOrm()
	var cfg Cfg
	cfg.CfgName = cfgName
	cfg.AppName = appName
	cfg.CfgFile = "<?xml version='1.0' encoding='utf-8' ?>" + el.ToString()
	cfg.MajorVersion = 1
	cfg.MinorVersion = 1
	cfg.Environment = env
	cfg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	cfg.UpdateTime = cfg.CreateTime
	id, err := o.Insert(&cfg)
	if err == nil {
		return id
	} else {
		return 0
	}
}

func GetCfgsByPage(offset int64, length int64, cfg CfgInputViewModel, env string) ([]*Cfg, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("Cfg")
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

	qs = qs.Filter("environment", env)

	var cfgs []*Cfg
	qs.Limit(length, offset).All(&cfgs)
	count, _ := qs.Count()
	return cfgs, count
}

func GetCfgByMinor(cfgName string, appName string, env string, minor int) int {
	o := orm.NewOrm()
	qs := o.QueryTable("Cfg")
	qs = qs.Filter("cfg_name", cfgName)
	if appName != "" {
		qs = qs.Filter("app_name", appName)
	}
	qs = qs.Filter("environment", env)
	qs = qs.Filter("minor_version__gt", minor)

	var cfgs []*Cfg
	qs.Limit(1).All(&cfgs)
	count, _ := qs.Count()

	if count >= 1 {
		return cfgs[0].Id
	} else {
		qs := o.QueryTable("Cfg")
		if cfgName != "" {
			qs = qs.Filter("cfg_name", cfgName)
		}
		qs = qs.Filter("environment", env)
		qs = qs.Filter("minor_version__gt", minor)

		var cfgs []*Cfg
		qs.Limit(1).All(&cfgs)
		count, _ := qs.Count()
		if count >= 1 {
			return cfgs[0].Id
		} else {
			return 0
		}
	}
}

func GetCfg(id int) Cfg {
	o := orm.NewOrm()
	cfg := Cfg{Id: id}
	err := o.Read(&cfg)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(cfg.Id, cfg.CfgName)
	}
	return cfg
}

func GetCfgFile(id int) string {
	o := orm.NewOrm()
	cfg := Cfg{Id: id}
	err := o.Read(&cfg)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(cfg.Id, cfg.CfgName)
	}
	return cfg.CfgFile
}

func UpdateCfg(cfg CfgUpdateViewModel) bool {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(cfg.Id)
	newcfg := Cfg{Id: id}
	if o.Read(&newcfg) == nil {
		el, err := LoadByXml(cfg.CfgFile)
		if err != nil {
			fmt.Println("err", err)
		}
		newcfg.MinorVersion++
		el.RemoveAttr("majorVersion")
		el.RemoveAttr("minorVersion")
		el.AddAttr("majorVersion", "1")
		el.AddAttr("minorVersion", strconv.Itoa(newcfg.MinorVersion))
		//fmt.Println(el.ToString())
		newcfg.CfgFile = "<?xml version='1.0' encoding='utf-8' ?>" + el.ToString()

		newcfg.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		if _, err := o.Update(&newcfg, "cfg_file", "minor_version", "update_time"); err == nil {
			return true
		}
	}
	return false
}
