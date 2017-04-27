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
		GetRedisClient().Put(strconv.FormatInt(id, 10), cfg.CfgFile, time.Hour*24*360)
		idKey := cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_idKey"
		minorKey := cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_minorKey"
		GetRedisClient().Put(idKey, strconv.FormatInt(id, 10), time.Hour*24*360)
		GetRedisClient().Put(minorKey, cfg.MinorVersion, time.Hour*24*360)
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

func GetCfgByMinorByRedis(cfgName string, appName string, env string, minor int) int {
	if minor < 0 {
		return 0
	}
	idKey := cfgName + "_" + appName + "_" + env + "_idKey"
	minorKey := cfgName + "_" + appName + "_" + env + "_minorKey"

	idKeyNoAppName := cfgName + "_" + "_" + env + "_idKey"
	minorKeyNoAppName := cfgName + "_" + "_" + env + "_minorKey"

	if GetRedisClient().IsExist(minorKey) {
		minorStr := string(GetRedisClient().Get(minorKey).([]uint8))
		minorInt, _ := strconv.Atoi(minorStr)
		fmt.Println("minorKey", cfgName, minor, minorInt)
		if minor < minorInt {
			idStr := string(GetRedisClient().Get(idKey).([]uint8))
			id, _ := strconv.Atoi(idStr)
			return id
		} else {
			return 0
		}
	} else if GetRedisClient().IsExist(minorKeyNoAppName) {
		minorStr := string(GetRedisClient().Get(minorKeyNoAppName).([]uint8))
		minorInt, _ := strconv.Atoi(minorStr)
		fmt.Println("minorKeyNoAppName", cfgName, minor, minorInt)
		if minor < minorInt {
			idStr := string(GetRedisClient().Get(idKeyNoAppName).([]uint8))
			id, _ := strconv.Atoi(idStr)
			return id
		} else {
			return 0
		}
	} else {
		return GetCfgByMinor(cfgName, appName, env, minor)
	}
}

func GetCfgByMinor(cfgName string, appName string, env string, minor int) int {

	o := orm.NewOrm()
	qs := o.QueryTable("Cfg")
	qs = qs.Filter("cfg_name", cfgName)
	if appName != "" {
		qs = qs.Filter("app_name", appName)
	}
	qs = qs.Filter("environment", env)
	qs = qs.Filter("minor_version__gte", minor)

	var cfgs []*Cfg
	qs.Limit(1).All(&cfgs)
	count, _ := qs.Count()
	//包含appname且版本匹配
	if count >= 1 {
		idKey := cfgs[0].CfgName + "_" + cfgs[0].AppName + "_" + env + "_idKey"
		minorKey := cfgs[0].CfgName + "_" + cfgs[0].AppName + "_" + env + "_minorKey"
		GetRedisClient().Put(idKey, cfgs[0].Id, time.Hour*24*360)
		GetRedisClient().Put(minorKey, cfgs[0].MinorVersion, time.Hour*24*360)
		return cfgs[0].Id
	} else {
		qs := o.QueryTable("Cfg")
		qs = qs.Filter("cfg_name", cfgName)
		if appName != "" {
			qs = qs.Filter("app_name", appName)
		}
		qs = qs.Filter("environment", env)
		count, _ := qs.Count()
		//包含appname但是版本达不到
		if count == 0 {
			qs := o.QueryTable("Cfg")
			if cfgName != "" {
				qs = qs.Filter("cfg_name", cfgName)
			}
			qs = qs.Filter("environment", env)
			qs = qs.Filter("minor_version__gte", minor)

			var cfgs []*Cfg
			qs.Limit(1).All(&cfgs)
			count, _ := qs.Count()
			//不包含appname但版本达到
			if count >= 1 {
				idKey := cfgs[0].CfgName + "_" + cfgs[0].AppName + "_" + env + "_idKey"
				minorKey := cfgs[0].CfgName + "_" + cfgs[0].AppName + "_" + env + "_minorKey"
				GetRedisClient().Put(idKey, cfgs[0].Id, time.Hour*24*360)
				GetRedisClient().Put(minorKey, cfgs[0].MinorVersion, time.Hour*24*360)
				return cfgs[0].Id
			} else {
				//不包含appname但版本达不到
				idKey := cfgName + "_" + "_" + env + "_idKey"
				minorKey := cfgName + "_" + "_" + env + "_minorKey"
				GetRedisClient().Put(idKey, 0, time.Hour*24*360)
				GetRedisClient().Put(minorKey, 0, time.Hour*24*360)
				return 0
			}
		} else {
			//包含appname但版本达不到
			idKey := cfgName + "_" + appName + "_" + env + "_idKey"
			minorKey := cfgName + "_" + appName + "_" + env + "_minorKey"
			GetRedisClient().Put(idKey, 0, time.Hour*24*360)
			GetRedisClient().Put(minorKey, 0, time.Hour*24*360)
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
	var key = strconv.Itoa(id)
	if GetRedisClient().IsExist(key) {
		return string(GetRedisClient().Get(key).([]uint8))
	} else {
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
		GetRedisClient().Put(strconv.Itoa(cfg.Id), cfg.CfgFile, time.Hour*24*360)
		return cfg.CfgFile
	}

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
			GetRedisClient().Put(cfg.Id, newcfg.CfgFile, time.Hour*24*360)
			idKey := newcfg.CfgName + "_" + newcfg.AppName + "_" + newcfg.Environment + "_idKey"
			minorKey := newcfg.CfgName + "_" + newcfg.AppName + "_" + newcfg.Environment + "_minorKey"
			GetRedisClient().Put(idKey, newcfg.Id, time.Hour*24*360)
			GetRedisClient().Put(minorKey, newcfg.MinorVersion, time.Hour*24*360)
			return true
		}
	}
	return false
}
