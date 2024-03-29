package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
)
func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	beego.Router("/login", &controllers.UserController{},"get:Login")
	beego.Router("/users/session", &controllers.UserController{},"post:Session")
	beego.Router("/articles", &controllers.ArticleController{},"get:Articles")
	beego.Router("/articles", &controllers.ArticleController{},"post:CreateArticle")
	beego.Router("/articles/new", &controllers.ArticleController{},"get:New")

}
