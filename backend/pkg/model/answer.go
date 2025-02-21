package model

type UserAnswer struct {
	UserID int    `json:"user_id"`
	LabId  int    `json:"laboratory_id"`
	Answer Answer `json:"answer"`
}

type Answer struct{}
