package order

import (
	"MoocShop/internal/dao"
	"MoocShop/internal/model"
	"errors"
	"sync"
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

var m sync.Mutex

// Create 创建商品接口
func (o *OrderManager) Create(gid int, mo *model.Order) error {
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

	m.Lock()
	defer m.Unlock()
	tx := dao.DB.Begin()
	// 查询商品信息是否存在
	var gm model.Goods
	if result := tx.Table("goods").Where("id = ? and stock > 0", gid).Find(&gm); result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("未查询到此商品")
	}
	// 校验商品是否有库存
	//fmt.Println("商品库存 =》", gm.Stock)
	if gm.Stock > 0 {
		// 库存扣减1，销量增加1
		gm.Stock -= 1
		gm.Sale += 1

		// 创建订单
		if err := tx.Create(&mo).Error; err != nil {
			tx.Rollback()
			return errors.New("订单创建失败")
		}
		if err := tx.Table("order_goods_ref").Create(&ogr).Error; err != nil {
			tx.Rollback()
			return errors.New("订单商品关系表创建失败")
		}
		// 提交事务
		if err := tx.Select("Stock", "Sale").Updates(&gm).Error; err != nil {
			tx.Rollback()
			return errors.New("商品库存扣减失败")
		}
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return errors.New("提交失败")
		}
		return nil
	}
	tx.Rollback()
	return errors.New("此商品库存不足")
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
