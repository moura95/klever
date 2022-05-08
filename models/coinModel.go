package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coin struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CoinName  string             `json:"coinName" bson:"coinName"`
	Upvote    int64              `json:"upvote" bson:"upvote"`
	Downvote  int64              `json:"downvote" bson:"downvote"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
