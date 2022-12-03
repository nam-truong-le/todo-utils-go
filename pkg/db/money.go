package db

import (
	"context"
	"time"

	"github.com/nam-truong-le/lambda-utils-go/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionMoney = "money"

type MoneyGroup struct {
	Paid    bool      `bson:"paid" json:"paid"`
	Created time.Time `bson:"created" json:"created"`
}

type Money struct {
	Amount      string      `bson:"amount" json:"amount"`
	Partner     string      `bson:"partner" json:"partner"`
	Currency    string      `bson:"currency" json:"currency"`
	Description string      `bson:"description" json:"description"`
	Group       *MoneyGroup `bson:"group" json:"group"`
	Created     time.Time   `bson:"created" json:"created"`
}

func CollectionMoney(ctx context.Context) (*mongo.Collection, error) {
	return mongodb.Collection(ctx, collectionMoney)
}
