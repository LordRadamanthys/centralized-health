package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type VaccinesDomain struct {
	ID        primitive.ObjectID `copier:"ID" bson:"_id,omitempty" json:"_id"`
	IdUser    primitive.ObjectID `copier:"IdUser" json:"id_user"`
	Title     string             `copier:"Title" json:"title"`
	Address   string             `copier:"Address" json:"address"`
	Doses     string             `copier:"Doses" json:"doses"`
	Dates     []string           `copier:"Dates" json:"dates"`
	Documents []DocumentsDomain  `copier:"Documents" json:"documents"`
}
