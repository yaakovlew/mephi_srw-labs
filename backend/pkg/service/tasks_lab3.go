package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	"math/rand"
	"time"
)

type tasksLab3 struct {
	repo *repository.Repo
}

func NewTaskLab3(repo *repository.Repo) *tasksLab3 {
	return &tasksLab3{
		repo: repo,
	}
}

func (s *tasksLab3) UpdateUserVarianceLab3(ctx context.Context, userId int, labId int, variance model.GeneratedLab3Variance) error {
	return s.repo.UpdateLab3Variance(userId, labId, variance)
}

func (s *tasksLab3) CheckLab3Variance(ctx context.Context, userId, labId int) error {
	return s.repo.CheckLab3Variance(userId, labId)
}

func (s *tasksLab3) GetVarianceLab3(ctx context.Context, userId, labId int) (model.UserLab3Task, [][]float64, error) {
	variance, err := s.repo.GetLab3Variance(userId, labId)
	if err != nil {
		return model.UserLab3Task{}, nil, err
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return model.UserLab3Task{}, nil, err
	}

	return s.convertToUserTask(ctx, variance), matrix, nil
}

func (s *tasksLab3) GenerateLab3Task(ctx context.Context, userId int) (model.GeneratedLab3Variance, error) {
	rand.Seed(time.Now().UnixNano())
	number, variantResponse, err := s.repo.GetRandomLab3VarianceFromBank()
	if err != nil {
		return model.GeneratedLab3Variance{}, err
	}

	var criteria []model.Criteria
	var sum float64 = 0
	for i, v := range variantResponse.Criteria {
		if len(variantResponse.Criteria)-1 == i {
			if sum > 0.92 {
				return model.GeneratedLab3Variance{}, fmt.Errorf("variantResponse has too many coefficients")
			}
			task := model.Criteria{
				Definition: v.Definition,
				ExtraInfo:  v.ExtraInfo,
				Func:       v.Func,
				RightFunc:  v.RightFunc,
				Gradation:  v.Gradation,
				Weight:     math.Round((1-sum)*100) / 100,
				FuncMark:   v.FuncMark,
			}
			criteria = append(criteria, task)
			continue
		}
		task := model.Criteria{
			Definition: v.Definition,
			ExtraInfo:  v.ExtraInfo,
			Func:       v.Func,
			RightFunc:  v.RightFunc,
			Gradation:  v.Gradation,
			Weight:     s.generateWeight(v.Weight),
			FuncMark:   v.FuncMark,
		}

		sum = sum + task.Weight
		criteria = append(criteria, task)
	}

	var alternative []model.Alternative
	for i := range variantResponse.Alternative {
		var alternativeCriteria []model.CriteriaCount
		for j := range variantResponse.Alternative[i].CriteriaCount {
			alternativeCriteria = append(alternativeCriteria, model.CriteriaCount{
				Count:         s.generateNotQualitativeAlternative(ctx, variantResponse.Alternative[i].CriteriaCount[j].Count),
				Value:         variantResponse.Alternative[i].CriteriaCount[j].Value,
				IsQualitative: variantResponse.Alternative[i].CriteriaCount[j].IsQualitative,
			})

		}

		alternativeExample := model.Alternative{
			Description:   variantResponse.Alternative[i].Description,
			CriteriaCount: alternativeCriteria,
		}
		alternative = append(alternative, alternativeExample)
	}

	variance := model.GeneratedLab3Variance{
		Number:                      number,
		Task:                        variantResponse.Task,
		Criteria:                    criteria,
		Alternative:                 alternative,
		Rule:                        variantResponse.Rule,
		ImportanceCriteriaMatrix:    variantResponse.ImportanceCriteriaMatrix,
		ImportanceAlternativeMatrix: variantResponse.ImportanceAlternativeMatrix,
	}

	for i := range variantResponse.ImportanceCriteriaMatrix {
		var lastVal float64 = 0
		for j := range variantResponse.ImportanceCriteriaMatrix[i].Points {
			variantResponse.ImportanceCriteriaMatrix[i].Points[j].X = s.generateImportancePoint(ctx, variantResponse.ImportanceCriteriaMatrix[i].Points[j].X)
			if variantResponse.ImportanceCriteriaMatrix[i].Points[j].X < lastVal {
				return model.GeneratedLab3Variance{}, fmt.Errorf("invalid importance point")
			}
			lastVal = variantResponse.ImportanceCriteriaMatrix[i].Points[j].X
		}
	}

	for i := range variantResponse.ImportanceAlternativeMatrix {
		var lastVal float64 = 0
		for j := range variantResponse.ImportanceAlternativeMatrix[i].Points {
			variantResponse.ImportanceAlternativeMatrix[i].Points[j].X = s.generateImportancePoint(ctx, variantResponse.ImportanceAlternativeMatrix[i].Points[j].X)
			if variantResponse.ImportanceAlternativeMatrix[i].Points[j].X < lastVal {
				return model.GeneratedLab3Variance{}, fmt.Errorf("invalid importance point")
			}
			lastVal = variantResponse.ImportanceAlternativeMatrix[i].Points[j].X
		}
	}

	variance.ImportanceCriteriaMatrix = variantResponse.ImportanceCriteriaMatrix
	variance.ImportanceAlternativeMatrix = variantResponse.ImportanceAlternativeMatrix

	return variance, nil
}

