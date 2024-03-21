package router

import (
	"server/internal/controller"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func CollectRoute() *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	r.GET("/ping", controller.PingController)
	r.POST("/api/user/login", controller.LoginController)
	r.GET("/api/user/info", middleware.AuthMiddleware(), controller.InfoController)

	return r
}
