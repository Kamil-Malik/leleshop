package util

import (
	"errors"
	"leleshop/dto/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "leleshop"

func GenerateToken(id uint, username string) string {
	claims := jwt.MapClaims{
		"id":        id,
		"user_name": username,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("please signin to continue")
	headerToken := ctx.Request.Header.Get("Authorization")

	if len(headerToken) == 0 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  false,
				Message: "Authorization header cannot be empty",
			},
		)
		return nil, nil
	}

	hasBearer := strings.HasPrefix(headerToken, "Bearer")

	if !hasBearer {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  false,
				Message: "Authorization should start with prefix Bearer",
			},
		)
		return nil, nil
	}

	removedPrefix, _ := strings.CutPrefix(headerToken, "Bearer")
	if len(removedPrefix) < 2 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  false,
				Message: "JWT cannot be empty",
			},
		)
		return nil, nil
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.ErrorResponse{
					Status:  false,
					Message: errResponse.Error(),
				},
			)
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  false,
				Message: "JWT doesn't meet server requirement",
			},
		)
		return nil, nil
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  false,
				Message: "Invalid JWT",
			},
		)
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
