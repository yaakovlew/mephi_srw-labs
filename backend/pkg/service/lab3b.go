package service

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"sort"
	"strconv"
	"strings"

	"backend/pkg/model"
	"backend/pkg/repository"
)

const Lab3BId = 5

type lab3bService struct {
	repo *repository.Repo
	commonEventService
}

func NewLab3bService(repo *repository.Repo) *lab3bService {
	return &lab3bService{
		repo:               repo,
		commonEventService: NewCommonEventService(),
	}
}

func (s *lab3bService) ValidateLab3BResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error) {
	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 || matrix[i][j] > 1 {
				return nil, fmt.Errorf("matrix coefficients are not valid")
			}
		}
	}

	critVal, alterVal, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return nil, err
	}

	_, points := s.GetPointsByLukasiewiczImplication(ctx, critVal, alterVal)
	if err := s.validatePoints(ctx, points); err != nil {
		return nil, err
	}

	num := s.GetPointEstimateByAlternativeMatrix(ctx, points)

	return num, nil
}

func (s *lab3bService) NextMatrix(ctx context.Context, userId int) ([][]string, error) {
	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return nil, err
	}

	var data [][]string
	for _, criteria := range variance.Criteria {
		var dataToAdd []string
		for i := range criteria.FuncMark {
			dataToAdd = append(dataToAdd, criteria.FuncMark[i].Name)
		}
		data = append(data, dataToAdd)
	}

	return data, nil
}

