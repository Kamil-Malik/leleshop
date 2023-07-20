package service

import (
	"leleshop/db"
	"leleshop/entity/user"
)

func SignUpUser(entity user.UserEntity) error {
	db := db.GetDB()
	return db.Create(&entity).Error
}
