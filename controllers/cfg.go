package controllers

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"xcfg-server/models"

	"github.com/astaxie/beego"
)

type CfgController struct {
	beego.Controller
}

/*
	查询配置文件状态
*/
func (c *CfgController) Post() {
	rcfgResult := models.RemoteConfigSectionCollection{}
	xml.Unmarshal(c.Ctx.Input.RequestBody, &rcfgResult)
	var sections = make([]*models.RemoteConfigSection, 0)
	for _, v := range rcfgResult.Sections {
		id := models.GetCfgByMinorByRedis(v.SectionName, rcfgResult.Application, rcfgResult.Environment, v.MinorVersion)
		if id > 0 {
			v.DownloadUrl = "xcfg/get?id=" + strconv.Itoa(id)
			v.TemplateUrl = "xcfg/getTemplate?id=" + strconv.Itoa(id)
			sections = append(sections, &models.RemoteConfigSection{v.SectionName, 0, 0, v.DownloadUrl,v.TemplateUrl})
		}

	}
	rcfgResult.Sections = sections
	//fmt.Println(rcfgResult.Machine, rcfgResult.Application)
	c.Data["xml"] = &rcfgResult
	c.ServeXML()
}

/*
	获取配置文件
*/
func (c *CfgController) Get() {
	id, _ := c.GetInt("id")
	cfgFile := models.GetCfgFile(id)
	c.Ctx.Output.ContentType("application/xml")
	c.Ctx.Output.Body([]byte(cfgFile))
}

/*
	获取模板
*/
func (c *CfgController) GetTemplate() {
	id, _ := c.GetInt("id")
	cfg := models.GetCfg(id)
	c.Ctx.Output.ContentType("text/plain")
	c.Ctx.Output.Body([]byte(cfg.ApolloTemplate))
}

func (c *CfgController) AddCfg() {
	env := c.GetString("env")
	cfgType := c.GetString("type")
	var cfg models.Cfg
	json.Unmarshal(c.Ctx.Input.RequestBody, &cfg)
	if cfg.CfgName==""{
		c.ParseForm(&cfg)
	}
	models.AddCfg(cfg.CfgName, cfg.AppName, cfg.CfgFile, env,cfg.ApolloTemplate,cfgType)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgController) Exists() {
	env := c.GetString("env")
	var cfg models.Cfg
	json.Unmarshal(c.Ctx.Input.RequestBody, &cfg)
	if cfg.Id==0{
		c.ParseForm(&cfg)
	}
	data := 0
	if models.IsExistsCfg(cfg.CfgName, cfg.AppName, env) {
		data = 1
	}
	jsonResult := models.JsonResult{Code: 200, Data: data}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgController) GetCfgs() {
	env := c.GetString("env")
	dtp := &models.DataTableParameter{}
	dtp.Unmarshal(c)
	cfgIvm := &models.CfgInputViewModel{}
	json.Unmarshal([]byte(dtp.Search), cfgIvm)
	fmt.Println(env)
	cfgs, count := models.GetCfgsByPage(dtp.DisplayStart, dtp.DisplayLength, *cfgIvm, env)
	for _, cfg := range cfgs {
		cfg.CfgFile = ""
	}
	jsonResult := models.DataTableResult{dtp.Echo, count, count, cfgs}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgController) GetCfg() {
	id, _ := c.GetInt("id")
	cfg := models.GetCfg(id)
	cfgfileBytes := []byte(cfg.CfgFile)
	encodeString := base64.StdEncoding.EncodeToString(cfgfileBytes)
	cfg.CfgFile = encodeString
	apolloTemplateBytes := []byte(cfg.ApolloTemplate)
	encodeApolloTemplateString := base64.StdEncoding.EncodeToString(apolloTemplateBytes)
	cfg.ApolloTemplate = encodeApolloTemplateString
	jsonResult := models.JsonResult{Code: 200, Data: cfg}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

/*
	更新配置
*/
func (c *CfgController) UpdateCfg() {
	var cfg models.CfgUpdateViewModel
	json.Unmarshal(c.Ctx.Input.RequestBody, &cfg)
	if cfg.Id==""{
		c.ParseForm(&cfg)
	}
	//fmt.Println(cfg.Id, "我的id", cfg.CfgFile)
	models.UpdateCfg(cfg)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}
