package requests

type TrainingRequest struct {
	Seg []TrainingDetails `copier:"Seg" json:"seg"`
	Ter []TrainingDetails `copier:"Ter" json:"ter"`
	Qua []TrainingDetails `copier:"Qua" json:"qua"`
	Qui []TrainingDetails `copier:"Qui" json:"qui"`
	Sex []TrainingDetails `copier:"Sex" json:"sex"`
	Sab []TrainingDetails `copier:"Sab" json:"sab"`
	Dom []TrainingDetails `copier:"Dom" json:"dom"`
}

type TrainingDetails struct {
	Muscle      string `copier:"Muscle" json:"muscle"`
	Activity    string `copier:"Activity" json:"activity"`
	Series      string `copier:"Series" json:"series"`
	Repetitions string `copier:"Repetitions" json:"repetitions"`
}
