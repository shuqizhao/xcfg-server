package routers

import (
	"xcfg-server/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/auth/checkpwd", &controllers.AuthController{}, "post:CheckPwd")
	beego.Router("/auth/mod", &controllers.AuthController{}, "post:ModPwd")
	beego.Router("/auth/menus", &controllers.AuthController{}, "get:GetMenus")

	beego.Router("/user/list", &controllers.UserController{}, "post:GetUsers")
	beego.Router("/user/disable", &controllers.UserController{}, "post:DisableUser")
	beego.Router("/user/enable", &controllers.UserController{}, "post:EnableUser")
	beego.Router("/user/add", &controllers.UserController{}, "post:AddUser")
	beego.Router("/user/exists", &controllers.UserController{}, "post:IsExists")
	beego.Router("/user/get", &controllers.UserController{}, "get:GetUser")
	beego.Router("/user/roles", &controllers.UserController{}, "get:GetUserRoles")
	beego.Router("/user/update", &controllers.UserController{}, "post:UpdateUser")

	beego.Router("/role/list", &controllers.RoleController{}, "post:GetRoles")
	beego.Router("/role/add", &controllers.RoleController{}, "post:AddRole")
	beego.Router("/role/exists", &controllers.RoleController{}, "post:IsExists")
	beego.Router("/role/get", &controllers.RoleController{}, "get:GetRole")
	beego.Router("/role/update", &controllers.RoleController{}, "post:UpdateRole")
	beego.Router("/role/resources", &controllers.RoleController{}, "get:GetRoleMenus")

	beego.Router("/menu/list", &controllers.RoleController{}, "post:GetMenus")

	beego.Router("/ConfigVersionHandler.ashx", &controllers.ConfigVersionController{})

	var FilterUser = func(ctx *context.Context) {
		adAuthCookie := ctx.Input.Cookie("adAuthCookie")
		if adAuthCookie != "true" && ctx.Request.RequestURI != "/" && ctx.Request.RequestURI != "/auth/login" {
			ctx.Redirect(302, "/")
		}

	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
}
