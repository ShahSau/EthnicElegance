package constant

const (
	APIVersion = "v1"

	BadRequestMessage = "request not fulfilled"

	//schedular constants
	HealthCheckRoute = "/health"
	Database         = "EthnicElegance"

	// email verification routes
	VerifyEmailRoute = "/verify-email"
	VerifyOtpRoute   = "/verify-otp"
	ResendEmailRoute = "/resend-email"

	// user related routes
	UserRegisterRoute = "/user-register"
	UserLoginRoute    = "/login"
	UserLogoutRoute   = "/logout"

	// product and adminroutes
	RegisterProductRoute    = "/product-register"
	ListProductRoute        = "/list-products"
	ListCategoryRoute       = "/list-category"
	ListSingleProductRoute  = "/product/:id"
	SearchProductRoute      = "/search"
	UpdateProductRoute      = "/update-product"
	DeleteProductRoute      = "/delete-product"
	AddToCartRoute          = "/cart"
	AddAddressRoute         = "/address"
	EditAddressRoute        = "/address"
	GetSingleUserRoute      = "/user/:id"
	UpdateUser              = "/update-user"
	CheckoutRoute           = "/user/:id"
	AddToFavoriteRoute      = "/favorite"
	RemoveFromFavoriteRoute = "/favorite"
	ListFavoriteRoute       = "/favorite"
	GetAllUserRoute         = "/users"
	BlockUserRoute          = "/block-user"
	UnblockUserRoute        = "/unblock-user"
	AddCategoryRoute        = "/category"
	UpdateCategoryRoute     = "/category"
	DeleteCategoryRoute     = "/category"
	AddCouponRoute          = "/coupon"
	DeleteCouponRoute       = "/coupon"
	RemoveFromCartRoute     = "/cart"
	ListCartRoute           = "/cart"
	EmptyCartRoute          = "/cart"
	ApplyCouponRoute        = "/cart"
	GetProductLinkRoute     = "/product-link"
)

const (
	NormalUser = "user"
	AdminUser  = "admin"
)

const (
	// time slot for otp validation
	OtpValidation = 60
)

// collections
const (
	VerificationsCollection = "verifications"
	UserCollection          = "user"
	ProductCollection       = "products"
	AddressCollection       = "user_addresses"
	CartCollection          = "user_cart"
)

// messages
const (
	AlreadyRegisterWithThisEmail = "already register with this email"
	EmailIsNotVerified           = "your email is not verified please verify your email"
	EmailValidationError         = "wrong email entered"
	OtpValidationError           = "wrong otp entered"
	OtpExpiredValidationError    = "otp expired"
	AlreadyVerifiedError         = "already verified"
	OptAlreadySentError          = "otp already sent to email"
	NotRegisteredUser            = "you are not a register user"
	PasswordNotMatchedError      = "email and password do not match"
	NotAuthorizedUserError       = "you are not authorized"
	NoProductAvaliable           = "no product avaliable"
	UserDoesNotExists            = "user not exists"
	AddressNotExists             = "address not exists. please add one address"
)
