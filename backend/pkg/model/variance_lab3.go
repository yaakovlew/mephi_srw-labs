package model

type Variance struct {
	Task                        string             `json:"task"`
	Criteria                    []Criteria         `json:"criteria"`
	Alternative                 []Alternative      `json:"alternative"`
	ImportanceCriteriaMatrix    []ImportancePoints `json:"importance_criteria"`
	ImportanceAlternativeMatrix []ImportancePoints `json:"importance_alternative"`
	Rule                        []Rule             `json:"rule"`
}

type Rule struct {
	Name string `json:"name"`
	Func string `json:"func"`
}

type GeneratedLab3Variance struct {
	Number                      int                `json:"number"`
	Task                        string             `json:"task"`
	Criteria                    []Criteria         `json:"criteria"`
	Alternative                 []Alternative      `json:"alternative"`
	Rule                        []Rule             `json:"rule"`
	ImportanceCriteriaMatrix    []ImportancePoints `json:"importance_criteria"`
	ImportanceAlternativeMatrix []ImportancePoints `json:"importance_alternative"`
}

type Alternative struct {
	Description   string          `json:"description"`
	CriteriaCount []CriteriaCount `json:"criteria_count"`
}

type CriteriaCount struct {
	Count         float64 `json:"count"`
	IsQualitative bool    `json:"is_qualitative"`
	Value         string  `json:"value"`
}

type Gradation struct {
	Name        string  `json:"name"`
	StartPointX float64 `json:"start_point"`
	EndPointX   float64 `json:"end_point"`
}

type Criteria struct {
	Definition string      `json:"definition"`
	ExtraInfo  string      `json:"extra_info"`
	Func       string      `json:"func"`
	RightFunc  string      `json:"right_func"`
	FuncMark   []FuncMark  `json:"func_mark"`
	Gradation  []Gradation `json:"gradation"`
	Weight     float64     `json:"weight"`
}

type FuncMark struct {
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Func     string `json:"func"`
}

type UserFuncMark struct {
	Name string `json:"name"`
	Func string `json:"func"`
}

type UserCriteria struct {
	Definition string         `json:"definition"`
	ExtraInfo  string         `json:"extra_info"`
	Func       string         `json:"func"`
	Weight     float64        `json:"weight"`
	FuncMark   []UserFuncMark `json:"func_mark"`
}

type UserLab3Task struct {
	Number                      int                `json:"number"`
	Task                        string             `json:"task"`
	Criteria                    []UserCriteria     `json:"criteria"`
	Alternative                 []UserAlternative  `json:"alternative"`
	Rule                        []UserRule         `json:"rule"`
	ImportanceCriteriaMatrix    []ImportancePoints `json:"importance_criteria"`
	ImportanceAlternativeMatrix []ImportancePoints `json:"importance_alternative"`
}

type UserAlternative struct {
	Description   string              `json:"description"`
	CriteriaCount []UserCriteriaCount `json:"criteria_count"`
}

type UserRule struct {
	Name string `json:"name"`
}

type UserCriteriaCount struct {
	Count interface{} `json:"count"`
	Value string      `json:"value"`
}

type DbBankVariance struct {
	Variance Variance `json:"variance"`
}
