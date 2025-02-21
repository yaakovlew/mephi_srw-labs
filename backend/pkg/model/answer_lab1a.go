package model

type AnswerLab1aResult struct {
	Set         []float64 `json:"set" binding:"required"`
	ChosenIndex int       `json:"chosen_index"`
}

type AnswerLab1ACommonMatrix struct {
	X []float64 `json:"x"`
	W []float64 `json:"w"`
	MW []float64 `json:"mw"`
	LambdaW []float64 `json:"lambda_w"`
	LambdaMax float64 `json:"lambda_max"`
	IS float64 `json:"is"`
	OS float64 `json:"os"`
}

type IsRight struct {
	Val float64 `json:"val"`
	IsRight bool `json:"is_right"`
} 

type IsRightInt struct {
	Val int `json:"val"`
	IsRight bool `json:"is_right"`
} 

type AnswerLab1ACommonMatrixIsRight struct {
	X []IsRight `json:"x"`
	W []IsRight `json:"w"`
	MW []IsRight `json:"mw"`
	LambdaW []IsRight `json:"lambda_w"`
	LambdaMax IsRight `json:"lambda_max"`
	IS IsRight `json:"is"`
	OS IsRight `json:"os"`
}

type AnswerLab1A5Step struct {
	Set []float64 `json:"set"`
	Index int `json:"index"`
}

type AnswerLab1A5StepIsRight struct {
	Set []IsRight `json:"set"`
	Index IsRightInt `json:"index"`
}

