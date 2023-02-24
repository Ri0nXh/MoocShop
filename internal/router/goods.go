package router

import "github.com/gin-gonic/gin"

func GoodsRouter(Router *gin.RouterGroup) {
	r := Router.Group("goods")
	{
		r.GET("/list", func(ctx *gin.Context) {
			ctx.String(200, "goods list ok")
		})
	}
}
