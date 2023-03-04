package router

import (
	"MoocShop/internal/api/order"
	"github.com/gin-gonic/gin"
)

func OrderRouter(Router *gin.RouterGroup) {
	r := Router.Group("/order")
	{
		r.POST("/create", order.Create)
		//r.GET("/delete/:gid", goods.Delete)
		//r.POST("/update", goods.Update)
		//r.GET("/detail/:gid", goods.Detail)
		//r.GET("/list", goods.List)
	}
}
