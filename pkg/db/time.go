package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/nam-truong-le/lambda-utils-go/v4/pkg/aws/ssm"
	"github.com/nam-truong-le/lambda-utils-go/v4/pkg/mongodb"
)

const collectionTime = "time"

func CollectionTime(ctx context.Context) (*mongo.Collection, error) {
	database, err := ssm.GetParameter(ctx, "/mongo/db", false)
	if err != nil {
		return nil, err
	}
	return mongodb.Collection(ctx, database, collectionTime)
}

type TimeItem struct {
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
}

type Time struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Project string             `bson:"project" json:"project"`
	Date    time.Time          `bson:"date" json:"date"`
	Items   []TimeItem         `bson:"items" json:"items"`
}
