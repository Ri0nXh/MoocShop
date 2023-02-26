package router

import (
	"MoocShop/internal/api/goods"
	"github.com/gin-gonic/gin"
)

func GoodsRouter(Router *gin.RouterGroup) {
	r := Router.Group("/goods")
	{
		r.POST("/create", goods.Create)
	}
}
