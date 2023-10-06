package modals

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user Schema
type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId       string             `json:"userId" bson:"userId"`
	FirstName    string             `json:"firstName" bson:"firstName" validate:"required,min=3,max=30"`
	LastName     string             `json:"lastName" bson:"lastName" validate:"required,min=3,max=30"`
	Email        string             `json:"email" bson:"email" validate:"required,email"`
	MobileNumber string             `json:"mobileNumber" bson:"mobileNumber" validate:"required,min=10,max=10"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Address      string             `json:"address" bson:"address"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
