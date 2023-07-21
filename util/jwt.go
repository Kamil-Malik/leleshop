package util

import "github.com/golang-jwt/jwt/v5"

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
