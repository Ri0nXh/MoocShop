package goods

import (
	"MoocShop/internal/api"
	"MoocShop/internal/dao/goods"
	"MoocShop/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// Create 创建商品
func Create(c *gin.Context) {
	var req ReqCreate

	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}
	// 赋值给模型
	gmodel := model.Goods{
		Stock:     req.Stock,
		Price:     req.Price,
		Name:      req.Name,
		PicUrl:    req.PicUrl,
		Brand:     req.Brand,
		Tag:       req.Tag,
		Detail:    req.Detail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	gm := goods.NewGoodsManager("goods")
	err = gm.Create(&gmodel)
	if err != nil {
		zap.L().Info(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	api.RespSuccess(c, nil)
}

// Delete 删除商品
func Delete(c *gin.Context) {
	gIdStr := c.Param("gid")
	gid, err := strconv.Atoi(gIdStr)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}

	// 查询数据是否存在
	gm := goods.NewGoodsManager("goods")
	gmodel, err := gm.SelectById(gid)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	if gmodel.Id == 0 {
		zap.L().Info(fmt.Sprintf("data select failed by %d", gid))
		api.RespError(c, api.CodeServerBusy)
		return
	}

	err = gm.Delete(gmodel.Id)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	api.RespSuccess(c, nil)
}

// Update 更新商品
func Update(c *gin.Context) {
	var req ReqUpdate

	// 校验数据
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}

	gm := goods.NewGoodsManager("goods")
	gmodel, err := gm.SelectById(req.Id)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	if gmodel.Id == 0 {
		zap.L().Info(fmt.Sprintf("data select failed by %d", req.Id))
		api.RespError(c, api.CodeServerBusy)
		return
	}

	// 更新数据
	mg := model.Goods{
		Id:        req.Id,
		Stock:     req.Stock,
		Price:     req.Price,
		Name:      req.Name,
		PicUrl:    req.PicUrl,
		Brand:     req.Brand,
		Tag:       req.Tag,
		Detail:    req.Detail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = gm.Update(mg)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	api.RespSuccess(c, mg)
}

// Detail 商品详情
func Detail(c *gin.Context) {
	gIdStr := c.Param("gid")
	gid, err := strconv.Atoi(gIdStr)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}

	gm := goods.NewGoodsManager("goods")
	gmodel, err := gm.SelectById(gid)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	if gmodel.Id == 0 {
		zap.L().Info(fmt.Sprintf("data select failed by %d", gid))
		api.RespError(c, api.CodeServerBusy)
		return
	}
	api.RespSuccess(c, gmodel)
}

// List 获取商品列表
func List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "20")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeInvalidParam)
		return
	}

	// 返回数据
	gm := goods.NewGoodsManager("goods")
	goodsList, err := gm.SelectAll(page, size)
	if err != nil {
		zap.L().Error(err.Error())
		api.RespError(c, api.CodeServerBusy)
		return
	}
	data := map[string]interface{}{
		"total": len(goodsList),
		"goods": goodsList,
	}
	api.RespSuccess(c, data)
}
