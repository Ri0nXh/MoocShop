package router

import "github.com/gin-gonic/gin"

func UserRouter(Router *gin.RouterGroup) {
	r := Router.Group("user")
	{
		r.GET("/list", func(ctx *gin.Context) {
			ctx.String(200, "user list ok")
		})
	}
}
