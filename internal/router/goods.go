package router

import (
	"MoocShop/internal/api/goods"
	"github.com/gin-gonic/gin"
)

func GoodsRouter(Router *gin.RouterGroup) {
	r := Router.Group("/goods")
	{
		r.POST("/create", goods.Create)
		r.GET("/delete/:gid", goods.Delete)
		r.POST("/update", goods.Update)
		r.GET("/detail/:gid", goods.Detail)
		r.GET("/list", goods.List)
	}
}
