package controller

import (
	"fmt"
	"net/http"
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

// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// // @Param			signup  body  models.UserDetails  true	"signup"
// // @Success		200	{object}	response.Response{}
// // @Failure		500	{object}	response.Response{}
// @Router			/users/signup [post]
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
		Address:   "",
		Favourite: []string{},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Id:        primitive.NewObjectID(),
	}

	_, insertErr := userCollection.InsertOne(c, dbUser)

	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": insertErr.Error()})
		return
	}

	// jwt token
	token, err := helper.GenerateToken(dbUser.Id.Hex(), dbUser.Email, dbUser.UserType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
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
	token, err := helper.GenerateToken(dbUser.Id.Hex(), dbUser.Email, dbUser.UserType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	dbUser.Password = ""

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": dbUser, "token": token})

}

func SignOut(c *gin.Context) {
	use := c.Value("user")
	fmt.Println("SignOut", use)
	// c.Set("user", "")
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
}

func AddAddress(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	// checking is admin or not
	_, _, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	type AddressData struct {
		Address string `json:"address" bson:"address"`
		Email   string `json:"email" bson:"email"`
	}

	var addAddress AddressData

	defer c.Request.Body.Close()

	// binding the request body to address
	reqErr := c.ShouldBindJSON(&addAddress)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": addAddress.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	// updating the address
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": addAddress.Email}, bson.M{"$set": bson.M{"address": addAddress.Address}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})

}

func EditAddress(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	type AddressData struct {
		Address string `json:"address" bson:"address"`
		Email   string `json:"email" bson:"email"`
	}

	var editAddress AddressData

	defer c.Request.Body.Close()

	// binding the request body to address
	reqErr := c.ShouldBindJSON(&editAddress)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": editAddress.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	// updating the address
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": editAddress.Email}, bson.M{"$set": bson.M{"address": editAddress.Address}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
}

func UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	type UpdatePassword struct {
		Email       string `json:"email" bson:"email"`
		OldPassword string `json:"oldPassword" bson:"oldPassword"`
		NewPassword string `json:"newPassword" bson:"newPassword"`
	}

	var updatePassword UpdatePassword

	defer c.Request.Body.Close()

	// binding the request body to updatePassword
	reqErr := c.ShouldBindJSON(&updatePassword)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": updatePassword.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	// checking the password
	if !helper.ComparePassword(dbUser.Password, updatePassword.OldPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "password not matched"})
		return
	}

	// updating the password
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": updatePassword.Email}, bson.M{"$set": bson.M{"password": helper.EncryptPassword(updatePassword.NewPassword)}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
}

func EditName(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	type NameData struct {
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	}

	var editName NameData

	defer c.Request.Body.Close()

	// binding the request body to address
	reqErr := c.ShouldBindJSON(&editName)

	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": reqErr.Error()})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": editName.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	// updating the name
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": editName.Email}, bson.M{"$set": bson.M{"name": editName.Name}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
}

func AddToFavorite(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	var req struct {
		Email     string `json:"email" bson:"email"`
		ProductId string `json:"productId" bson:"productId"`
	}

	defer c.Request.Body.Close()

	// binding the request body to address
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": req.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}
	// updating the favourite
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": req.Email}, bson.M{"$push": bson.M{"favourite": req.ProductId}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Added to Favorite"})
}

func RemoveFromFavorite(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	var req struct {
		Email     string `json:"email" bson:"email"`
		ProductId string `json:"productId" bson:"productId"`
	}

	defer c.Request.Body.Close()

	// binding the request body to address
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": req.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}
	// updating the favourite
	_, updateErr := userCollection.UpdateOne(c, bson.M{"email": req.Email}, bson.M{"$pull": bson.M{"favourite": req.ProductId}})

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Removed from Favorite"})
}

func ListFavorite(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	var req struct {
		Email string `json:"email" bson:"email"`
	}

	defer c.Request.Body.Close()

	// binding the request body to address
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	var dbUser types.User

	// checking if user exists
	emailExists := userCollection.FindOne(c, bson.M{"email": req.Email}).Decode(&dbUser)

	if emailExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "email not found"})
		return
	}

	var allFav []types.Product
	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)
	for _, v := range dbUser.Favourite {
		var singleProduct types.Product
		productCollection.FindOne(c, bson.M{"id": v}).Decode(&singleProduct)
		allFav = append(allFav, singleProduct)
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": allFav})
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
