package controllers

import (
	"encoding/xml"
	"fmt"
	"xcfg-server/models"

	"github.com/astaxie/beego"
)

type ConfigVersionController struct {
	beego.Controller
}

func (c *ConfigVersionController) Post() {
	//fmt.Println(string(c.Ctx.Input.RequestBody))

	rcfg_result := models.RemoteConfigSectionCollection{}
	xml.Unmarshal(c.Ctx.Input.RequestBody, &rcfg_result)
	fmt.Println(rcfg_result.Machine, rcfg_result.Application)
	c.Data["xml"] = &rcfg_result
	c.ServeXML()
}