func (s *tasksLab3) convertToUserTask(ctx context.Context, variantResponse model.GeneratedLab3Variance) model.UserLab3Task {
	var criteria []model.UserCriteria
	for _, v := range variantResponse.Criteria {
		task := model.UserCriteria{
			Definition: v.Definition,
			ExtraInfo:  v.ExtraInfo,
			Func:       v.Func,
			Weight:     v.Weight,
		}
		for i := range v.FuncMark {
			task.FuncMark = append(task.FuncMark, model.UserFuncMark{
				Name: v.FuncMark[i].Name,
				Func: v.FuncMark[i].Func,
			})
		}
		criteria = append(criteria, task)
	}

	var alternative []model.UserAlternative
	for i := range variantResponse.Alternative {
		var alternativeCriteria []model.UserCriteriaCount
		for j := range variantResponse.Alternative[i].CriteriaCount {
			if variantResponse.Alternative[i].CriteriaCount[j].IsQualitative {
				alternativeCriteria = append(alternativeCriteria, model.UserCriteriaCount{
					Count: s.generateLevelQualitativeAlternative(ctx, variantResponse.Alternative[i].CriteriaCount[j].Count, variantResponse.Criteria[j].Gradation),
					Value: variantResponse.Alternative[i].CriteriaCount[j].Value,
				})
			} else {
				alternativeCriteria = append(alternativeCriteria, model.UserCriteriaCount{
					Count: variantResponse.Alternative[i].CriteriaCount[j].Count,
					Value: variantResponse.Alternative[i].CriteriaCount[j].Value,
				})
			}
		}

		alternativeExample := model.UserAlternative{
			Description:   variantResponse.Alternative[i].Description,
			CriteriaCount: alternativeCriteria,
		}
		alternative = append(alternative, alternativeExample)
	}

	var rules []model.UserRule
	for _, v := range variantResponse.Rule {
		rules = append(rules, model.UserRule{Name: v.Name})
	}

	return model.UserLab3Task{
		Number:                      variantResponse.Number,
		Task:                        variantResponse.Task,
		Criteria:                    criteria,
		Alternative:                 alternative,
		Rule:                        rules,
		ImportanceAlternativeMatrix: variantResponse.ImportanceAlternativeMatrix,
		ImportanceCriteriaMatrix:    variantResponse.ImportanceCriteriaMatrix,
	}
}

func (s *tasksLab3) generateUserTask(ctx context.Context, variantResponse model.GeneratedLab3Variance) model.UserLab3Task {
	var criteria []model.UserCriteria

	for _, v := range variantResponse.Criteria {
		task := model.UserCriteria{
			Definition: v.Definition,
			ExtraInfo:  v.ExtraInfo,
			Func:       v.Func,
			Weight:     s.generateWeight(v.Weight),
		}
		criteria = append(criteria, task)
	}

	var alternative []model.UserAlternative
	for i := range variantResponse.Alternative {
		var alternativeCriteria []model.UserCriteriaCount
		for j := range variantResponse.Alternative[i].CriteriaCount {
			if variantResponse.Alternative[i].CriteriaCount[j].IsQualitative {
				alternativeCriteria = append(alternativeCriteria, model.UserCriteriaCount{
					Count: s.generateLevelQualitativeAlternative(ctx, variantResponse.Alternative[i].CriteriaCount[j].Count, variantResponse.Criteria[j].Gradation),
					Value: variantResponse.Alternative[i].CriteriaCount[j].Value,
				})
			} else {
				alternativeCriteria = append(alternativeCriteria, model.UserCriteriaCount{
					Count: s.generateNotQualitativeAlternative(ctx, variantResponse.Alternative[i].CriteriaCount[j].Count),
					Value: variantResponse.Alternative[i].CriteriaCount[j].Value,
				})
			}
		}

		alternativeExample := model.UserAlternative{
			Description:   variantResponse.Alternative[i].Description,
			CriteriaCount: alternativeCriteria,
		}
		alternative = append(alternative, alternativeExample)
	}

	return model.UserLab3Task{
		Number:      variantResponse.Number,
		Task:        variantResponse.Task,
		Criteria:    criteria,
		Alternative: alternative,
	}

}

