package main

import (
	"github.com/astaxie/beego"
	_ "github.com/beego-pp/routers"
)

func main() {
	beego.SetStaticPath("/static", "public")
	beego.Run()
}
