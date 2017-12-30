package models

import (
	"regexp"
	"fmt"
	"strconv"
	"strings"
	"time"
	"bytes"

	"github.com/astaxie/beego/orm"
)

type Cfg struct {
	Id           int
	AppName      string
	CfgName      string
	MajorVersion int
	MinorVersion int
	CfgFile      string `orm:"type(text)"`
	CfgType string
	ApolloTemplate string `orm:"type(text)"`
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
	CfgType string
	ApolloTemplate string
}

func IsExistsCfg(cfgName string, appName string, env string) bool {
	fmt.Println(env, cfgName)
	o := orm.NewOrm()
	qs := o.QueryTable("cfg").Filter("CfgName", cfgName).Filter("AppName", appName).Filter("Environment", env)
	return qs.Exist()
}

func AddCfg(cfgName string, appName string, cfgFile string, env string,apolloTemplate string,cfgType string) int64 {
	// el, err := LoadByXml(cfgFile)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// el.RemoveAttr("majorVersion")
	// el.RemoveAttr("minorVersion")
	// el.AddAttr("majorVersion", "1")
	// el.AddAttr("minorVersion", "1")
	b := []byte(cfgFile)
	reg := regexp.MustCompile(`\<\?.*?\?\>`)
	rep := []byte("${1}")
	cfgFile = string(reg.ReplaceAll(b, rep))
	xmlDoc, _ := LoadDocument(strings.NewReader(cfgFile))
	rootNode := xmlDoc.FirstChild().ToElement()
	rootNode.SetAttribute("majorVersion","1")
	rootNode.SetAttribute("minorVersion","1")
	buf := bytes.NewBufferString("")
	xmlDoc.Accept(NewSimplePrinter(buf, PrintStream))

	o := orm.NewOrm()
	var cfg Cfg
	cfg.CfgName = cfgName
	cfg.AppName = appName
	cfg.CfgFile = "<?xml version='1.0' encoding='utf-8' ?>" + buf.String()
	cfg.MajorVersion = 1
	cfg.MinorVersion = 1
	cfg.Environment = env
	cfg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	cfg.UpdateTime = cfg.CreateTime
	cfg.CfgType = cfgType
	cfg.ApolloTemplate = strings.Trim(apolloTemplate," ")
	id, err := o.Insert(&cfg)
	if err == nil {
		GetRedisClient().Put(strconv.FormatInt(id, 10), cfg.CfgFile, time.Hour*24*360)
		idKey := cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_idKey"
		minorKey := cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_minorKey"
		GetRedisClient().Put(strings.ToLower(idKey), strconv.FormatInt(id, 10), time.Hour*24*360)
		GetRedisClient().Put(strings.ToLower(minorKey), cfg.MinorVersion, time.Hour*24*360)
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
	idKey := strings.ToLower(cfgName + "_" + appName + "_" + env + "_idKey")
	minorKey := strings.ToLower(cfgName + "_" + appName + "_" + env + "_minorKey")

	idKeyNoAppName := strings.ToLower(cfgName + "_" + "_" + env + "_idKey")
	minorKeyNoAppName := strings.ToLower(cfgName + "_" + "_" + env + "_minorKey")

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
		return GetCfgIdByCfgName(cfgName, appName, env)
	}
}

func GetCfgIdByCfgName(cfgName string, appName string, env string) int {
	BuildCfgCacheByCfgName(cfgName, appName, env)
	o := orm.NewOrm()
	var cfgs []*Cfg
	count, _ := o.Raw("SELECT * from cfg WHERE lower(cfg_name) = lower(?) and (app_name = '' or lower(app_name)=lower(?)) and environment=? limit 2", cfgName, appName, env).QueryRows(&cfgs)

	if count == 1 {
		if cfgs[0].AppName != "" && strings.ToLower(cfgs[0].AppName) == appName {
			return cfgs[0].Id
		} else if cfgs[0].AppName == "" {
			return cfgs[0].Id
		}
	} else if count == 2 {
		for _, cfg := range cfgs {
			if strings.ToLower(cfg.AppName) == appName && cfg.AppName != "" {
				return cfg.Id
			}
		}
	}
	return 0
}

