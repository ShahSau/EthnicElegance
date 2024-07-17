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
