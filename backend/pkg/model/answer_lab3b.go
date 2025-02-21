package model

type AnswerLab3bRulesValue struct {
	Matrices map[string][]float64 `json:"matrices" binding:"required"`
	Step     int                  `json:"step" binding:"required"`
}

type AnswerLab3bRulesNumber struct {
	Matrix [][]float64 `json:"matrix" binding:"required"`
}

type AnswerLab3bAllMatrices struct {
	Matrices [][]Point `json:"matrices" binding:"required"`
	Step     int       `json:"step" binding:"required"`
}

type AnswerLab3bMatricesIntersection struct {
	Matrix [][]Point `json:"matrix" binding:"required"`
}

type AnswerLab3bLevelSet struct {
	AnswerLevelSet []AnswerLevelSet `json:"answer_level_set" binding:"required"`
	Step           int              `json:"step" binding:"required"`
}

type AnswerLevelSet struct {
	Set      []float64 `json:"set"`
	Delta    float64   `json:"delta"`
	Powerful float64   `json:"powerful"`
}

type AnswerLab3bLevelSetResponse struct {
	AnswerLevelSet []AnswerLevelResponse `json:"answer_level_set""`
}

type AnswerLevelResponse struct {
	Set      []DataResponse `json:"set"`
	Delta    DataResponse   `json:"delta"`
	Powerful DataResponse   `json:"powerful"`
}

type AnswerLab3bResult struct {
	Set   []float64 `json:"set" binding:"required"`
	Index int       `json:"index"`
}
