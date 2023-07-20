package util

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) string {
	salt := 8
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(plainPassword, hashedPassword string) error {
	login, local := []byte(plainPassword), []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(local, login)
	return err
}
