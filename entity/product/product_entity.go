package product

import "time"

type ProductEntity struct {
	ID          uint                 `gorm:"primaryKey;column:id"`
	SellerID    uint                 `gorm:"column:seller_id"`
	Name        string               `gorm:"column:name"`
	Description string               `gorm:"column:description"`
	Price       float64              `gorm:"column:price"`
	Quantity    int                  `gorm:"column:quantity"`
	CategoryID  uint                 `gorm:"column:category_id"`
	IsSold      bool                 `gorm:"column:is_sold"`
	CreatedAt   time.Time            `gorm:"column:created_at"`
	UpdatedAt   time.Time            `gorm:"column:updated_at"`
	Images      []ProductImageEntity `gorm:"foreignKey:ProductID"`
}
