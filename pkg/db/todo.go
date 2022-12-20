package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/nam-truong-le/lambda-utils-go/pkg/aws/ssm"
	"github.com/nam-truong-le/lambda-utils-go/pkg/mongodb"
)

const collectionTodo = "todo"

type Todo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Group       string             `bson:"group" json:"group"`
	Description string             `bson:"description" json:"description"`
	Priority    string             `bson:"priority" json:"priority"`
	Deadline    *time.Time         `bson:"deadline,omitempty" json:"deadline,omitempty"`
	Finished    bool               `bson:"finished" json:"finished"`
	Created     time.Time          `bson:"created" json:"created"`
}

func CollectionTodo(ctx context.Context) (*mongo.Collection, error) {
	database, err := ssm.GetParameter(ctx, "/mongo/db", false)
	if err != nil {
		return nil, err
	}
	return mongodb.Collection(ctx, database, collectionTodo)
}
