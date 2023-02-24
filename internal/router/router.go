package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	g := gin.Default()
	ApiGroup := g.Group("/api/v1/")
	// 商品路由
	GoodsRouter(ApiGroup)

	// 用户路由
	UserRouter(ApiGroup)
	return g
}
