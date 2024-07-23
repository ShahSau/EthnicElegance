package controller

import (
	"context"
	"os"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/helper"
	"github.com/ShahSau/EthnicElegance/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary List all products
// @Description List all products
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}  string
// @Router /v1/ecommerce/products [get]
func ListProductsController(c *gin.Context) {

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	var products []types.Product

	cursor, err := productCollection.Find(context.Background(), bson.D{})

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product types.Product
		cursor.Decode(&product)
		if product.Stock > 0 {
			products = append(products, product)
		}
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

// @Summary Search product
// @Description Search product
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}  string
// @Router /search [get]
func SearchProductController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search Product",
	})
}

// @Summary List all categories
// @Description List all categories
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}  string
// @Router /v1/ecommerce/list-category [get]
func ListCategoryController(c *gin.Context) {
	var categoryCollection *mongo.Collection = database.GetCollection(database.DB, constant.CategoryCollection)

	var categories []types.Category

	cursor, err := categoryCollection.Find(context.Background(), bson.D{})

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var category types.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

// @Summary List single product by id
// @Description List single product by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object}  string
// @Router /v1/ecommerce/product/{id} [get]
func ListSingleProductController(c *gin.Context) {
	Id := c.Param("id")

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	var product types.Product

	err := productCollection.FindOne(context.Background(), bson.D{{"id", Id}}).Decode(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}

// @Summary Get product link
// @Description Get product link
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object}  string
// @Router /v1/ecommerce/product-link/:id [get]
func GetProductLink(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}
	// verified user
	_, _, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	Id := c.Param("id")
	frontend := os.Getenv("frontEndUrl")
	link := frontend + "/products/" + Id

	c.JSON(200, gin.H{
		"link": link,
	})

}

// @Summary Give rating
// @Description Give rating
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object}  string
// @Router /v1/ecommerce/rating/{id} [post]
func GiveRating(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}
	// verified user
	_, _, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	Id := c.Param("id")

	var rating types.Rating

	err = c.ShouldBindJSON(&rating)

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	var product types.Product

	err = productCollection.FindOne(context.Background(), bson.D{{"id", Id}}).Decode(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	product.Rating = (product.Rating*float64(product.NumRating) + rating.Rating) / float64(product.NumRating+1)
	product.NumRating = product.NumRating + 1

	_, updateErr := productCollection.UpdateOne(c, bson.M{"id": Id}, bson.M{"$set": bson.M{"rating": product.Rating, "numrating": product.NumRating}})
	if updateErr != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Rating added",
	})
}

// @Summary Comment on product
// @Description Comment on product
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object}  string
// @Router /v1/ecommerce/comment/{id} [post]
func CommentOnProduct(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"message": "Token is required",
		})
		return
	}
	// verified user
	_, _, err := helper.VerifyToken(token)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	Id := c.Param("id")

	var comment types.Comment

	err = c.ShouldBindJSON(&comment)

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	var productCollection *mongo.Collection = database.GetCollection(database.DB, constant.ProductCollection)

	var product types.Product

	err = productCollection.FindOne(context.Background(), bson.D{{"id", Id}}).Decode(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	output := append(product.Comments, comment)

	_, updateErr := productCollection.UpdateOne(c, bson.M{"id": Id}, bson.M{"$set": bson.M{"comments": output}})
	if updateErr != nil {
		c.JSON(400, gin.H{
			"message": constant.BadRequestMessage,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Comment added",
	})
}
