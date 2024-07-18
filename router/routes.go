package router

import (
	"net/http"

	"github.com/ShahSau/EthnicElegance/constant"
	"github.com/ShahSau/EthnicElegance/controller"
)

// health check service
var healthCheckRoutes = Routes{
	Route{"Health check", http.MethodGet, constant.HealthCheckRoute, controller.HealthCheck},
}

var userRoutes = Routes{
	// Route{"VerifyEmail", http.MethodPost, constant.VerifyEmailRoute, controller.VerifyEmail},
	// Route{"VerifyOtp", http.MethodPost, constant.VerifyOtpRoute, controller.VerifyOtp},
	// Route{"ResendEmail", http.MethodPost, constant.ResendEmailRoute, controller.VerifyEmail},

	// Register User
	Route{"Register User", http.MethodPost, constant.UserRegisterRoute, controller.RegisterUser},
	Route{"Login User", http.MethodPost, constant.UserLoginRoute, controller.UserLogin},

	Route{"Sign Out", http.MethodPost, constant.UserLogoutRoute, controller.SignOut},
}

var productGlobalRoutes = Routes{
	Route{"List Product", http.MethodGet, constant.ListProductRoute, controller.ListProductsController},
	Route{"Search Product", http.MethodGet, constant.SearchProductRoute, controller.SearchProductController},
	Route{"List Category", http.MethodGet, constant.ListCategoryRoute, controller.ListCategoryController},
	Route{"List Single Product", http.MethodGet, constant.ListSingleProductRoute, controller.ListSingleProductController},
	Route{"Get Product Link", http.MethodGet, constant.GetProductLinkRoute, controller.GetProductLink},
}

var userAuthRoutes = Routes{
	Route{"Add to cart", http.MethodPost, constant.AddToCartRoute, controller.AddToCart},
	Route{"Add Address", http.MethodPost, constant.AddAddressRoute, controller.AddAddress},
	Route{"Edit Address", http.MethodPut, constant.EditAddressRoute, controller.EditAddress},
	Route{"Edit Name", http.MethodPut, constant.EditNameRoute, controller.EditName},
	Route{"Update User", http.MethodPut, constant.UpdateUser, controller.UpdateUser},
	Route{"Checkout Order", http.MethodPut, constant.CheckoutRoute, controller.CheckoutOrder},
	Route{"Add to Favorite", http.MethodPost, constant.AddToFavoriteRoute, controller.AddToFavorite},
	Route{"Remove from Favorite", http.MethodDelete, constant.RemoveFromFavoriteRoute, controller.RemoveFromFavorite},
	Route{"List Favorite", http.MethodGet, constant.ListFavoriteRoute, controller.ListFavorite},
	Route{"Remove from Cart", http.MethodDelete, constant.RemoveFromCartRoute, controller.RemoveFromCart},
	Route{"List Cart", http.MethodGet, constant.ListCartRoute, controller.ListCart},
	Route{"Empty Cart", http.MethodDelete, constant.EmptyCartRoute, controller.EmptyCart},
	Route{"Apply Coupon", http.MethodPost, constant.ApplyCouponRoute, controller.ApplyCoupon},
}

var adminRoutes = Routes{
	Route{"All Users", http.MethodGet, constant.GetAllUserRoute, controller.ListAllUsers},
	Route{"Block User", http.MethodPut, constant.BlockUserRoute, controller.BlockUser},
	Route{"Unblock User", http.MethodPut, constant.UnblockUserRoute, controller.UnblockUser},
	Route{"Register Product", http.MethodPost, constant.RegisterProductRoute, controller.RegisterProduct},
	Route{"Update Product", http.MethodPut, constant.UpdateProductRoute, controller.UpdateProduct},
	Route{"Delete Product", http.MethodDelete, constant.DeleteProductRoute, controller.DeleteProduct},
	Route{"Add Category", http.MethodPost, constant.AddCategoryRoute, controller.AddCategory},
	Route{"Update Category", http.MethodPut, constant.UpdateCategoryRoute, controller.UpdateCategory},
	Route{"Delete Category", http.MethodDelete, constant.DeleteCategoryRoute, controller.DeleteCategory},
	Route{"Add coupons", http.MethodPost, constant.AddCouponRoute, controller.AddCoupon},
	Route{"Delete coupons", http.MethodDelete, constant.DeleteCouponRoute, controller.DeleteCoupon},
	Route{"List Coupons", http.MethodGet, constant.ListCouponRoute, controller.ListCoupons},
	Route{"List all products", http.MethodGet, constant.ListProductRoute, controller.ListProducts},
	Route{"List all orders", http.MethodGet, constant.ListOrders, controller.ListAllOrders},
	Route{"Add Stock", http.MethodPut, constant.UpdateStockRoute, controller.AddStock},
}
