package service

import (
	"leleshop/db"
	dto "leleshop/dto/user"
	entity "leleshop/entity/user"
	"leleshop/mapper"
)

func SignUpUser(entity entity.UserEntity) error {
	db := db.GetDB()
	return db.Create(&entity).Error
}

func GetUserByUsername(userName string) (dto dto.UserDto, err error) {
	db := db.GetDB()
	var user entity.UserEntity
	err = db.Where("user_name = ?", userName).First(&user).Error
	dto = mapper.ToUserDto(user)
	return dto, err
}
