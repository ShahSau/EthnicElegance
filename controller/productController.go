package controller

import "github.com/gin-gonic/gin"

func ListProductsController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Products",
	})
}

func SearchProductController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search Product",
	})
}

func ListCategoryController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Category",
	})
}

func ListSingleProductController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List Single Product",
	})
}

func GetProductLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Product Link",
	})
}
