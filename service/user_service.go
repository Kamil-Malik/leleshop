package service

import (
	"errors"
	"leleshop/db"
	dto "leleshop/dto/user"
	entity "leleshop/entity/user"
	"leleshop/mapper"

	"gorm.io/gorm"
)

func SignUpUser(newUser entity.UserEntity) error {
	db := db.GetDB()

	// Check if the username is already taken
	var existingUser entity.UserEntity
	if err := db.Where("user_name = ?", newUser.UserName).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// The username is not taken, proceed to create the new user record
			return db.Create(&newUser).Error
		}
		// An error occurred while querying the database, return it
		return err
	}

	// The username is already taken, return an error indicating that the username is not available
	return errors.New("username is already taken")
}

func GetUserByUsername(userName string) (dto dto.UserDto, err error) {
	db := db.GetDB()
	var user entity.UserEntity
	err = db.Where("user_name = ?", userName).First(&user).Error
	dto = mapper.ToUserDto(user)
	return dto, err
}
