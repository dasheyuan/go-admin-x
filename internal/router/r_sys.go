package routers

import (
	"github.com/gin-gonic/gin"
	"go-admin-x/internal/controllers/sys"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := sys.Menu{}
	app.GET("/menu/list", menu.List)
	app.GET("/menu/detail", menu.Detail)
	app.GET("/menu/allmenu", menu.AllMenu)
	app.GET("/menu/menubuttonlist", menu.MenuButtonList)
	app.POST("/menu/delete", menu.Delete)
	app.POST("/menu/update", menu.Update)
	app.POST("/menu/create", menu.Create)
	user := sys.User{}
	app.GET("/user/info", user.Info)
	app.POST("/user/login", user.Login)
	app.POST("/user/logout", user.Logout)
	app.POST("/user/editpwd", user.EditPwd)
	admin := sys.Admin{}
	app.GET("/admin/list", admin.List)
	app.GET("/admin/detail", admin.Detail)
	app.GET("/admin/adminroleidlist", admin.AdminsRoleIDList)
	app.POST("/admin/delete", admin.Delete)
	app.POST("/admin/update", admin.Update)
	app.POST("/admin/create", admin.Create)
	app.POST("/admin/setrole", admin.SetRole)
	role := sys.Role{}
	app.GET("/role/list", role.List)
	app.GET("/role/detail", role.Detail)
	app.GET("/role/rolemenuidlist", role.RoleMenuIDList)
	app.GET("/role/allrole", role.AllRole)
	app.POST("/role/delete", role.Delete)
	app.POST("/role/update", role.Update)
	app.POST("/role/create", role.Create)
	app.POST("/role/setrole", role.SetRole)
}
