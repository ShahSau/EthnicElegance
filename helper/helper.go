package helper

import (
	"errors"
	"fmt"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func IsUserAdmin(c *gin.Context, email string) (bool, error) {
	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)
	var dbUser types.User

	err := userCollection.FindOne(c, bson.M{"email": email}).Decode(&dbUser)

	if err != nil {
		return false, err
	}
	if dbUser.UserType != "admin" {
		return false, nil
	}
	return true, nil

}
