package store

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/selvamtech08/helpdesk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

var User = userImpl{
	collection: GetCollection("users"),
	ctx:        context.TODO(),
}

// insert new user detail in db and return the entires added db
// return error if db update get failed
func (u *userImpl) New(user model.User) error {
	user.ID = primitive.NewObjectID()
	result, err := u.collection.InsertOne(u.ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user name alreay exisits in database")
		}
		return err
	}
	log.Println("user document added:", result.InsertedID)
	return nil
}

func (u *userImpl) GetUserByName(userName string) (*model.User, error) {
	var user model.User
	filter := bson.M{"name": userName}
	result := u.collection.FindOne(u.ctx, filter)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userImpl) UpdateUser(userName string, user model.UpdateUser) error {
	filter := bson.M{"name": userName}
	var updates = bson.D{}

	if user.Name != "" {
		updates = append(updates, bson.E{Key: "name", Value: user.Name})
	}
	if user.Email != "" {
		updates = append(updates, bson.E{Key: "email", Value: user.Email})
	}
	if user.Password != "" {
		hashPassword, err := u.HashPassword(user.Password)
		if err == nil {
			updates = append(updates, bson.E{Key: "password", Value: hashPassword})
		}
	}

	updates = append(updates, bson.E{Key: "updated_at", Value: time.Now()})

	result, err := u.collection.UpdateOne(u.ctx, filter, bson.M{"$set": updates})
	if err != nil {
		return err
	}
	log.Printf("user %s modified: %d\n", userName, result.ModifiedCount)
	return nil
}
