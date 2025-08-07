package router

import (
	"github.com/gin-gonic/gin"

	"src/internal/server/http/controller"
)

func Setup() *gin.Engine {
	router := gin.New()

	// 中间件
	mws := []gin.HandlerFunc{}
	router.Use(mws...)

	// 路由注册
	controller.Register(router)

	return router
}
