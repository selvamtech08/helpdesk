package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// hold the all users details
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name,omitempty" bson:"name"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Role      string             `json:"role,omitempty" bson:"role"`
	Group     string             `json:"group,omitempty" bson:"group"`
	IsActive  bool               `json:"is_active,omitempty" bson:"is_active"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
}

// update bt user
type UpdateUser struct {
	Name      string     `json:"name,omitempty" bson:"name"`
	Email     string     `json:"email,omitempty" bson:"email"`
	Password  string     `json:"password,omitempty" bson:"password"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
