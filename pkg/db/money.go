package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/nam-truong-le/lambda-utils-go/pkg/aws/ssm"
	"github.com/nam-truong-le/lambda-utils-go/pkg/mongodb"
)

const collectionMoney = "money"

type MoneyGroup struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Paid    bool               `bson:"paid" json:"paid"`
	Created time.Time          `bson:"created" json:"created"`
}

type Money struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Amount      string             `bson:"amount" json:"amount"`
	Partner     string             `bson:"partner" json:"partner"`
	Currency    string             `bson:"currency" json:"currency"`
	Description string             `bson:"description" json:"description"`
	Group       *MoneyGroup        `bson:"group,omitempty" json:"group,omitempty"`
	Created     time.Time          `bson:"created" json:"created"`
}

func CollectionMoney(ctx context.Context) (*mongo.Collection, error) {
	database, err := ssm.GetParameter(ctx, "/mongo/db", false)
	if err != nil {
		return nil, err
	}
	return mongodb.Collection(ctx, database, collectionMoney)
}
