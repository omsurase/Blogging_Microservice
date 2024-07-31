package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  *string            `bson:"username" json:"username" validate:"required,min=3,max=30"`
	Email     *string            `bson:"email" json:"email" validate:"required,email"`
	Password  *string            `bson:"password" json:"-" validate:"required,min=8"`
	FirstName *string            `bson:"first_name" json:"first_name" validate:"required"`
	LastName  *string            `bson:"last_name" json:"last_name" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Token     *string            `bson:"token" json:"token" validate:"required,jwt"`
	Bio       *string            `bson:"bio" json:"bio"`
}
