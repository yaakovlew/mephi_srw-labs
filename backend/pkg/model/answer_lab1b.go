package model

type AnswerLab1BResult struct {
	Set         []float64 `json:"set" binding:"required"`
	ChosenIndex int       `json:"chosen_index"`
}

type AnswerMatrixTaskTask struct {
	Task string `json:"task"`
}

type AnswerMatrixTaskAlternative struct {
	Alternatives []string `json:"alternatives"`
}

type AnswerMatrixTaskCriteria struct {
	Criteria []MainCriteria `json:"criteria"`
}

type Lab1BMatrix struct {
	Matrix [][]float64 `json:"matrix"`
}

type AnswerLab1BCommonMatrix struct {
	Matrix    [][]float64 `json:"matrix"`
	X         []float64   `json:"x"`
	W         []float64   `json:"w"`
	MW        []float64   `json:"mw"`
	LambdaW   []float64   `json:"lambda_w"`
	LambdaMax float64     `json:"lambda_max"`
	IS        float64     `json:"is"`
	OS        float64     `json:"os"`
}

type Lab1BMarkAligning struct {
	WIS    []float64 `json:"w_is"`
	WCC    []float64 `json:"w_cc"`
	M      float64   `json:"m"`
	TildaM float64   `json:"tilda_m"`
	OSI    float64   `json:"osi"`
}

type Lab1BWeights struct {
	Weights []float64 `float:"weights"`
}

type Lab1BCountCriteria struct {
	Marks [][]float64 `json:"marks"`
}

type Lab1BCountCriteriaISRight struct {
	Marks [][]IsRight `json:"marks"`
}

type Lab1BWeightsIsRight struct {
	Weights []IsRight `float:"weights"`
}

type Lab1BMarkAligningIsRight struct {
	WIS    []IsRight `json:"w_is"`
	WCC    []IsRight `json:"w_cc"`
	M      IsRight   `json:"m"`
	TildaM IsRight   `json:"tilda_m"`
	OSI    IsRight   `json:"osi"`
}

type AnswerLab1BCommonMatrixIsRight struct {
	X         []IsRight `json:"x"`
	W         []IsRight `json:"w"`
	MW        []IsRight `json:"mw"`
	LambdaW   []IsRight `json:"lambda_w"`
	LambdaMax IsRight   `json:"lambda_max"`
	IS        IsRight   `json:"is"`
	OS        IsRight   `json:"os"`
}

type AnswerLab1B5Step struct {
	W   float64 `json:"w"`
	IS  float64 `json:"is"`
	WIS float64 `json:"w_is"`
	CC  float64 `json:"cc"`
	WCC float64 `json:"w_cc"`
}

type AnswerLab1B5StepIsRight struct {
	WIS IsRight `json:"w_is"`
	CC  IsRight `json:"cc"`
	WCC IsRight `json:"w_cc"`
}

type AnswerLab1B6Step struct {
	M      float64 `json:"m"`
	TildaM float64 `json:"tilda_m"`
	OSI    float64 `json:"osi"`
}

type AnswerLab1B6StepIsRight struct {
	M      IsRight `json:"m"`
	TildaM IsRight `json:"tilda_m"`
	OSI    IsRight `json:"osi"`
}

type AnswerLab1B7Step struct {
	Weights []float64 `json:"weights"`
}

type AnswerLab1B7StepIsRight struct {
	Weights []IsRight `json:"weights"`
}

type AnswerLab1B9Step struct {
	Set   []float64 `json:"set"`
	Index int       `json:"index"`
}

type AnswerLab1A9StepIsRight struct {
	Set   []IsRight  `json:"set"`
	Index IsRightInt `json:"index"`
}
