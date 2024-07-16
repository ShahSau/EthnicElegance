package controller

import "github.com/gin-gonic/gin"

func RegisterUser(router *gin.Context) {
	router.JSON(200, gin.H{
		"message": "Register User",
	})
}

func VerifyEmail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Verify Email",
	})
}

func VerifyOtp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Verify Otp",
	})
}

//Resend Email should call verify email function

func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Login",
	})
}

func Signout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Signout",
	})
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
