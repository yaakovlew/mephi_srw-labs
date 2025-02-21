package model

type GeneratedLab1AVariance struct {
	Number   int           `json:"number" db:"id"`
	Variance Lab1AVariance `json:"variance" db:"variance"`
}

type Lab1AVariance struct {
	Matrices [][][]float64 `json:"matrices"`
}
