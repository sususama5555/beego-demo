package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Get")
    beego.Router("/test", &controllers.MainController{}, "get:Test")
}
