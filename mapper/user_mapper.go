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

func ToUserResponseDto(user dto.UserDto) dto.UserResponseDto {
	return dto.UserResponseDto{
		Id:             user.Id,
		UserName:       user.UserName,
		FullName:       user.FullName,
		FcmToken:       user.FcmToken,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		ProfilePicture: user.ProfilePicture,
		IsSeller:       user.IsSeller,
		IsAdmin:        user.IsAdmin,
	}
}
