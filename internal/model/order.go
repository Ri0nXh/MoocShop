package model

import "time"

// Order 订单表
type Order struct {
	Id               int       `json:"id"`
	Number           int64     `json:"number"`
	UserId           int       `json:"user_id"`
	PayType          int       `json:"pay_type"`
	Status           int       `json:"status"`
	Price            float64   `json:"price"`
	CouponPrice      float64   `json:"coupon_price"`
	ActualPrice      float64   `json:"actual_price"`
	Remark           string    `json:"remark"`
	ConsigneeName    string    `json:"consignee_name"`
	ConsigneePhone   string    `json:"consignee_phone"`
	ConsigneeAddress string    `json:"consignee_address"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// OrderGoodsRef 订单商品关系表
type OrderGoodsRef struct {
	OrderId        int       `json:"order_id"`
	GoodsId        int       `json:"goods_id"`
	GoodsOptionsId int       `json:"goods_options_id"`
	Count          int       `json:"count"`
	Remark         string    `json:"remark"`
	Price          float64   `json:"price"`
	CouponPrice    float64   `json:"coupon_price"`
	ActualPrice    float64   `json:"actual_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
