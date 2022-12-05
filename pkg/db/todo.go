package db

import (
	"context"
	"time"

	"github.com/nam-truong-le/lambda-utils-go/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	return mongodb.Collection(ctx, collectionTodo)
}
