package model

import "time"

// Goods 商品表
type Goods struct {
	Id        int       `json:"id"`
	Stock     int       `json:"stock"`
	Sale      int       `json:"sale"`
	Price     float64   `json:"price"`
	Name      string    `json:"name"`
	PicUrl    string    `json:"pic_url"`
	Brand     string    `json:"brand"`
	Tag       string    `json:"tag"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
