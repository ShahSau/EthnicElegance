package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/helper"
	"github.com/ShahSau/EthnicElegance/types"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JwtClaim struct {
	UserId   primitive.ObjectID
	Email    string
	UserType string
	jwt.StandardClaims
}

func RegisterUser(c *gin.Context) {
	var userClient types.UserClient
	var dbUser types.User

	defer c.Request.Body.Close()
	// binding the request body to userClient
	reqErr := c.ShouldBindJSON(&userClient)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	// checking the payload
	err := helper.CheckUserValidation(userClient)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking if email is unique
	emailExists := userCollection.FindOne(c, bson.M{"email": userClient.Email}).Decode(&dbUser)

	fmt.Println("emaiailExists:", emailExists)

	if emailExists == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email already exists"})
		return
	}

	// creating the user object
	dbUser = types.User{
		Name:      userClient.Name,
		Email:     userClient.Email,
		Phone:     userClient.Phone,
		Password:  helper.EncryptPassword(userClient.Password),
		UserType:  "user",
		IsBlocked: false,
		Favourite: []string{},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Id:        primitive.NewObjectID(),
	}

	result, insertErr := userCollection.InsertOne(c, dbUser)

	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": insertErr.Error()})
		return
	}

	// jwt token

	claims := &JwtClaim{
		UserId:   result.InsertedID.(primitive.ObjectID),
		UserType: dbUser.UserType,
		Email:    dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: &jwt.Time{time.Now().Add(time.Hour * time.Duration(48))},
			Issuer:    os.Getenv("jwtIssuer"),
		},
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := token1.SignedString([]byte(os.Getenv("jwtSecret")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	dbUser.Password = ""

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": dbUser, "token": token})

}

func UserLogin(c *gin.Context) {
	var loginReq types.Login

	defer c.Request.Body.Close()

	// binding the request body to loginReq
	reqErr := c.ShouldBindJSON(&loginReq)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if email exists
	emailExists := userCollection.FindOne(c, bson.M{"email": loginReq.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	// checking the password
	if !helper.ComparePassword(dbUser.Password, loginReq.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "password not matched"})
		return
	}

	// jwt token

	claims := &JwtClaim{
		UserId:   dbUser.Id,
		UserType: dbUser.UserType,
		Email:    dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: &jwt.Time{time.Now().Add(time.Hour * time.Duration(48))},
			Issuer:    os.Getenv("jwtIssuer"),
		},
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := token1.SignedString([]byte(os.Getenv("jwtSecret")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	dbUser.Password = ""

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": dbUser, "token": token})

}

func Signout(c *gin.Context) {
	var user types.User

	claims := c.MustGet("claims").(*JwtClaim)

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	userCollection.FindOne(c, bson.M{"_id": claims.UserId}).Decode(&user)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": user})
}

func AddAddress(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add Address",
	})
}

func EditAddress(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Edit Address",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update User",
	})
}

func AddToFavorite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add to Favorite",
	})
}

func RemoveFromFavorite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Remove from Favorite",
	})
}

func ListFavorite(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Favorite",
	})
}

func AddToCart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add to Cart",
	})
}

func CheckoutOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Checkout Order",
	})
}

func RemoveFromCart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Remove from Cart",
	})
}

func ListCart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Cart",
	})
}

func EmptyCart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Empty Cart",
	})
}

func ApplyCoupon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Apply Coupoun",
	})
}
