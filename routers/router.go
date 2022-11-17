package routers

import (
	"github.com/astaxie/beego"
	"github.com/xichengliudui/kubernetes-webshell/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/terminal", &controllers.TerminalController{}, "get:Get")
	beego.Handler("/terminal/ws", &controllers.TerminalSockjs{}, true)
}
