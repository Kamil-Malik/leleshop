package product

type ProductDto struct {
	ID          uint              `json:"id" form:"id"`
	SellerID    uint              `json:"seller_id" form:"seller_id"`
	Name        string            `json:"name" form:"name"`
	Description string            `json:"description" form:"description"`
	Price       float64           `json:"price" form:"price"`
	Quantity    int               `json:"quantity" form:"quantity"`
	CategoryID  uint              `json:"category_id" form:"category_id"`
	IsSold      bool              `json:"is_sold" form:"is_sold"`
	Images      []ProductImageDto `json:"images" form:"images"`
}
