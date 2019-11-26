package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"models"
)

type CfgHistroyController struct {
	beego.Controller
}
func (c *CfgHistroyController) GetCfgHistorys() {
	id := c.GetString("id")
	dtp := &models.DataTableParameter{}
	dtp.Unmarshal(c)
	cfgIvm := &models.CfgInputViewModel{}
	json.Unmarshal([]byte(dtp.Search), cfgIvm)
	cfgs, count := models.GetCfgHistorysByPage(dtp.DisplayStart, dtp.DisplayLength, *cfgIvm, id)
	for _, cfg := range cfgs {
		cfg.CfgFile = ""
	}
	jsonResult := models.DataTableResult{dtp.Echo, count, count, cfgs}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgHistroyController) RollBack() {
	var ids models.Ids
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	if len(ids.Ids)==0{
		c.ParseForm(&ids)
	}
	models.RollBack(ids.Ids[0])
	jsonResult := models.JsonResult{Code: 200, Data: true}
	c.Data["json"] = &jsonResult
	c.ServeJSON()
}

func (c *CfgHistroyController) Get() {
	id, _ := c.GetInt("id")
	cfgFile := models.GetCfgHistoryFile(id)
	c.Ctx.Output.ContentType("application/xml")
	c.Ctx.Output.Body([]byte(cfgFile))
}

func (c *CfgHistroyController) GetTemplate() {
	id, _ := c.GetInt("id")
	cftTemplate := models.GetCfgHistoryTemplate(id)
	c.Ctx.Output.ContentType("text/plain")
	c.Ctx.Output.Body([]byte(cftTemplate))
}