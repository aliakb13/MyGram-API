package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ParamAndUserId struct {
	Param  int
	UserId string
}

func HashPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return "", err
	}

	return string(encryptPassword), nil

}

func ComparedPassword(hashPass string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

	return err == nil
}

func GenerateToken(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})

	return token.SignedString([]byte("warteg bahari enak"))
}

func VerifyToken(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid method")
		}

		return []byte("warteg bahari enak"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetIdFromClaims(mapClaims any) (string, error) {
	claims, ok := mapClaims.(map[string]any)

	if !ok {
		return "", errors.New("invalid claims")
	}

	id, ok := claims["sub"]

	if !ok {
		return "", errors.New("sub not found")
	}

	return id.(string), nil
}
