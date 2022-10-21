package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDomain struct {
	ID       primitive.ObjectID `copier:"ID" json:"_id"`
	Name     string             `copier:"Name" json:"name"`
	Password string             `copier:"Password" json:"password"`
	Email    string             `copier:"Email" json:"email"`
	Birth    string             `copier:"Birth" json:"birth"`
	Phones   []string           `copier:"Phone" json:"phones"`
}
