package model

type AnswerLab3aAlternativeSets struct {
	Sets [][]float64 `json:"sets" binding:"required"`
}

type AnswerLab3aDiffMatrices struct {
	Matrices [][]float64 `json:"matrices" binding:"required"`
	Step     int         `json:"step" binding:"required"`
}

type AnswerLab3aMatricesIntersection struct {
	Matrix [][]float64 `json:"matrix" binding:"required"`
}

type AnswerLab3aCheckNonDominatedSet struct {
	Set []float64 `json:"set" binding:"required"`
}

type AnswerLab3aMatricesWithCoefficients struct {
	Matrix [][]float64 `json:"matrix" binding:"required"`
}

type AnswerLab3aSecondNonDominatedSets struct {
	Set []float64 `json:"set" binding:"required"`
}

type AnswerLab3aResult struct {
	Set         []float64 `json:"set" binding:"required"`
	ChosenIndex int       `json:"chosen_index"`
}