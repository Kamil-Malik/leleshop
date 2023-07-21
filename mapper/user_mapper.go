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
		FcmToken:       dto.FcmToken,
		PhoneNumber:    dto.PhoneNumber,
		ProfilePicture: dto.ProfilePicture,
		Email:          dto.Email,
		Password:       dto.Password,
		IsSeller:       dto.IsSeller,
		IsAdmin:        dto.IsAdmin,
	}
}

func ToUserDto(entity entity.UserEntity) dto.UserDto {
	return dto.UserDto{
		Id:             entity.Id,
		UserName:       entity.UserName,
		FullName:       entity.FullName,
		FcmToken:       entity.FcmToken,
		Email:          entity.Email,
		Password:       entity.Password,
		PhoneNumber:    entity.PhoneNumber,
		ProfilePicture: entity.ProfilePicture,
		IsSeller:       entity.IsSeller,
		IsAdmin:        entity.IsAdmin,
	}
}
