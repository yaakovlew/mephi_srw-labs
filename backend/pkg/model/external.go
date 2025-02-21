package model

type UserId struct {
	UserId int `json:"user_id" binding:"required"`
}

type NewToken struct {
	Token string `json:"token" binding:"required"`
}