func (s *tasksLab3) generateLevelQualitativeAlternative(ctx context.Context, val float64, gradation []model.Gradation) string {
	for i := range gradation {
		if val >= gradation[i].StartPointX && val < gradation[i].EndPointX {
			return gradation[i].Name
		}
	}

	if val >= gradation[len(gradation)-1].StartPointX {
		return gradation[len(gradation)-1].Name
	} else if val <= 0 {
		return gradation[0].Name
	}

	return gradation[len(gradation)/2].Name
}

func (s *tasksLab3) generateWeight(val float64) float64 {
	rand.Seed(time.Now().UnixNano())
	percentage := rand.Intn(10)

	count := math.Round(val*float64(100+percentage)) / 100

	if count >= 0.9 {
		return 0.9
	}
	if count <= 0.08 {
		return 0.08
	}

	return count
}

func (s *tasksLab3) generateImportancePoint(ctx context.Context, val float64) float64 {
	rand.Seed(time.Now().UnixNano())
	percentage := rand.Intn(10)

	count := math.Round(val*float64(100+percentage)) / 100

	if count >= 1 {
		return 1
	}
	if count <= 0.08 {
		return 0.08
	}

	return count
}

func (s *tasksLab3) generateNotQualitativeAlternative(ctx context.Context, val float64) float64 {
	rand.Seed(time.Now().UnixNano())
	percentage := rand.Intn(10)

	return math.Round(val*float64(100+percentage)) / 100
}

func (s *tasksLab3) createMatrixByCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	result, err := s.getResultOfCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *tasksLab3) getResultOfCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	var matrix [][]float64

	for i := range variance.Criteria {
		var criteria []float64
		for j := range variance.Alternative {
			if variance.Alternative[j].CriteriaCount[i].IsQualitative {
				var x float64
				for _, v := range variance.Criteria[i].Gradation {
					if variance.Alternative[j].CriteriaCount[i].Count < v.EndPointX && variance.Alternative[j].CriteriaCount[i].Count >= v.StartPointX {
						x = v.StartPointX
						break
					}
				}
				res, err := s.funcEval(variance.Criteria[i].RightFunc, x)
				if err != nil {
					return nil, err
				}

				criteria = append(criteria, math.Round(res*100)/100)
			} else {
				res, err := s.funcEval(variance.Criteria[i].RightFunc, variance.Alternative[j].CriteriaCount[i].Count)
				if err != nil {
					return nil, err
				}

				criteria = append(criteria, math.Round(res*100)/100)
			}
		}
		matrix = append(matrix, criteria)
	}

	return matrix, nil
}

func (s *tasksLab3) funcEval(expString string, count float64) (float64, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"log": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("log function requires exactly two arguments")
			}

			base, ok := args[0].(float64)
			if !ok {
				return nil, fmt.Errorf("first argument of log function must be a number")
			}

			value, ok := args[1].(float64)
			if !ok {
				return nil, fmt.Errorf("second argument of log function must be a number")
			}

			if base <= 0 {
				return nil, fmt.Errorf("base of log must be greater than 0")
			}

			if value <= 0 {
				return nil, fmt.Errorf("value of log must be greater than 0")
			}

			result := math.Log(value) / math.Log(base)
			return result, nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)
	if err != nil {
		return 0, err
	}

	result, err := expression.Evaluate(map[string]interface{}{"x": count})
	if err != nil {
		return 0, err
	}

	val, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("can't get result of expression")
	}

	return val, nil
}
