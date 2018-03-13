package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego-pp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
