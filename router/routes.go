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
	// Route{"VerifyEmail", http.MethodPost, constant.VerifyEmailRoute, controller.},
	// Route{"VerifyOtp", http.MethodPost, constant.VerifyOtpRoute, controller.},
	// Route{"ResendEmail", http.MethodPost, constant.ResendEmailRoute, controller.},

	// Register User
	// Route{"Register User", http.MethodPost, constant.UserRegisterRoute, controller.},
	// Route{"Login User", http.MethodPost, constant.UserLoginRoute, controller.},

	// Route{"Signout", http.MethodGet, constant.UserLogoutRoute, controller.},
}

var productGlobalRoutes = Routes{
	// Route{"List Product", http.MethodGet, constant.ListProductRoute, controller.},
	// Route{"Search Product", http.MethodGet, constant.SearchProductRoute, controller.},
	// Route{"List Category", http.MethodGet, constant.ListCategoryRoute, controller.},
	// Route{"List Single Product", http.MethodGet, constant.ListSingleProductRoute, controller.},
}

var userAuthRoutes = Routes{
	// Route{"Add to cart", http.MethodPost, constant.AddToCartRoute, controller.},
	// Route{"Add Address", http.MethodPost, constant.AddAddressRoute, controller.},
	// Route{"Edit Address", http.MethodPut, constant.EditAddressRoute, controller.},
	// Route{"Update User", http.MethodPut, constant.UpdateUser, controller.},
	// Route{"Checkout Order", http.MethodPut, constant.CheckoutRoute, controller.},
	// Route{"Add to Favorite", http.MethodPost, constant.AddToFavoriteRoute, controller.},
	// Route{"Remove from Favorite", http.MethodDelete, constant.RemoveFromFavoriteRoute, controller.},
	// Route{"List Favorite", http.MethodGet, constant.ListFavoriteRoute, controller.},
}

var adminRoutes = Routes{
	// Route{"All Users", http.MethodGet, constant.GetAllUserRoute, controller.},
	// Route{"Block User", http.MethodPut, constant.BlockUserRoute, controller.},
	// Route{"Unblock User", http.MethodPut, constant.UnblockUserRoute, controller.},
	// Route{"Register Product", http.MethodPost, constant.RegisterProductRoute, controller.},
	// Route{"Update Product", http.MethodPut, constant.UpdateProductRoute, controller.},
	// Route{"Delete Product", http.MethodDelete, constant.DeleteProductRoute, controller.},
	// Route{"Add Category", http.MethodPost, constant.AddCategoryRoute, controller.},
	// Route{"Update Category", http.MethodPut, constant.UpdateCategoryRoute, controller.},
	// Route{"Delete Category", http.MethodDelete, constant.DeleteCategoryRoute, controller.},
	// Route{"Add coupons", http.MethodPost, constant.AddCouponRoute, controller.},
	// Route{"Delete coupons", http.MethodDelete, constant.DeleteCouponRoute, controller.},
}
