package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrainingDomain struct {
	IdUser primitive.ObjectID `copier:"IdUser" bson:"id_user,omitempty" json:"id_user"`
	Seg    []TrainingDetails  `copier:"Seg" bson:"seg,omitempty" json:"seg"`
	Ter    []TrainingDetails  `copier:"Ter" bson:"ter,omitempty" json:"ter"`
	Qua    []TrainingDetails  `copier:"Qua" bson:"qua,omitempty" json:"qua"`
	Qui    []TrainingDetails  `copier:"Qui" bson:"qui,omitempty" json:"qui"`
	Sex    []TrainingDetails  `copier:"Sex" bson:"sex,omitempty" json:"sex"`
	Sab    []TrainingDetails  `copier:"Sab" bson:"sab,omitempty" json:"sab"`
	Dom    []TrainingDetails  `copier:"Dom" bson:"dom,omitempty" json:"dom"`
}

type TrainingDetails struct {
	Muscle      string `copier:"Muscle" json:"muscle"`
	Activity    string `copier:"Activity" json:"Activity"`
	Series      string `copier:"Series" json:"series"`
	Repetitions string `copier:"Repetitions" json:"repetitions"`
}
