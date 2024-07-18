package helper

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ShahSau/EthnicElegance/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserValidation(u types.UserClient) error {
	fmt.Println("CheckUserValidation:", u)
	if u.Email == "" {
		return errors.New("email can't be empty")
	}
	if u.Name == "" {
		return errors.New("name can't be empty")
	}
	if u.Phone == "" {
		return errors.New("phone can't be empty")
	}
	if u.Password == "" {
		return errors.New("password can't be empty")
	}
	return nil
}

func EncryptPassword(s string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ComparePassword(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

func IsUserAdmin(c *gin.Context, tokenString string) (bool, error) {

	return true, nil

}

// func GenerateID() string {
// 	return bson.NewObjectID().Hex()
// }

func GenerateToken(userId string, email string, userType string) (string, error) {
	secretKey := os.Getenv("secretKey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"type":   userType,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (string, string, error) {
	secretKey := os.Getenv("secretKey")
	token, err := jwt.Parse((tokenString), func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Print("error in isuseradmin", err)
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("could not parse claims")
	}

	email := claims["type"].(string)
	userType := claims["type"].(string)

	return email, userType, nil
}
