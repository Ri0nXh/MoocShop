package goods

type (
	ReqCreate struct {
		Stock  int     `json:"stock" binding:"required"`
		Price  float64 `json:"price" binding:"required"`
		Name   string  `json:"name" binding:"required"`
		PicUrl string  `json:"pic_url"`
		Brand  string  `json:"brand"`
		Tag    string  `json:"tag"`
		Detail string  `json:"detail"`
	}
)
