package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DtoSignUp struct {
	ID           primitive.ObjectID `bson:"_id"`
	Password     string             `json:"password" validate:"required,min=6"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserId       string             `json:"userid"`

	FirstName string `json:"firstname" validate:"required,min=2,max=100"`
	LastName  string `json:"lastname" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"email,required"`
	UserType  string `json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
}
