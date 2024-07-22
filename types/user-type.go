package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Phone     string             `json:"phone" bson:"phone"`
	Password  string             `json:"password" bson:"password"`
	UserType  string             `json:"user_type" bson:"user_type"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
	Favourite []string           `json:"favourite" bson:"favourite"`
	IsBlocked bool               `json:"is_blocked" bson:"is_blocked"`
	Address   string             `json:"address" bson:"address"`
}

type UserClient struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

type Verification struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Otp       int64              `json:"otp" bson:"otp"`
	Status    bool               `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Coupon struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Discount int                `json:"discount"`
	Expiry   string             `json:"expiry"`
}

type Product struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Description string             `json:"description"`
	Images      string             `json:"images"`
	Rating      float64            `json:"rating"`
	Stock       int                `json:"stock"`
	Keywords    []string           `json:"keywords"`
	NumRating   int                `json:"num_rating"`
	Commnets    []Comment          `json:"comments"`
	CategoryId  string             `json:"category_id"`
}

type Comment struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Comment  string `json:"comment"`
}

type Category struct {
	ID       primitive.ObjectID `json:"id"`
	Category string             `json:"category"`
}

type Offer struct {
	CategoryId int  `json:"category_id"`
	Discount   int  `json:"discount"`
	Expiry     bool `json:"expiry"`
}

type AddressData struct {
	Address string `json:"address" bson:"address"`
	Email   string `json:"email" bson:"email"`
}

type UpdatePassword struct {
	Email       string `json:"email" bson:"email"`
	OldPassword string `json:"oldPassword" bson:"oldPassword"`
	NewPassword string `json:"newPassword" bson:"newPassword"`
}

type NameData struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
