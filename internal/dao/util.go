package dao

import "gorm.io/gorm"

// Paginate 数据库分页查询
func Paginate(page, size int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
