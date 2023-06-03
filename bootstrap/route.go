package bootstrap

import "github.com/gin-gonic/gin"

//  路由初始化

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	//  注册 API 路由
	router.RegisterAPIRoutes(router)
	//  配置 404 路由
	setup404Handler(router)
}

// 配置 404 路由
func setup404Handler(router *gin.Engine) {
 router.NoRoute(func(c *gin.Context) {

	 acceptString :=c.Request.Header.Get("Accept")
	 if strings.c
 })
}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery())
}
