package model

type AnswerLab3AlternativeImportance struct {
	Matrix []string `json:"matrix" binding:"required"`
	Step   int      `json:"step" binding:"required"`
}

type AnswerLab3CriteriaImportance struct {
	Set []string `json:"set" binding:"required"`
}

type AnswerLab3CEstimation struct {
	Matrix []Point `json:"matrix" binding:"required"`
	Step   int     `json:"step" binding:"required"`
}

type AnswerLab3CCurrentMatrix struct {
	Matrix [][]Point `json:"matrix" binding:"required"`
	Step   int       `json:"step" binding:"required"`
}

type AnswerLab3CAlternativeMatrices struct {
	Matrices [][]Point `json:"matrices" binding:"required"`
	Step     int       `json:"step" binding:"required"`
}

type AnswerLab3CCriteriaMatrix struct {
	Matrix [][]Point `json:"matrix" binding:"required"`
}

type AnswerLab3CArea struct {
	Set  float64 `json:"set"`
	Step int     `json:"step" binding:"required"`
}

type AnswerLab3CLineParameters struct {
	Parameters LineParameters `json:"parameters"`
}

type AnswerLab3CLineParametersRequest struct {
	Parameters LineParametersRequest `json:"parameters" binding:"required"`
	Step       int                   `json:"step" binding:"required"`
}

type QuadraticParameters struct {
	A1 DataResponse `json:"a1"`
	A2 DataResponse `json:"a2"`
	A3 DataResponse `json:"a3"`
}

type QuadraticParametersRequest struct {
	A1 float64 `json:"a1"`
	A2 float64 `json:"a2"`
	A3 float64 `json:"a3"`
}

type AnswerLab3CMiddleOfArea struct {
	Set   []float64 `json:"set" binding:"required"`
	Index int       `json:"index"`
}

type LineParameters struct {
	K DataResponse `json:"k"`
	B DataResponse `json:"b"`
}

type LineParametersRequest struct {
	K float64 `json:"k"`
	B float64 `json:"b"`
}

type AnswerLab3CQuadraticParameters struct {
	Parameters QuadraticParameters `json:"parameters"`
}

type AnswerLab3CQuadraticParametersRequest struct {
	Parameters QuadraticParametersRequest `json:"parameters" binding:"required"`
	Step       int                        `json:"step" binding:"required"`
}
