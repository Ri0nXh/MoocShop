package goods

import (
	"MoocShop/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var req ReqCreate

	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}
	api.RespSuccessWithoutData(c)
}
