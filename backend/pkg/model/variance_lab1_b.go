package model

type GeneratedVarianceLab1B struct {
	Number   int           `json:"number" db:"id"`
	Variance VarianceLab1B `json:"variance" db:"variance"`
}

type VarianceLab1B struct {
	Task         string         `json:"task"`
	MainCriteria []MainCriteria `json:"main_criteria"`
	Alternatives []string       `json:"alternatives"`
}

type MainCriteria struct {
	Criteria string          `json:"criteria"`
	Extra    []ExtraCriteria `json:"extra"`
}

type ExtraCriteria struct {
	Criteria  string    `json:"criteria"`
	IsCount   bool      `json:"is_count"`
	Value     []float64 `json:"value"`
	IsReverse bool      `json:"is_reverse"`
}

type UserVarianceLab1B struct {
	Variance              GeneratedVarianceLab1B `json:"variance"`
	MainCriteriaMatrix    [][]float64            `json:"main_matrix"`
	CriteriaMatrix        [][][]float64          `json:"criteria_matrix"`
	QualityCriteriaMatrix [][][]float64          `json:"quality_matrix"`
}
