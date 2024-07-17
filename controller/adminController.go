package controller

import (
	"fmt"
	"net/http"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/helper"
	"github.com/ShahSau/EthnicElegance/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListAllUsers(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking is admin or not

	isAdmin, err := helper.IsUserAdmin(c, req.Email)
	fmt.Println(isAdmin, err)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	results, err := userCollection.Find(c.Request.Context(), bson.M{}, nil)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error fetching users",
		})
		return
	}

	defer results.Close(c.Request.Context())

	var users []types.User

	for results.Next(c.Request.Context()) {
		var singleUser types.User
		if err = results.Decode(&singleUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		}

		users = append(users, singleUser)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched",
		"users":   users,
		"error":   false,
	})

}

func BlockUser(c *gin.Context) {
	var req struct {
		Email     string `json:"email"`
		UserEmail string `json:"user_email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking is admin or not

	isAdmin, err := helper.IsUserAdmin(c, req.Email)
	fmt.Println(isAdmin, err)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	_, err = userCollection.UpdateOne(c.Request.Context(), bson.M{"email": req.UserEmail}, bson.M{"$set": bson.M{"is_blocked": true}})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error blocking user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User blocked",
	})

}

func UnblockUser(c *gin.Context) {
	var req struct {
		Email     string `json:"email"`
		UserEmail string `json:"user_email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking is admin or not

	isAdmin, err := helper.IsUserAdmin(c, req.Email)
	fmt.Println(isAdmin, err)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}
	var dbUser types.User
	err = userCollection.FindOne(c.Request.Context(), bson.M{"email": req.UserEmail}).Decode(&dbUser)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error checking user block status",
		})
		return
	}
	if !dbUser.IsBlocked {
		c.JSON(400, gin.H{
			"message": "User is not blocked",
		})
		return
	}
	_, err = userCollection.UpdateOne(c.Request.Context(), bson.M{"email": req.UserEmail}, bson.M{"$set": bson.M{"is_blocked": false}})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error unblocking user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User unblocked",
	})
}

func RegisterProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register Product",
	})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Product",
	})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Product",
	})
}

func AddCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add Category",
	})
}

func UpdateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Category",
	})
}

func DeleteCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Category",
	})
}

func AddCoupon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add Coupon",
	})
}

func DeleteCoupon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Coupon",
	})
}
