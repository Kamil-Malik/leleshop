package mapper

import (
	dto "leleshop/dto/user"
	entity "leleshop/entity/user"
)

func ToUserEntity(dto dto.UserDto) entity.UserEntity {
	return entity.UserEntity{
		Id:             dto.Id,
		UserName:       dto.UserName,
		FullName:       dto.FullName,
		PhoneNumber:    dto.PhoneNumber,
		ProfilePicture: dto.ProfilePicture,
		Email:          dto.Email,
		Password:       dto.Password,
		IsSeller:       dto.IsSeller,
		IsAdmin:        dto.IsAdmin,
	}
}
