package goods

import (
	"MoocShop/internal/dao"
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
	Update(model.Goods) error
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
	err := dao.DB.Table(gm.tableName).Create(&goods).Error
	return err
}

// Delete 删除商品
func (gm *GoodsManager) Delete(gid int) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = %d", gm.tableName, gid)
	err := dao.DB.Exec(sql).Error
	return err
}

// Update 更新商品
func (gm *GoodsManager) Update(mg model.Goods) (err error) {
	//err = dao.DB.Table(gm.tableName).Where("id = ?", gid).Updates(m).Error
	err = dao.DB.Table(gm.tableName).Where("id = ?", mg.Id).Save(mg).Error
	return err
}

// SelectById 查询单个商品
func (gm *GoodsManager) SelectById(gid int) (model.Goods, error) {
	var mg model.Goods
	err := dao.DB.Table(gm.tableName).Where("id = ?", gid).Find(&mg).Error
	return mg, err
}

// SelectAll 查询所有商品
func (gm *GoodsManager) SelectAll(page int, size int) (goodsList []*model.Goods, err error) {
	err = dao.DB.Table(gm.tableName).Order("id asc").Scopes(dao.Paginate(page, size)).Find(&goodsList).Error
	return goodsList, err
}
