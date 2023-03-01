package jwttoken

import (
	"Pet_1/internal/domain/entity"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

//var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(c entity.CustomerLoginViewModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["user"] = c.PhoneNumber
	claims["password"] = c.Password

	tokenString, err := token.SignedString([]byte(viper.GetString("Jwt.Secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(r *http.Request) (*jwt.Token, error) {
	//t := r.Header["Token"][0]
	t := r.Header.Get("Authorization")
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Wrong signature method")
		}
		return []byte(viper.GetString("Jwt.Secret")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Unauthorized due to error parsing JWT", err)
	}
	return token, nil
}

func ValidateJWT(r *http.Request) error {
	token, err := VerifyJWT(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return fmt.Errorf("Unauthorized due to error validation JWT", err)
	}
	return nil
}