func BuildCfgCacheByCfgName(cfgName string, appName string, env string) {
	o := orm.NewOrm()
	var cfgs []*Cfg
	count, _ := o.Raw("SELECT * from cfg WHERE lower(cfg_name) = lower(?) and environment=? ", cfgName, env).QueryRows(&cfgs)

	if count > 0 {
		for _, cfg := range cfgs {
			BuildCfgCache(cfg, false)
		}
	} else {
		BuildCfgCache(&Cfg{CfgName: cfgName, Environment: env}, true)
	}
}

func BuildCfgCache(cfg *Cfg, init bool) {
	var idKey, minorKey string
	if cfg.AppName != "" {
		idKey = cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_idKey"
		minorKey = cfg.CfgName + "_" + cfg.AppName + "_" + cfg.Environment + "_minorKey"
	} else {
		idKey = cfg.CfgName + "_" + "_" + cfg.Environment + "_idKey"
		minorKey = cfg.CfgName + "_" + "_" + cfg.Environment + "_minorKey"
	}
	idKey = strings.ToLower(idKey)
	minorKey = strings.ToLower(minorKey)
	if init {
		GetRedisClient().Put(idKey, 0, time.Hour*24*360)
		GetRedisClient().Put(minorKey, 0, time.Hour*24*360)
	} else {
		GetRedisClient().Put(idKey, cfg.Id, time.Hour*24*360)
		GetRedisClient().Put(minorKey, cfg.MinorVersion, time.Hour*24*360)
		GetRedisClient().Put(strconv.Itoa(cfg.Id), cfg.CfgFile, time.Hour*24*360)
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
			GetRedisClient().Put(strconv.Itoa(cfg.Id), cfg.CfgFile, time.Hour*24*360)
		}
		return cfg.CfgFile
	}

}

func UpdateCfg(cfg CfgUpdateViewModel) bool {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(cfg.Id)
	AddCfgHistory(id)
	newcfg := Cfg{Id: id}
	if o.Read(&newcfg) == nil {
		// el, err := LoadByXml(cfg.CfgFile)
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		// newcfg.MinorVersion++
		// el.RemoveAttr("majorVersion")
		// el.RemoveAttr("minorVersion")
		// el.AddAttr("majorVersion", "1")
		// el.AddAttr("minorVersion", strconv.Itoa(newcfg.MinorVersion))
		// //fmt.Println(el.ToString())
		// newcfg.CfgFile = "<?xml version='1.0' encoding='utf-8' ?>" + el.ToString()

		newcfg.MinorVersion++
		b := []byte(cfg.CfgFile)
		reg := regexp.MustCompile(`\<\?.*?\?\>`)
		rep := []byte("${1}")
		cfg.CfgFile = string(reg.ReplaceAll(b, rep))
		xmlDoc, _ := LoadDocument(strings.NewReader(cfg.CfgFile))
		rootNode := xmlDoc.FirstChild().ToElement()
		rootNode.SetAttribute("majorVersion","1")
		rootNode.SetAttribute("minorVersion",strconv.Itoa(newcfg.MinorVersion))
		buf := bytes.NewBufferString("")
		xmlDoc.Accept(NewSimplePrinter(buf, PrintStream))
		//fmt.Println(buf.String())
		newcfg.CfgFile = "<?xml version='1.0' encoding='utf-8' ?>" + buf.String()
		
		newcfg.ApolloTemplate = strings.Trim(cfg.ApolloTemplate," ")

		newcfg.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		if _, err := o.Update(&newcfg, "cfg_file", "minor_version", "update_time","apollo_template"); err == nil {
			GetRedisClient().Put(cfg.Id, newcfg.CfgFile, time.Hour*24*360)
			idKey := newcfg.CfgName + "_" + newcfg.AppName + "_" + newcfg.Environment + "_idKey"
			minorKey := newcfg.CfgName + "_" + newcfg.AppName + "_" + newcfg.Environment + "_minorKey"
			GetRedisClient().Put(strings.ToLower(idKey), newcfg.Id, time.Hour*24*360)
			GetRedisClient().Put(strings.ToLower(minorKey), newcfg.MinorVersion, time.Hour*24*360)
			return true
		}
	}
	return false
}
