package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	DevId         primitive.ObjectID   `bson:"dev_id" json:"dev_id"`
	Title         string               `bson:"title" json:"title"`
	Description   string               `bson:"description" json:"description"`
	Genres        []primitive.ObjectID `bson:"genres" json:"genres"`
	ImgURL        string               `bson:"img_url" json:"img_url"`
	VideoURL      string               `bson:"video_url" json:"video_url"`
	GameURL       string               `bson:"game_url" json:"game_url"`
	SecretKey     string               `bson:"secret_key" json:"secret_key"`
	EncryptionKey string               `bson:"encrption_key" json:"encrption_key"`
	CreatedAt     primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt     primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}

type GameRating struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	Star      int                `bson:"star" json:"star"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type GameComment struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId     primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId     primitive.ObjectID `bson:"game_id" json:"game_id"`
	Comment    string             `bson:"comment" json:"comment"`
	ReplyCmtId primitive.ObjectID `bson:"reply_cmt_id,omitempty" json:"reply_cmt_id,omitempty"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type GameFeedback struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	GameId    primitive.ObjectID `bson:"game_id" json:"game_id"`
	Feedback  string             `bson:"feedback" json:"feedback"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
