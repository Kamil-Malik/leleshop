package product

import "time"

type ProductImageEntity struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	ProductID uint      `gorm:"column:product_id"`
	ImageURL  string    `gorm:"column:image_url"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
