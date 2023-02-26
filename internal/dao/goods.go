package dao

import (
	"MoocShop/internal/model"
	"fmt"
)

type GoodsDAO struct {
	model.Goods
}

// IGoods 商品接口
type IGoods interface {
	Create(*model.Goods) error
	Delete(int) error
	Update(int, map[string]interface{}) error
	SelectById(int) (model.Goods, error)
	SelectAll(int, int) ([]*model.Goods, error)
}

type GoodsManager struct {
	tableName string
}

func NewGoodsManager(tableName string) IGoods {
	gm := &GoodsManager{
		tableName: tableName,
	}
	return gm
}

// Create 创建商品
func (gm *GoodsManager) Create(goods *model.Goods) error {
	err := db.Table(gm.tableName).Create(&goods).Error
	return err
}

// Delete 删除商品
func (gm *GoodsManager) Delete(gid int) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = %s", gm.tableName, gid)
	err := db.Exec(sql).Error
	return err
}

// Update 更新商品
func (gm *GoodsManager) Update(gid int, m map[string]interface{}) (err error) {
	err = db.Table(gm.tableName).Where("id = ?", gid).Updates(m).Error
	return err
}

// SelectById 查询单个商品
func (gm *GoodsManager) SelectById(gid int) (model.Goods, error) {
	var mg model.Goods
	err := db.Table(gm.tableName).Where("id = ?", gid).Find(&mg).Error
	return mg, err
}

// SelectAll 查询所有商品
func (gm *GoodsManager) SelectAll(page int, size int) (goodsList []*model.Goods, err error) {
	err = db.Table(gm.tableName).Order("id asc").Scopes(Paginate(page, size)).Find(&goodsList).Error
	return goodsList, err
}
