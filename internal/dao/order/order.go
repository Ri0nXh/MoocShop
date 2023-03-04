package order

import (
	"MoocShop/internal/dao"
	"MoocShop/internal/model"
	"errors"
	"time"
)

type OrderInfoIncludeGoods struct {
	model.Order
	GoodInfo model.Goods `json:"goods_info"`
}

type IOrder interface {
	Create(int, *model.Order) error
	Delete(int) error
	Update(model.Order) error
	SelectById(int) (model.Order, error)
	SelectAll(page int, size int) ([]*model.Order, error)
	SelectByIdIncludeGoods(int)
}
type OrderManager struct {
	TableName string
}

func NewOrderManager(tableName string) *OrderManager {
	return &OrderManager{
		TableName: tableName,
	}
}

// Create 创建商品接口
func (o *OrderManager) Create(gid int, mo *model.Order) error {

	//m.Lock()
	//defer m.Unlock()
	tx := dao.DB.Begin()
	sqlStr := "update goods set stock = stock - 1,sale = sale + 1 where stock > 0 and id = ?"
	exec := tx.Exec(sqlStr, gid)
	if err := exec.Error; err != nil {
		return err
	}
	if exec.RowsAffected > 0 {
		// 创建订单商品关系表
		ogr := model.OrderGoodsRef{
			GoodsId:        gid,
			OrderId:        mo.Id,
			GoodsOptionsId: 1,
			Count:          1,
			Remark:         mo.Remark,
			Price:          mo.Price,
			CouponPrice:    mo.CouponPrice,
			ActualPrice:    mo.ActualPrice,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		// 创建订单
		if err := tx.Create(&mo).Error; err != nil {
			tx.Rollback()
			return errors.New("订单创建失败")
		}
		if err := tx.Table("order_goods_ref").Create(&ogr).Error; err != nil {
			tx.Rollback()
			return errors.New("订单商品关系表创建失败")
		}
	}
	tx.Commit()
	return nil
}

func (o *OrderManager) Delete(i int) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderManager) Update(order model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderManager) SelectById(i int) (model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderManager) SelectAll(page int, size int) ([]*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderManager) SelectByIdIncludeGoods(i int) {
	//TODO implement me
	panic("implement me")
}
