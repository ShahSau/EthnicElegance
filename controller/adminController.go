package controller

import (
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
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/users [get]
func ListAllUsers(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	var userCollection *mongo.Collection = database.GetCollection(database.DB, constant.UsersCollection)

	// checking is admin or not
	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Block user
// @Description block user by the admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/block-user [put]
func BlockUser(c *gin.Context) {
	var req struct {
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
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary unblock user
// @Description unblock user by the admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/unblock-user [put]
func UnblockUser(c *gin.Context) {
	var req struct {
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
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Add Product
// @Description Add product by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/product-register [post]
func RegisterProduct(c *gin.Context) {
	var req struct {
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
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Update Product
// @Description Update product by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/update-product/:id [put]
func UpdateProduct(c *gin.Context) {
	var req struct {
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
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err = productCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$set": bson.M{"name": req.Name, "price": req.Price, "description": req.Description, "images": req.Images, "rating": req.Rating, "stock": req.Stock, "keywords": req.Keywords, "num_rating": req.NumRating, "comments": req.Commnets, "category_id": req.CategoryId}})

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

// @Summary Delete Product
// @Description Delete product by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/delete-product/:id [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// checking is admin or not

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err = productCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

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

// @Summary List all products
// @Description List all products from the database by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/list-products-admin [get]
func ListProducts(c *gin.Context) {
	// checking is admin or not

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Add Category
// @Description Add category by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/category [post]
func AddCategory(c *gin.Context) {
	var req struct {
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary List all categories
// @Description List all categories from the database by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/category/:id [put]
func UpdateCategory(c *gin.Context) {
	var req struct {
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
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}
	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	_, err = categoryCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$set": bson.M{"category": req.Category}})

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

// @Summary Delete Category
// @Description Delete category by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/category/:id [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	// checking is admin or not

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	_, err = categoryCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

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

// @Summary Add Coupon
// @Description Add coupon by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/coupon [post]
func AddCoupon(c *gin.Context) {
	var req struct {
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

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Delete Coupon
// @Description Delete coupon by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/coupon/:id [delete]
func DeleteCoupon(c *gin.Context) {
	id := c.Param("id")

	// checking is admin or not

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var couponCollection *mongo.Collection = database.GetCollection(database.DB, constant.CouponCollection)

	_, err = couponCollection.DeleteOne(c.Request.Context(), bson.M{"id": id})

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

// @Summary List all coupons
// @Description List all coupons from the database by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/coupon [get]
func ListCoupons(c *gin.Context) {

	// checking is admin or not
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
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

// @Summary Add Stock
// @Description Add stock by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/update-stock/:id [put]
func AddStock(c *gin.Context) {
	var req struct {
		Stock int `json:"stock"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	id := c.Param("id")

	// checking is admin or not

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	_, err = productCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$inc": bson.M{"stock": req.Stock}})

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

// @Summary Add Offer
// @Description Add offer by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/offer [post]
func AddOffer(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	var req struct {
		CategoryId int  `json:"category_id"`
		Discount   int  `json:"discount"`
		Expiry     bool `json:"expiry"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// checking is admin or not
	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var offerCollection *mongo.Collection = database.GetCollection(database.DB, constant.OfferCollection)

	offer, err := offerCollection.InsertOne(c.Request.Context(), bson.M{"category_id": req.CategoryId, "discount": req.Discount, "expiry": req.Expiry, "id": primitive.NewObjectID().Hex()})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error adding offer",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Offer added",
		"offer":   offer,
	})
}

// @Summary List all offers
// @Description List all offers from the database by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/offer [get]
func ListAllOffers(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	// checking is admin or not
	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var offerCollection *mongo.Collection = database.GetCollection(database.DB, constant.OfferCollection)

	results, err := offerCollection.Find(c.Request.Context(), bson.M{}, nil)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error fetching offers",
		})
		return
	}

	defer results.Close(c.Request.Context())

	var offers []types.Offer

	for results.Next(c.Request.Context()) {
		var singleOffer types.Offer
		if err = results.Decode(&singleOffer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		}

		offers = append(offers, singleOffer)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Offers fetched",
		"offers":  offers,
		"error":   false,
	})
}

// @Summary Change Offer Status
// @Description Change expiry of the offers by admin
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/ecommerce/offer/:id [put]
func ChangeOffersStatus(c *gin.Context) {
	id := c.Param("id")
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}

	// checking is admin or not
	_, userType, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType != "admin" {
		c.JSON(400, gin.H{
			"message": "User is not an admin",
		})
		return
	}

	var offerCollection *mongo.Collection = database.GetCollection(database.DB, constant.OfferCollection)

	var dbOffer types.Offer
	_, err = offerCollection.UpdateOne(c.Request.Context(), bson.M{"id": id}, bson.M{"$set": bson.M{"expiry": !dbOffer.Expiry}})

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error changing offer status",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Offer status changed",
	})

}
