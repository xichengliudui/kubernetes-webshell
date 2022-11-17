package kubernetes_webshell

import (
	"github.com/astaxie/beego"
	_ "github.com/xichengliudui/kubernetes-webshell/routers"
)

func RunBeego() {
	beego.Run()
}
