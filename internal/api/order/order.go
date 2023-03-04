package order

import (
	"MoocShop/internal/api"
	"MoocShop/internal/dao/order"
	"MoocShop/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Create(c *gin.Context) {
	var req OrderCreate

	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("params error:" + err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}

	// 生成时间戳
	mo := model.Order{
		Number:           time.Now().UnixMilli(),
		Status:           req.Status,
		PayType:          req.PayType,
		Price:            req.Price,
		CouponPrice:      req.CouponPrice,
		ActualPrice:      req.ActualPrice,
		Remark:           req.Remark,
		ConsigneeName:    req.ConsigneeName,
		ConsigneeAddress: req.ConsigneeAddress,
		ConsigneePhone:   req.ConsigneePhone,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	fmt.Println(mo)
	om := order.NewOrderManager("order")
	err = om.Create(req.GoodsId, &mo)
	if err != nil {
		zap.L().Error("create order error: " + err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	api.RespSuccess(c, &mo)
}
