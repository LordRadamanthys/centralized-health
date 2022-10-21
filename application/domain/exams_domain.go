package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExamsDomain struct {
	IdUser        primitive.ObjectID `copier:"IdUser" json:"id_user"`
	DoctorName    string             `copier:"DoctorName" json:"doctor_name"`
	Description   string             `copier:"Description" json:"description"`
	DateVisit     string             `copier:"DateVisit" json:"date_visit"`
	DateReturn    string             `copier:"DateReturn" json:"date_return"`
	Speciality    string             `copier:"Speciality" json:"speciality"`
	ListMedicines []MedicinesStruct  `copier:"ListMedicines" json:"list_medicines"`
	Documents     []DocumentsDomain
}

type MedicinesStruct struct {
	Title string `copier:"Title" json:"title"`
	Rules string `copier:"Rules" json:"rules"`
}