func (s *lab3bService) CheckLab3BRulesValue(ctx context.Context, userId int, step int, userMatrix map[string][]float64) (int, int, map[string][]model.DataResponse, error) {
	maxMark := 5
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	matrix = s.transposeMatrix(ctx, matrix)
	if step >= len(matrix) {
		return 0, maxMark, nil, fmt.Errorf("step %d is out of range", step)
	}

	result, err := s.getRuleData(ctx, step, variance, matrix[step])
	if err != nil {
		return 0, maxMark, nil, err
	}

	for key := range result {
		if _, ok := userMatrix[key]; !ok {
			incorrect = incorrect + len(result[key])
			all = all + len(result[key])
			continue
		}
		if len(result[key]) != len(userMatrix[key]) {
			incorrect = incorrect + len(result[key])
			all = all + len(result[key])
			continue
		}

		for i := range result[key] {
			if result[key][i].Data != userMatrix[key][i] {
				incorrect++
			} else {
				result[key][i].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, result, nil
}

func (s *lab3bService) getRuleData(ctx context.Context, step int, variance model.GeneratedLab3Variance, matrix []float64) (map[string][]model.DataResponse, error) {
	mapka := make(map[string][]model.DataResponse)

	for j := range variance.Criteria[step].FuncMark {
		for k := range matrix {
			if data, err := s.eval(variance.Criteria[step].FuncMark[j].Func, math.Round(matrix[k]*100)/100); err != nil {
				return nil, err
			} else {
				mapka[variance.Criteria[step].FuncMark[j].Name] = append(mapka[variance.Criteria[step].FuncMark[j].Name], model.DataResponse{
					Data: math.Round(data*100) / 100,
					Flag: false,
				})
			}
		}
	}

	return mapka, nil
}

func (s *lab3bService) CheckLab3BCriteriaValue(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error) {
	maxMark := 20
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	critVal, _, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return 0, maxMark, nil, err
	}

	var data [][]model.DataResponse
	for i := range critVal {
		var dataToAdd []model.DataResponse
		for j := range critVal[i] {
			dataToAdd = append(dataToAdd, model.DataResponse{
				Data: critVal[i][j],
				Flag: false,
			})
		}
		data = append(data, dataToAdd)
	}

	if len(data) != len(userMatrix) {
		return 0, maxMark, data, nil
	}

	for i := range data {
		if len(data[i]) != len(userMatrix[i]) {
			return 0, maxMark, data, nil
		}
		for j := range data[i] {
			if data[i][j].Data != userMatrix[i][j] {
				incorrect++
			} else {
				data[i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3bService) CheckLab3BAllMatrices(ctx context.Context, userId int, step int, userPoints [][]model.Point) (int, int, [][]model.PointCheck, error) {
	maxMark := 10
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	critVal, alterVal, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return 0, maxMark, nil, err
	}

	allPoints, _ := s.GetPointsByLukasiewiczImplication(ctx, critVal, alterVal)
	var data [][][]model.PointCheck
	for i := range allPoints {
		var dataToAdd [][]model.PointCheck
		for j := range allPoints[i] {
			var dataToAdd1 []model.PointCheck
			for k := range allPoints[i][j] {
				dataToAdd1 = append(dataToAdd1, model.PointCheck{
					X:    allPoints[i][j][k].X,
					Y:    allPoints[i][j][k].Y,
					Flag: false,
				})
			}
			dataToAdd = append(dataToAdd, dataToAdd1)
		}
		data = append(data, dataToAdd)
	}

	if step >= len(data) {
		return 0, maxMark, nil, fmt.Errorf("step %d is out of range", step)
	}

	if len(data) != len(userPoints) {
		return 0, maxMark, data[step], nil
	}

	for i := range data[step] {
		if len(data[step][i]) != len(userPoints[i]) {
			return 0, maxMark, data[step], nil
		}
		for j := range data[step][i] {
			if data[step][i][j].X != userPoints[i][j].X || data[step][i][j].Y != userPoints[i][j].Y {
				incorrect++
			} else {
				data[step][i][j].Flag = true
			}

			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[step], nil
}

func (s *lab3bService) CheckLab3BMatricesIntersection(ctx context.Context, userId int, userPoints [][]model.Point) (int, int, [][]model.PointCheck, error) {
	maxMark := 15
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	critVal, alterVal, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return 0, maxMark, nil, err
	}

	_, points := s.GetPointsByLukasiewiczImplication(ctx, critVal, alterVal)

	var data [][]model.PointCheck
	for i := range points {
		var dataToAdd []model.PointCheck
		for j := range points[i] {
			dataToAdd = append(dataToAdd, model.PointCheck{
				X:    points[i][j].X,
				Y:    points[i][j].Y,
				Flag: false,
			})
		}
		data = append(data, dataToAdd)
	}

	if len(data) != len(userPoints) {
		return 0, maxMark, data, nil
	}

	for i := range data {
		if len(data[i]) != len(userPoints[i]) {
			return 0, maxMark, data, nil
		}
		for j := range data[i] {
			if data[i][j].X != userPoints[i][j].X || data[i][j].Y != userPoints[i][j].Y {
				incorrect++
			} else {
				data[i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3bService) CheckLab3BAnswerLab3bLevelSet(ctx context.Context, userId int, step int, userLevelSet []model.AnswerLevelSet) (int, int, model.AnswerLab3bLevelSetResponse, error) {
	maxMark := 10
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark, model.AnswerLab3bLevelSetResponse{}, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, model.AnswerLab3bLevelSetResponse{}, err
	}

	critVal, alterVal, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return 0, maxMark, model.AnswerLab3bLevelSetResponse{}, err
	}

	_, points := s.GetPointsByLukasiewiczImplication(ctx, critVal, alterVal)
	levelSets := s.getLevelSets(ctx, points)

	if step >= len(levelSets) {
		return 0, maxMark, model.AnswerLab3bLevelSetResponse{}, fmt.Errorf("step %d is out of range", step)
	}

	var data []model.AnswerLab3bLevelSetResponse
	for i := range levelSets {
		var dataToAdd []model.AnswerLevelResponse
		for j := range levelSets[i] {
			var dataToAdd2 []model.DataResponse
			for k := range levelSets[i][j].Set {
				dataToAdd2 = append(dataToAdd2, model.DataResponse{
					Data: levelSets[i][j].Set[k],
					Flag: false,
				})
			}
			dataToAdd = append(dataToAdd, model.AnswerLevelResponse{
				Set: dataToAdd2,
				Delta: model.DataResponse{
					Data: levelSets[i][j].Delta,
					Flag: false,
				},
				Powerful: model.DataResponse{
					Data: levelSets[i][j].Powerful,
					Flag: false,
				},
			})
		}

		data = append(data, model.AnswerLab3bLevelSetResponse{
			AnswerLevelSet: dataToAdd,
		})
	}

	if step >= len(data) {
		return 0, maxMark, model.AnswerLab3bLevelSetResponse{}, fmt.Errorf("step %d is out of range", step)
	}

	if len(data[step].AnswerLevelSet) != len(userLevelSet) {
		return 0, maxMark, data[step], nil
	}

	for i := range data[step].AnswerLevelSet {
		if userLevelSet[i].Delta != data[step].AnswerLevelSet[i].Delta.Data {
			incorrect++
		} else {
			data[step].AnswerLevelSet[i].Delta.Flag = true
		}

		if userLevelSet[i].Powerful != data[step].AnswerLevelSet[i].Powerful.Data {
			incorrect++
		} else {
			data[step].AnswerLevelSet[i].Powerful.Flag = true
		}

		var check1 []float64
		check2 := make([]float64, len(userLevelSet[i].Set))
		copy(check2, userLevelSet[i].Set)
		for j := range data[step].AnswerLevelSet {
			toAdd, ok := data[step].AnswerLevelSet[i].Set[j].Data.(float64)
			if !ok {
				continue
			}
			check1 = append(check1, toAdd)
		}

		if s.arrayIsEqual(check1, check2) {
			for j := range data[step].AnswerLevelSet {
				for k := range data[step].AnswerLevelSet[j].Set {
					data[step].AnswerLevelSet[j].Set[k].Flag = true
				}
			}
		} else {
			incorrect++
		}
		all = all + 3
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[step], nil
}

func (s *lab3bService) CheckLab3BResult(ctx context.Context, userId int, userIndex int, userSet []float64) (int, int, model.DataResponse, []model.DataResponse, error) {
	maxMark := 10
	maxMark1 := 5
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3BId)
	if err != nil {
		return 0, maxMark1 + maxMark, model.DataResponse{}, nil, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark1 + maxMark, model.DataResponse{}, nil, err
	}

	critVal, alterVal, _, err := s.parseRuleResult(ctx, matrix, variance.Criteria, variance.Rule)
	if err != nil {
		return 0, maxMark1 + maxMark, model.DataResponse{}, nil, err
	}

	_, points := s.GetPointsByLukasiewiczImplication(ctx, critVal, alterVal)

	num := s.GetPointEstimateByAlternativeMatrix(ctx, points)
	var maxVal float64 = 0
	currentIndex := 0
	var data []model.DataResponse
	for i := range num {
		data = append(data, model.DataResponse{
			Data: num[i],
			Flag: false,
		})
		if num[i] > maxVal {
			maxVal = num[i]
			currentIndex = i
		}
	}
	index := model.DataResponse{
		Data: currentIndex,
		Flag: false,
	}
	if len(num) != len(userSet) {
		return 0, maxMark1 + maxMark, index, data, nil
	}

	for i := range data {
		if data[i].Data != userSet[i] {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	mark := s.getMark(float64(maxMark), incorrect, all)
	if userIndex == currentIndex {
		mark = mark + maxMark1
		index.Flag = true
	}

	return mark, maxMark1 + maxMark, index, data, nil
}

func (s *lab3bService) arrayIsEqual(arr1, arr2 []float64) bool {
	sort.Float64s(arr1)
	sort.Float64s(arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func (s *lab3bService) validatePoints(ctx context.Context, points [][]model.Point) error {
	for i := range points {
		flag := false
		if len(points[i]) < 2 {
			flag = true
			continue
		}
		for j := 0; j < len(points[i])-1; j++ {
			if points[i][j].X != points[i][j+1].X {
				flag = true
			}
		}

		if !flag {
			return fmt.Errorf("line")
		}
	}

	return nil
}

func (s *lab3bService) createMatrixByCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	result, err := s.getResultOfCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *lab3bService) parseRuleResult(ctx context.Context, matrix [][]float64, criteria []model.Criteria, rules []model.Rule) ([][]float64, []string, [][]map[string]float64, error) {
	var alternatives []string
	formRules := make([][]float64, len(rules))

	for _, rule := range rules {
		str := strings.Split(rule.Name, "Y")
		if len(str) < 2 {
			return nil, nil, nil, fmt.Errorf("can't parse rule result: %s", rule.Name)
		}
		result := strings.ReplaceAll(str[1], "=", "")
		result = strings.ReplaceAll(result, " ", "")
		alternatives = append(alternatives, result)
	}

	var rulesData [][]map[string]float64
	for i, rule := range rules {
		var resultData []map[string]float64
		var minRes []float64
		for i := range matrix {
			isDone := true
			strToAdd := rule
			data := make(map[string]float64)
			for isDone {
				res := ""
				index := strings.Index(strToAdd.Func, "$")
				if index != -1 {
					res = strToAdd.Func[index : index+3]
					if res != "" && len(res) == 3 {
						v, vFloat, err := s.getValueOfRule(ctx, matrix[i], criteria, res)
						if err != nil {
							return nil, nil, nil, err
						}
						strToAdd.Func = strings.ReplaceAll(strToAdd.Func, res, v)
						data[string("c"+res[1:2])] = vFloat
					}
					continue
				} else {
					ruleRes, err := s.prepareRule(ctx, strToAdd.Func)
					if err != nil {
						return nil, nil, nil, err
					}

					minRes = append(minRes, ruleRes)
					isDone = false
					continue
				}
			}
			resultData = append(resultData, data)
		}
		rulesData = append(rulesData, resultData)
		formRules[i] = minRes
	}

	return formRules, alternatives, rulesData, nil
}

func (s *lab3bService) getValueOfRule(ctx context.Context, alternative []float64, criteria []model.Criteria, res string) (string, float64, error) {
	indexCriteria, err := strconv.Atoi(string(res[1]))
	if err != nil {
		return "", 0, err
	}
	indexCriteria--

	if indexCriteria >= len(criteria) || indexCriteria < 0 {
		return "", 0, fmt.Errorf("can't find criteria with index %d", indexCriteria)
	}

	var funcMarkRes string
	for _, v := range criteria[indexCriteria].FuncMark {
		if v.RealName == string(res[2]) {
			funcMarkRes = v.Func
			break
		}
	}

	if funcMarkRes == "" {
		return "", 0, fmt.Errorf("can't find func mark with name %s", string(res[2]))
	}

	resValue, err := s.eval(funcMarkRes, math.Round(alternative[indexCriteria]*100)/100)
	if err != nil {
		return "", 0, err
	}

	return strconv.FormatFloat(math.Round(resValue*100)/100, 'f', -1, 64), math.Round(100*resValue) / 100, nil
}

func (s *lab3bService) GetPointsByLukasiewiczImplication(ctx context.Context, matrix [][]float64, ruleResult []string) ([][][]model.Point, [][]model.Point) {
	points := make([][]model.Point, len(matrix))
	for i := range points {
		points[i] = make([]model.Point, 11)
		var x float64 = 0
		for j := range points[i] {
			points[i][j] = model.Point{X: math.Round(x*100) / 100, Y: 1}
			x += 0.1
		}
	}

	var matrixPoints [][][]model.Point
	for i := range matrix {
		var dataToAdd [][]model.Point
		for j := range matrix[i] {
			var x float64 = 0
			var data []model.Point
			for x <= 1 {
				mb := s.getRuleResult(ctx, ruleResult[i], math.Round(x*10)/10)
				res := math.Min(1, math.Round((1-matrix[i][j]+mb)*100)/100)
				data = append(data, model.Point{X: math.Round(x*100) / 100, Y: math.Round(res*100) / 100})
				x += 0.1
			}
			dataToAdd = append(dataToAdd, data)
		}
		matrixPoints = append(matrixPoints, dataToAdd)
	}

	for k := range matrixPoints {
		for i := range matrixPoints[k] {
			for j := range matrixPoints[k][i] {
				points[i][j].Y = math.Min(points[i][j].Y, matrixPoints[k][i][j].Y)
			}
		}
	}

	return matrixPoints, points
}

func (s *lab3bService) prepareRule(ctx context.Context, expr string) (float64, error) {
	node, err := parser.ParseExpr(expr)
	if err != nil {
		return 0, err
	}

	result := s.evaluateExpression(ctx, node)
	return result, nil
}

func (s *lab3bService) evaluateExpression(ctx context.Context, expr ast.Expr) float64 {
	switch node := expr.(type) {
	case *ast.BinaryExpr:
		left := s.evaluateExpression(ctx, node.X)
		right := s.evaluateExpression(ctx, node.Y)
		switch node.Op {
		case token.ADD:
			return left + right
		case token.SUB:
			return left - right
		case token.MUL:
			return left * right
		case token.QUO:
			return left / right
		default:
			return 0
		}
	case *ast.CallExpr:
		funcName := node.Fun.(*ast.Ident).Name
		if funcName == "min" {
			return s.minOrMax(ctx, node, true)
		} else if funcName == "max" {
			return s.minOrMax(ctx, node, false)
		} else {
			return 0
		}
	case *ast.BasicLit:
		val, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}

func (s *lab3bService) minOrMax(ctx context.Context, call *ast.CallExpr, isMin bool) float64 {
	if len(call.Args) == 0 {
		return 0
	}

	result := s.evaluateExpression(ctx, call.Args[0])

	for _, arg := range call.Args[1:] {
		val := s.evaluateExpression(ctx, arg)
		if isMin {
			result = math.Min(result, val)
		} else {
			result = math.Max(result, val)
		}
	}

	return result
}

func (s *lab3bService) getRuleResult(ctx context.Context, quality string, x float64) float64 {
	switch quality {
	case "удовлетворительный":
		return x
	case "неудовлетворительный":
		return 1 - x
	case "безупречный":
		if x == 1 {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func (s *lab3bService) getRuleCriteria(ctx context.Context, index int, quality string, x float64) float64 {
	switch quality {
	case "большое":
		if index == 3 {
			return 1 - x
		}
		return x
	case "высокое":
		if index == 3 {
			return 1 - x
		}
		return x
	case "низкое":
		if index == 3 {
			return x
		}
		return 1 - x
	case "маленькое":
		if index == 3 {
			return x
		}
		return 1 - x
	case "очень маленькое":
		return x * x
	case "очень низкое":
		return x * x
	default:
		return 0
	}
}

func (s *lab3bService) GetPointEstimateByAlternativeMatrix(ctx context.Context, matrix [][]model.Point) []float64 {
	var result []float64

	levelSets := s.getLevelSets(ctx, matrix)
	for i := range levelSets {
		var maxVal float64 = 0
		for j := range levelSets[i][0].Set {
			if levelSets[i][0].Set[j] > maxVal {
				maxVal = levelSets[i][0].Set[j]
			}
		}
		var res float64 = 0
		for j := range levelSets[i] {
			res = res + levelSets[i][j].Delta*levelSets[i][j].Powerful
		}
		if maxVal != 0 {
			result = append(result, math.Round(res/maxVal*100)/100)
		} else {
			return result
		}
	}

	return result
}

func (s *lab3bService) getLevelSets(ctx context.Context, matrix [][]model.Point) [][]model.AnswerLevelSet {
	var result [][]model.AnswerLevelSet

	for i := range matrix {
		sliceToWork := matrix[i]
		var lastMin float64 = 0
		var currentAnswer []model.AnswerLevelSet
		for len(sliceToWork) != 0 {
			length := len(sliceToWork)
			sum := s.sumByXValue(ctx, sliceToWork)
			slice, minValue := s.getSliceWithoutMinimum(ctx, sliceToWork)
			delta := minValue - lastMin
			lastMin = minValue
			var sl []float64
			for k := range sliceToWork {
				sl = append(sl, sliceToWork[k].X)
			}
			currentAnswer = append(currentAnswer, model.AnswerLevelSet{
				Set:      sl,
				Delta:    math.Round(delta*100) / 100,
				Powerful: math.Round(sum/float64(length)*100) / 100,
			})
			sliceToWork = slice
		}
		result = append(result, currentAnswer)
	}

	return result
}

func (s *lab3bService) getSliceWithoutMinimum(ctx context.Context, oldSlice []model.Point) ([]model.Point, float64) {
	slice := make([]model.Point, len(oldSlice))
	copy(slice, oldSlice)

	minValue := slice[0].Y

	for i := range slice {
		if slice[i].Y < minValue {
			minValue = slice[i].Y
		}
	}

	var newSlice []model.Point
	for i := range slice {
		if slice[i].Y != minValue {
			newSlice = append(newSlice, slice[i])
		}
	}

	return newSlice, minValue
}

func (s *lab3bService) sumByXValue(ctx context.Context, slice []model.Point) float64 {
	var sum float64 = 0
	for i := range slice {
		sum += slice[i].X
	}

	return sum
}

func (s *lab3bService) getResultOfCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	var matrix [][]float64

	for i := range variance.Criteria {
		var criteria []float64
		for j := range variance.Alternative {
			if variance.Alternative[j].CriteriaCount[i].IsQualitative {
				var x float64
				for k := range variance.Criteria[i].Gradation {
					if k == len(variance.Criteria[i].Gradation)-1 && variance.Alternative[j].CriteriaCount[i].Count >= variance.Criteria[i].Gradation[k].StartPointX {
						x = variance.Criteria[i].Gradation[k].StartPointX
						break
					} else if variance.Criteria[i].Gradation[k].StartPointX == variance.Criteria[i].Gradation[k].EndPointX && variance.Criteria[i].Gradation[k].StartPointX == variance.Alternative[j].CriteriaCount[i].Count {
						x = variance.Criteria[i].Gradation[k].StartPointX
						break
					} else if variance.Alternative[j].CriteriaCount[i].Count < variance.Criteria[i].Gradation[k].EndPointX && variance.Alternative[j].CriteriaCount[i].Count >= variance.Criteria[i].Gradation[k].StartPointX {
						x = variance.Criteria[i].Gradation[k].StartPointX
						break
					}
				}
				res, err := s.eval(variance.Criteria[i].RightFunc, x)
				if err != nil {
					return nil, err
				}

				criteria = append(criteria, math.Round(res*100)/100)
			} else {
				res, err := s.eval(variance.Criteria[i].RightFunc, variance.Alternative[j].CriteriaCount[i].Count)
				if err != nil {
					return nil, err
				}

				criteria = append(criteria, math.Round(res*100)/100)
			}
		}
		matrix = append(matrix, criteria)
	}

	return s.transposeMatrix(ctx, matrix), nil
}

func (s *lab3bService) transposeMatrix(ctx context.Context, matrix [][]float64) [][]float64 {
	numRows := len(matrix)
	numCols := len(matrix[0])

	transposed := make([][]float64, numCols)
	for i := range transposed {
		transposed[i] = make([]float64, numRows)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
