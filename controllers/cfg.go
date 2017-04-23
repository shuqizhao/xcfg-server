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
		id := models.GetCfgByMinor(v.SectionName, rcfgResult.Application, rcfgResult.Environment, v.MinorVersion)
		if id > 0 {
			v.DownloadUrl = "xcfg/get?id=" + strconv.Itoa(id)
			sections = append(sections, &models.RemoteConfigSection{v.SectionName, 0, 0, v.DownloadUrl})
		}

	}
	rcfgResult.Sections = sections
	fmt.Println(rcfgResult.Machine, rcfgResult.Application)
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

func (c *CfgController) AddCfg() {
	env := c.GetString("env")
	var cfg models.Cfg
	json.Unmarshal(c.Ctx.Input.RequestBody, &cfg)
	models.AddCfg(cfg.CfgName, cfg.AppName, cfg.CfgFile, env)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgController) Exists() {
	env := c.GetString("env")
	fmt.Println(env)
	var cfg models.Cfg
	json.Unmarshal(c.Ctx.Input.RequestBody, &cfg)
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
	//fmt.Println(cfg.Id, "我的id", cfg.CfgFile)
	models.UpdateCfg(cfg)
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}
