package router

import (
	"server/internal/controller"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()
var tr = gin.Default()

func CollectRoute() *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	r.LoadHTMLFiles("../html/index.html")
	r.Static("css", "../html/css")
	r.Static("js", "../html/js")
	r.Static("img", "../html/img")

	r.GET("/ping", controller.PingController)
	r.GET("/", controller.IndexController)
	r.POST("/api/user/login", controller.LoginController)

	return r
}

func TrCollectRoute() *gin.Engine {
	tr.Use(middleware.CORSMiddleware())

	tr.POST("/api/user/login", controller.LoginCertificateController)

	return tr
}
