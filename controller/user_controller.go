package controller

import (
	"leleshop/dto/response"
	"leleshop/dto/user"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var userNameLoginDto user.UserNameLoginDto
	contentType := ctx.ContentType()

	// Get JSON or Body Form from HTTP Request
	if contentType == "applicatoin/json" {
		if err := ctx.ShouldBindJSON(&userNameLoginDto); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				response.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
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
		}
	}

	// Validate Struct with validator
	if _, err := govalidator.ValidateStruct(userNameLoginDto); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
	}
}
