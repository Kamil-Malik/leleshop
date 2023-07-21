package user

import (
	"leleshop/entity/product"
	"time"
)

type UserEntity struct {
	Id             uint                    `gorm:"primaryKey, column:id"`
	UserName       string                  `gorm:"index, column:user_name"`
	FullName       string                  `gorm:"column:full_name"`
	FcmToken       string                  `gorm:"column:fcm_token"`
	Email          string                  `gorm:"column:email"`
	Password       string                  `gorm:"column:password"`
	PhoneNumber    string                  `gorm:"column:phone_number"`
	ProfilePicture string                  `gorm:"column:profile_picture"`
	IsSeller       bool                    `gorm:"column:is_seller"`
	IsAdmin        bool                    `gorm:"column:is_admin"`
	CreatedAt      time.Time               `gorm:"column:created_at"`
	UpdatedAt      time.Time               `gorm:"column:updated_at"`
	Products       []product.ProductEntity `gorm:"foreignKey:SellerID"`
}
