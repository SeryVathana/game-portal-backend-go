package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GamePlayer struct {
	Id           primitive.ObjectID    `bson:"_id,omitempty" json:"_id"`
	UserId       primitive.ObjectID    `bson:"user_id" json:"user_id"`
	GameId       primitive.ObjectID    `bson:"game_id" json:"game_id"`
	DisplayName  string                `bson:"display_name" json:"display_name"`
	AvatarImgURL string                `bson:"avatar_img_url" json:"avatar_img_url"`
	Language     string                `bson:"language" json:"language"`
	Currency     []GamePlayerCurrency  `bson:"currency" json:"currency"`
	Inventory    []GamePlayerInventory `bson:"inventory" json:"inventory"`
	Statistic    []GamePlayerStatistic `bson:"statistic" json:"statistic"`
	CreatedAt    primitive.DateTime    `bson:"created_at" json:"created_at"`
	UpdatedAt    primitive.DateTime    `bson:"updated_at" json:"updated_at"`
}

type GamePlayerCurrency struct {
	Code  string `bson:"code" json:"code"`
	Value int    `bson:"value" json:"value"`
}

type GamePlayerInventory struct {
	Key   string `bson:"key" json:"key"`
	Value int    `bson:"value" json:"value"`
}

type GamePlayerStatistic struct {
	Code  string `bson:"code" json:"code"`
	Value int    `bson:"value" json:"value"`
}

type GamePlayerRequest struct {
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	Token     string             `bson:"token" json:"token"`
	SecretKey string             `bson:"secret_key" json:"secret_key"`
}

type GamePlayerData struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	Data      []GameData         `bson:"game_data" json:"game_data"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type GameData struct {
	Key   string `bson:"key" json:"key"`
	Value int    `bson:"value" json:"value"`
}

type GamePlayerPurchase struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	ItemId    primitive.ObjectID `bson:"item_id" json:"item_id"`
	Price     []PurchasePrice    `bson:"price" json:"price"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type PurchasePrice struct {
	Key   string `bson:"key" json:"key"`
	Value int    `bson:"value" json:"value"`
}

type GamePlayerSession struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	StartedAt primitive.DateTime `bson:"started_at" json:"started_at"`
	StoppedAt primitive.DateTime `bson:"stopped_at" json:"stopped_at"`
}
