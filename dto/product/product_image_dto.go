package product

type ProductImageDto struct {
	ID        uint   `json:"id" form:"id"`
	ProductID uint   `json:"product_id" form:"product_id"`
	ImageURL  string `json:"image_url" form:"image_url"`
}
