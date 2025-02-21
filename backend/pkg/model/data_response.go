package model

type DataResponse struct {
	Data interface{} `json:"data"`
	Flag bool        `json:"flag"`
}
