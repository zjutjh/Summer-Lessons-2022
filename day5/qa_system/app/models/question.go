package models

type Question struct {
	QuestionnaireID int
	Stem            string
	TypeNum         int
	Options         string
	Answer          string
}
