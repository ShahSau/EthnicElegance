package controller

import "github.com/gin-gonic/gin"

func ListAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List All Users",
	})
}

func BlockUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Block User",
	})
}

func UnblockUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Unblock User",
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
