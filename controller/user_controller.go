package controller

import (
	"leleshop/dto/response"
	"leleshop/dto/user"
	"leleshop/mapper"
	"leleshop/service"
	"leleshop/util"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user user.UserDto
	contentType := ctx.Request.Header.Get(util.ContentType)

	if contentType != util.ApplicationJson {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "Only JSON is being accepted for this endpoint, please refer to our docs",
			},
		)
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "Invalid JSON format",
			},
		)
		return
	}

	userEntity := mapper.ToUserEntity(user)
	userEntity.Password = util.HashPass(userEntity.Password)
	if err := service.SignUpUser(userEntity); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		response.Response{
			Status:  true,
			Message: "Ok",
			Data:    "Account Created",
		},
	)
}

func Login(ctx *gin.Context) {
	var userNameLoginDto user.UserNameLoginDto
	contentType := ctx.Request.Header.Get(util.ContentType)

	// Get JSON or Body Form from HTTP Request
	if contentType == util.ApplicationJson {
		if err := ctx.ShouldBindJSON(&userNameLoginDto); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				response.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&userNameLoginDto); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				response.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return
		}
	}

	// Validate Struct with validator
	if _, err := valid.ValidateStruct(&userNameLoginDto); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
		return
	}

	user, err := service.GetUserByUsername(userNameLoginDto.UserName)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			response.ErrorResponse{
				Status:  false,
				Message: "Username not found",
			},
		)
		return
	}

	if err := util.ComparePassword(userNameLoginDto.Password, user.Password); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusForbidden,
			response.ErrorResponse{
				Status:  false,
				Message: "Invalid username / password",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.Response{
			Status:  true,
			Message: "Ok",
			Data:    util.GenerateToken(user.Id, user.UserName),
		},
	)
}

func GetUserProfile(ctx *gin.Context) {
	userName := ctx.Param("user_name")
	user, err := service.GetUserByUsername(userName)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			response.ErrorResponse{
				Status:  false,
				Message: "Username not found",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.Response{
			Status:  true,
			Message: "Ok",
			Data:    mapper.ToUserResponseDto(user),
		},
	)
}
