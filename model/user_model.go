package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	DeviceId     string             `bson:"device_id" json:"device_id"`
	WalletId     string             `bson:"wallet_id" json:"wallet_id"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"password"`
	Nickname     string             `bson:"nickname" json:"nickname"`
	AvatarImgURL string             `bson:"avatar_img_url" json:"avatar_img_url"`
	IsGuest      string             `bson:"is_guest" json:"is_guest"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type UserRegister struct {
	DeviceId     string `bson:"device_id" json:"device_id"`
	Email        string `bson:"email" json:"email"`
	Password     string `bson:"password" json:"password"`
	Nickname     string `bson:"nickname" json:"nickname"`
	AvatarImgURL string `bson:"avatar_img_url" json:"avatar_img_url"`
}

type UserInfo struct {
	UserId       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	DeviceId     string             `bson:"device_id" json:"device_id"`
	Email        string             `bson:"email" json:"email"`
	Nickname     string             `bson:"nickname" json:"nickname"`
	AvatarImgURL string             `bson:"avatar_img_url" json:"avatar_img_url"`
	IsGuest      string             `bson:"is_guest" json:"is_guest"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type UserWallet struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	SecretKey string             `bson:"secret_key" json:"secret_key"`
	Value     string             `bson:"value" json:"value"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type UserSession struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserInfo  UserInfo           `bson:"user_info" json:"user_info"`
	IP        string             `bson:"ip" json:"ip"`
	OS        string             `bson:"os,omitempty" json:"os"`
	Platform  string             `bson:"platform,omitempty" json:"platform"`
	ExpireAt  primitive.DateTime `bson:"expire_at" json:"expire_at"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type UserPreference struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type UserRecentGame struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
