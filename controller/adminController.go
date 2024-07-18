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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary List all users
// @Description List all users from the database by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param email header string true "Email of the user"
// @Success 200 {object} string
// @Router /v1/ecommerce/users [get]
func ListAllUsers(c *gin.Context) {
	// var req struct {
	// 	Email string `json:"email"`
	// }
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"message": "Invalid request",
	// 	})
	// 	return
	// }

	email := c.Param("email")
	fmt.Println(email)

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking is admin or not

	isAdmin, err := helper.IsUserAdmin(c, email)
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
	var req struct {
		Email       string          `json:"email"`
		Name        string          `json:"name"`
		Price       int             `json:"price"`
		Description string          `json:"description"`
		Images      string          `json:"images"`
		Rating      float64         `json:"rating"`
		Stock       int             `json:"stock"`
		Keywords    []string        `json:"keywords"`
		NumRating   int             `json:"num_rating"`
		Commnets    []types.Comment `json:"comments"`
		CategoryId  string          `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	product, err := productCollection.InsertOne(c.Request.Context(), bson.M{"name": req.Name, "price": req.Price, "description": req.Description, "images": req.Images, "rating": req.Rating, "stock": req.Stock, "keywords": req.Keywords, "num_rating": req.NumRating, "comments": req.Commnets, "category_id": req.CategoryId, "id": primitive.NewObjectID().Hex()})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error adding product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product added",
		"product": product,
	})

}

func UpdateProduct(c *gin.Context) {
	var req struct {
		Email       string          `json:"email"`
		Name        string          `json:"name"`
		Price       int             `json:"price"`
		Description string          `json:"description"`
		Images      string          `json:"images"`
		Rating      float64         `json:"rating"`
		Stock       int             `json:"stock"`
		Keywords    []string        `json:"keywords"`
		NumRating   int             `json:"num_rating"`
		Commnets    []types.Comment `json:"comments"`
		CategoryId  string          `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err := productCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$set": bson.M{"name": req.Name, "price": req.Price, "description": req.Description, "images": req.Images, "rating": req.Rating, "stock": req.Stock, "keywords": req.Keywords, "num_rating": req.NumRating, "comments": req.Commnets, "category_id": req.CategoryId}})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error updating product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product updated",
	})
}

func DeleteProduct(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err := productCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error deleting product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product deleted",
	})

}

func ListProducts(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	results, err := productCollection.Find(c.Request.Context(), bson.M{}, nil)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error fetching products",
		})
		return
	}

	defer results.Close(c.Request.Context())

	var products []types.Product

	for results.Next(c.Request.Context()) {
		var singleProduct types.Product
		if err = results.Decode(&singleProduct); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		}

		products = append(products, singleProduct)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Products fetched",
		"products": products,
		"error":    false,
	})

}

func AddCategory(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	cat, err := categoryCollection.InsertOne(c.Request.Context(), bson.M{"category": req.Category, "id": primitive.NewObjectID().Hex()})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error adding category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Category added",
		"category": cat,
	})

}

func UpdateCategory(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	_, err := categoryCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$set": bson.M{"category": req.Category}})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error updating category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category updated",
	})

}

func DeleteCategory(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	_, err := categoryCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error deleting category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category deleted",
	})
}

func AddCoupon(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Discount int    `json:"discount"`
		Expiry   string `json:"expiry"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var coupounCollection *mongo.Collection = database.GetCollection(database.DB, constant.CouponCollection)

	coupon, err := coupounCollection.InsertOne(c.Request.Context(), bson.M{"name": req.Name, "discount": req.Discount, "expiry": req.Expiry, "id": primitive.NewObjectID().Hex()})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error adding coupon",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Coupon added",
		"coupon":  coupon,
	})

}

func DeleteCoupon(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var couponCollection *mongo.Collection = database.GetCollection(database.DB, constant.CouponCollection)

	_, err := couponCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error deleting coupon",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Coupon deleted",
	})

}

func ListCoupons(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var couponCollection *mongo.Collection = database.GetCollection(database.DB, constant.CouponCollection)

	results, err := couponCollection.Find(c.Request.Context(), bson.M{}, nil)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error fetching coupons",
		})
		return
	}

	defer results.Close(c.Request.Context())

	var coupons []types.Coupon

	for results.Next(c.Request.Context()) {
		var singleCoupon types.Coupon
		if err = results.Decode(&singleCoupon); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		}

		coupons = append(coupons, singleCoupon)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Coupons fetched",
		"coupons": coupons,
		"error":   false,
	})

}

func ListAllOrders(c *gin.Context) {
}

func AddStock(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Stock int    `json:"stock"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	isAdmin, _ := helper.IsUserAdmin(c, req.Email)

	if !isAdmin {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err := productCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$inc": bson.M{"stock": req.Stock}})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error updating stock",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Stock updated",
	})

}
