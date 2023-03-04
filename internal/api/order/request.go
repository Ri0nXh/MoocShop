package order

type (
	OrderCreate struct {
		GoodsId          int     `json:"goods_id"`
		UserId           int     `json:"user_id"`
		PayType          int     `json:"pay_type"`
		Status           int     `json:"status"`
		Price            float64 `json:"price"`
		CouponPrice      float64 `json:"coupon_price"`
		ActualPrice      float64 `json:"actual_price"`
		Remark           string  `json:"remark"`
		ConsigneeName    string  `json:"consignee_name"`
		ConsigneePhone   string  `json:"consignee_phone"`
		ConsigneeAddress string  `json:"consignee_address"`
	}
)
