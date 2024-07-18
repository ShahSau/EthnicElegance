package controller

import (
	"context"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func SearchProductController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search Product",
	})
}

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

func GetProductLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Product Link",
	})
}
