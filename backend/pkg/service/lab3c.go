package service

import (
	"context"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"backend/pkg/model"
	"backend/pkg/repository"
)

const Lab3CId = 6

type lab3cService struct {
	repo *repository.Repo
	commonEventService
}

func NewLab3cService(repo *repository.Repo) *lab3cService {
	return &lab3cService{
		repo:               repo,
		commonEventService: NewCommonEventService(),
	}
}

func (s *lab3cService) ValidateLab3CResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error) {
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] < 0 || mat[i][j] > 1 {
				return nil, fmt.Errorf("matrix coefficients are not valid")
			}
		}
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	for i := range ch {
		var lastElem float64 = -1
		for j := range ch[i] {
			if lastElem == ch[i][j].X {
				return nil, fmt.Errorf("not valid variant")
			}
			lastElem = ch[i][j].X
		}
	}

	mid := s.GetMiddleOfArea(ctx, ch)

	return mid, nil
}

func (s *lab3cService) CheckLab3CMiddleOfArea(ctx context.Context, userId int, userIndex int, userMid []float64) (int, int, model.DataResponse, []model.DataResponse, error) {
	maxMark1 := 10
	maxMark2 := 5
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark1 + maxMark2, model.DataResponse{}, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark1 + maxMark2, model.DataResponse{}, nil, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)
	mid := s.GetMiddleOfArea(ctx, ch)

	currentIndex := 0
	var maxMid float64 = 0
	for i := range mid {
		if mid[i] >= maxMid {
			maxMid = mid[i]
			currentIndex = i
		}
	}

	var data []model.DataResponse
	for i := range mid {
		data = append(data, model.DataResponse{
			Data: mid[i],
			Flag: false,
		})
	}
	index := model.DataResponse{
		Data: currentIndex,
		Flag: false,
	}

	if len(mid) != len(userMid) {
		return 0, maxMark1 + maxMark2, index, data, nil
	}

	for i := range data {
		if data[i].Data != userMid[i] {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	mark := s.getMark(float64(maxMark1), incorrect, all)
	if userIndex == currentIndex {
		index.Flag = true
		mark = mark + maxMark2
	}

	return mark, maxMark1 + maxMark2, index, data, nil
}

func (s *lab3cService) CheckLab3CQuadraticParameters(ctx context.Context, userId int, userData model.AnswerLab3CQuadraticParametersRequest) (int, int, model.AnswerLab3CQuadraticParameters, error) {
	maxMark := 3
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, model.AnswerLab3CQuadraticParameters{}, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, model.AnswerLab3CQuadraticParameters{}, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var data []model.QuadraticParameters
	for _, c := range ch {
		_, param := s.getCoefficientsOfQuadraticEquation(ctx, c)
		data = append(data, param)
	}

	if userData.Step >= len(data) {
		return 0, maxMark, model.AnswerLab3CQuadraticParameters{}, fmt.Errorf("step %d is out of range", userData.Step)
	}

	if data[userData.Step].A1.Data != userData.Parameters.A1 {
		incorrect++
	} else {
		data[userData.Step].A1.Flag = true
	}

	if data[userData.Step].A2.Data != userData.Parameters.A2 {
		incorrect++
	} else {
		data[userData.Step].A2.Flag = true
	}

	if data[userData.Step].A3.Data != userData.Parameters.A3 {
		incorrect++
	} else {
		data[userData.Step].A3.Flag = true
	}
	all = all + 3

	return s.getMark(float64(maxMark), incorrect, all), maxMark, model.AnswerLab3CQuadraticParameters{Parameters: data[userData.Step]}, nil
}

func (s *lab3cService) CheckLab3CLineParameters(ctx context.Context, userId int, userData model.AnswerLab3CLineParametersRequest) (int, int, model.AnswerLab3CLineParameters, error) {
	maxMark := 3
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, model.AnswerLab3CLineParameters{}, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, model.AnswerLab3CLineParameters{}, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var data []model.LineParameters
	for _, c := range ch {
		param, _ := s.getCoefficientsOfQuadraticEquation(ctx, c)
		data = append(data, param)
	}

	if userData.Step >= len(data) {
		return 0, maxMark, model.AnswerLab3CLineParameters{}, fmt.Errorf("step %d is out of range", userData.Step)
	}

	if data[userData.Step].K.Data != userData.Parameters.K {
		incorrect++
	} else {
		data[userData.Step].K.Flag = true
	}
	if data[userData.Step].B.Data != userData.Parameters.B {
		incorrect++
	} else {
		data[userData.Step].B.Flag = true
	}
	all = all + 2

	return s.getMark(float64(maxMark), incorrect, all), maxMark, model.AnswerLab3CLineParameters{Parameters: data[userData.Step]}, nil
}

func (s *lab3cService) CheckLab3CArea(ctx context.Context, userId int, userArea model.AnswerLab3CArea) (int, int, model.DataResponse, error) {
	maxMark := 3
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, model.DataResponse{}, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, model.DataResponse{}, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var area []model.DataResponse
	for i := range ch {
		area = append(area, model.DataResponse{
			Data: math.Round((ch[i][2].X-ch[i][0].X)/2*100) / 100,
			Flag: false,
		})
	}

	if userArea.Step >= len(area) {
		return 0, maxMark, model.DataResponse{}, fmt.Errorf("step %d is out of range", userArea.Step)
	}

	if userArea.Set != area[userArea.Step].Data {
		incorrect++
	} else {
		area[userArea.Step].Flag = true
	}
	all++

	return s.getMark(float64(maxMark), incorrect, all), maxMark, area[userArea.Step], nil
}

func (s *lab3cService) CheckLab3CCriteriaMatrix(ctx context.Context, userId int, userMatrix [][]model.Point) (int, int, [][]model.PointCheck, error) {
	maxMark := 15
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, matrix, _ := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var data [][]model.PointCheck
	for i := range matrix {
		var dataToAdd []model.PointCheck
		for j := range matrix[i] {
			dataToAdd = append(dataToAdd, model.PointCheck{
				X:    matrix[i][j].X,
				Y:    matrix[i][j].Y,
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
			if data[i][j].X != userMatrix[i][j].X || data[i][j].Y != userMatrix[i][j].Y {
				incorrect++
			} else {
				data[i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3cService) CheckLab3CAlternativeMatrices(ctx context.Context, userId int, step int, userMatrices [][]model.Point) (int, int, [][]model.PointCheck, error) {
	maxMark := 20
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	matrices, _, _ := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var data [][][]model.PointCheck
	for i := range matrices {
		var dataToAdd2 [][]model.PointCheck
		for j := range matrices[i] {
			var dataToAdd1 []model.PointCheck
			for n := range matrices[i][j] {
				dataToAdd1 = append(dataToAdd1, model.PointCheck{
					X:    matrices[i][j][n].X,
					Y:    matrices[i][j][n].Y,
					Flag: false,
				})
			}
			dataToAdd2 = append(dataToAdd2, dataToAdd1)
		}
		data = append(data, dataToAdd2)
	}

	if step >= len(data) {
		return 0, maxMark, nil, fmt.Errorf("step: %d out of range", step)
	}

	if len(data[step]) != len(userMatrices) {
		return 0, maxMark, data[step], nil
	}

	for i := range data[step] {
		if len(data[step][i]) != len(userMatrices[i]) {
			return 0, maxMark, data[step], nil
		}
		for j := range data[step][i] {
			if data[step][i][j].X != userMatrices[i][j].X || data[step][i][j].Y != userMatrices[i][j].Y {
				incorrect++
			} else {
				data[step][i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[step], nil
}

func (s *lab3cService) CheckLab3CCurrentMatrix(ctx context.Context, userId int, userEstimation model.AnswerLab3CCurrentMatrix) (int, int, [][]model.PointCheck, error) {
	maxMark := 5
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	data := s.getResultOfMultiplyAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	if userEstimation.Step >= len(data) {
		return 0, maxMark, nil, fmt.Errorf("step: %d out of range", userEstimation.Step)
	}

	for i := range data[userEstimation.Step] {
		for j := range data[userEstimation.Step][i] {
			if data[userEstimation.Step][i][j].X != userEstimation.Matrix[i][j].X || data[userEstimation.Step][i][j].Y != userEstimation.Matrix[i][j].Y {
				incorrect++
			} else {
				data[userEstimation.Step][i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[userEstimation.Step], nil
}

func (s *lab3cService) CheckLab3CEstimation(ctx context.Context, userId int, userEstimation model.AnswerLab3CEstimation) (int, int, []model.PointCheck, error) {
	maxMark := 5
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	alternativeImportance, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	_, _, ch := s.GetResultAlternativesData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, alternativeImportance, criteriaImportance)

	var data [][]model.PointCheck
	for i := range ch {
		var dataToAdd []model.PointCheck
		for j := range ch[i] {
			dataToAdd = append(dataToAdd, model.PointCheck{
				X:    ch[i][j].X,
				Y:    ch[i][j].Y,
				Flag: false,
			})
		}
		data = append(data, dataToAdd)
	}

	if userEstimation.Step >= len(data) {
		return 0, maxMark, nil, fmt.Errorf("step: %d out of range", userEstimation.Step)
	}

	for i := range userEstimation.Matrix {
		if data[userEstimation.Step][i].X != userEstimation.Matrix[i].X || data[userEstimation.Step][i].Y != userEstimation.Matrix[i].Y {
			incorrect++
		} else {
			data[userEstimation.Step][i].Flag = true
		}

		all++
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[userEstimation.Step], nil
}

func (s *lab3cService) CheckLab3CAlternativesImportance(ctx context.Context, userId int, step int, userAlterMatrix []string) (int, int, []model.DataResponse, error) {
	maxMark1 := 7
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark1, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark1, nil, err
	}

	alternativeImportance, _ := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)
	for i := range alternativeImportance {
		for j := range alternativeImportance[i] {
			alternativeImportance[i][j] = s.extractFirstLetters(alternativeImportance[i][j])
		}
	}

	if step >= len(alternativeImportance) {
		return 0, maxMark1, nil, fmt.Errorf("step: %d out of range", step)
	}

	var data []model.DataResponse
	for i := range alternativeImportance[step] {
		data = append(data, model.DataResponse{
			Data: alternativeImportance[step][i],
			Flag: false,
		})
	}

	for i := range data {
		if data[i].Data != userAlterMatrix[i] {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	return s.getMark(float64(maxMark1), incorrect, all), maxMark1, data, nil
}

func (s *lab3cService) CheckLab3CCriteriaImportance(ctx context.Context, userId int, userCriteriaMatrix []string) (int, int, []model.DataResponse, error) {
	maxMark := 7
	all := 0
	incorrect := 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3CId)
	if err != nil {
		return 0, maxMark, nil, err
	}
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	mat, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}

	_, criteriaImportance := s.GetImportanceData(ctx, variance.ImportanceCriteriaMatrix, variance.ImportanceAlternativeMatrix, mat, coefficients)

	var data []model.DataResponse
	for i := range criteriaImportance {
		data = append(data, model.DataResponse{
			Data: s.extractFirstLetters(criteriaImportance[i]),
			Flag: false,
		})
	}

	if len(userCriteriaMatrix) != len(data) {
		return 0, maxMark, data, nil
	}

	all = 0
	incorrect = 0
	for i := range data {
		if userCriteriaMatrix[i] != data[i].Data {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3cService) extractFirstLetters(input string) string {
	words := strings.Fields(input)
	var result strings.Builder

	for _, word := range words {
		r, _ := utf8.DecodeRuneInString(word)
		result.WriteRune(r)
	}

	return result.String()
}

func (s *lab3cService) GetImportanceData(ctx context.Context, criteriaData, alternativeData []model.ImportancePoints, criteria [][]float64, criteriaWeight []float64) ([][]string, []string) {
	result := make([][]string, len(criteria))
	criteriaResultImportance := make([]string, len(criteriaWeight))
	for i := 0; i < len(criteria); i++ {
		criteriaResult := make([]string, len(criteria[i]))
		for j := 0; j < len(criteria[i]); j++ {
			var maxCount float64 = 0
			importance := ""
			for _, v := range alternativeData {
				if s.getYValueOnLine(ctx, v.Points, model.Point{X: criteria[i][j]}) > maxCount {
					maxCount = s.getYValueOnLine(ctx, v.Points, model.Point{X: criteria[i][j]})
					importance = v.Importance
				}
			}
			if maxCount == 0 {
				var minDist float64 = 1
				for _, v := range alternativeData {
					if s.getXMinDistance(ctx, v.Points, model.Point{X: criteria[i][j]}) <= minDist {
						minDist = s.getXMinDistance(ctx, v.Points, model.Point{X: criteria[i][j]})
						importance = v.Importance
					}
				}
			}

			criteriaResult[j] = importance
		}

		result[i] = criteriaResult
	}

	for i := 0; i < len(criteriaWeight); i++ {
		var criteriaMaxCount float64 = 0
		criteriaImportance := ""
		for j := 0; j < len(criteriaData); j++ {
			if s.getYValueOnLine(ctx, criteriaData[j].Points, model.Point{X: criteriaWeight[i]}) >= criteriaMaxCount {
				criteriaMaxCount = s.getYValueOnLine(ctx, criteriaData[j].Points, model.Point{X: criteriaWeight[i]})
				criteriaImportance = criteriaData[j].Importance
			}
		}
		if criteriaMaxCount == 0 {
			minDistance := s.getXMinDistance(ctx, criteriaData[0].Points, model.Point{X: criteriaWeight[i]})
			criteriaImportance = criteriaData[0].Importance
			for j := 0; j < len(criteriaData); j++ {
				if s.getXMinDistance(ctx, criteriaData[j].Points, model.Point{X: criteriaWeight[i]}) <= minDistance {
					minDistance = s.getXMinDistance(ctx, criteriaData[j].Points, model.Point{X: criteriaWeight[i]})
					criteriaImportance = criteriaData[j].Importance
				}
			}
		}

		criteriaResultImportance[i] = criteriaImportance
	}

	return result, criteriaResultImportance
}

func (s *lab3cService) getResultOfMultiplyAlternativesData(ctx context.Context, criteriaData, alternativeData []model.ImportancePoints, alternativeImportance [][]string, criteriaImportance []string) [][][]model.PointCheck {
	data := make([][][]model.Point, len(alternativeImportance))
	coffData := make([][]model.Point, len(criteriaImportance))

	for i := range criteriaImportance {
		for _, v := range criteriaData {
			if criteriaImportance[i] == v.Importance {
				coffData[i] = v.Points
			}
		}
	}

	for i := range alternativeImportance {
		var dataToAdd [][]model.Point
		for j := range alternativeImportance[i] {
			for _, v := range alternativeData {
				if alternativeImportance[i][j] == v.Importance {
					dataToAdd = append(dataToAdd, v.Points)
				}
			}
		}
		data[i] = dataToAdd
	}

	var result [][][]model.PointCheck
	for i := range data {
		var dataToAdd2 [][]model.PointCheck
		for j := range data[i] {
			var dataToAdd1 []model.PointCheck
			for k := range data[i][j] {
				dataToAdd1 = append(dataToAdd1, model.PointCheck{
					X:    math.Round(data[i][j][k].X*coffData[j][k].X*100) / 100,
					Y:    math.Round(coffData[j][k].Y*100) / 100,
					Flag: false,
				})
			}
			dataToAdd2 = append(dataToAdd2, dataToAdd1)
		}
		result = append(result, dataToAdd2)
	}

	return result
}

func (s *lab3cService) GetResultAlternativesData(ctx context.Context, criteriaData, alternativeData []model.ImportancePoints, alternativeImportance [][]string, criteriaImportance []string) ([][][]model.Point, [][]model.Point, [][]model.Point) {
	data := make([][][]model.Point, len(alternativeImportance))
	coffData := make([][]model.Point, len(criteriaImportance))

	for i := range criteriaImportance {
		for _, v := range criteriaData {
			if criteriaImportance[i] == v.Importance {
				coffData[i] = v.Points
			}
		}
	}

	for i := range alternativeImportance {
		var dataToAdd [][]model.Point
		for j := range alternativeImportance[i] {
			for _, v := range alternativeData {
				if alternativeImportance[i][j] == v.Importance {
					dataToAdd = append(dataToAdd, v.Points)
				}
			}
		}
		data[i] = dataToAdd
	}

	var pointsSum [][]model.Point
	for _, result := range data {
		points := make([]model.Point, len(result[0]))
		for i := range result {
			for j := range result[i] {
				points[j].X = math.Round(math.Round(points[j].X*100)+math.Round(result[i][j].X*coffData[i][j].X*100)) / 100
				points[j].Y = result[i][j].Y
			}
		}
		pointsSum = append(pointsSum, points)
	}

	for i := range pointsSum {
		for j := range pointsSum[i] {
			pointsSum[i][j].X = math.Round(pointsSum[i][j].X*100) / 100
			pointsSum[i][j].Y = math.Round(pointsSum[i][j].Y*100) / 100
		}
	}

	return data, coffData, pointsSum
}

func (s *lab3cService) GetMiddleOfArea(ctx context.Context, points [][]model.Point) []float64 {
	var result []float64
	for i := range points {
		answer := math.Round(s.getAnswerOfQuadraticEquation(ctx, points[i])*100) / 100
		result = append(result, answer)
	}

	return result
}

func (s *lab3cService) getYValueOnLine(ctx context.Context, points []model.Point, targetPoint model.Point) float64 {
	var maxValue float64 = 0
	for i := 0; i < len(points)-1; i++ {
		if points[i].X == points[i+1].X && points[i].X == targetPoint.X {
			maxValue = 1
		}
		if (targetPoint.X > points[i].X && targetPoint.X > points[i+1].X) || (targetPoint.X < points[i].X && targetPoint.X < points[i+1].X) {
			continue
		}
		k := (points[i+1].Y - points[i].Y) / (points[i+1].X - points[i].X)
		b := points[i].Y - k*points[i].X
		result := k*targetPoint.X + b
		if result > maxValue {
			maxValue = math.Round(result*100) / 100
		}
	}

	return maxValue
}

func (s *lab3cService) getXMinDistance(ctx context.Context, points []model.Point, targetPoint model.Point) float64 {
	minValue1 := points[0].X - targetPoint.X
	minValue2 := points[2].X - targetPoint.X

	return math.Min(minValue1, minValue2)
}

func (s *lab3cService) getAnswerOfQuadraticEquation(ctx context.Context, points []model.Point) float64 {
	areaTriangle1 := (points[1].X - points[0].X) * points[1].Y / 2
	areaTriangle2 := (points[2].X - points[1].X) * points[2].Y / 2

	var k, b, a1, a2, a3 float64
	if areaTriangle1 > areaTriangle2 {
		k = math.Round(100*((points[2].Y-points[1].Y)/(points[2].X-points[1].X))) / 100
		b = math.Round(100*(points[1].Y-points[1].X*(points[2].Y-points[1].Y)/(points[2].X-points[1].X))) / 100
		a1 = math.Round(100*k) / 100
		a2 = math.Round(100*(b-k*points[2].X)) / 100
		a3 = math.Round(100*((points[2].X-points[0].X)/2-b*points[2].X)) / 100
	} else if areaTriangle1 < areaTriangle2 {
		k = math.Round(100*((points[1].Y-points[0].Y)/(points[1].X-points[0].X))) / 100
		b = math.Round(100*(points[1].Y-points[0].X*(points[1].Y-points[0].Y)/(points[1].X-points[0].X))) / 100
		a1 = math.Round(100*k) / 100
		a2 = math.Round(100*(b-k*points[1].X)) / 100
		a3 = math.Round(100*((points[2].X-points[0].X)/2-b*points[1].X)) / 100
	} else {
		return math.Round(100*points[1].X) / 100
	}

	desc := a2*a2 - 4*a1*a3
	if desc < 0 {
		return -1
	}

	x1 := (-a2 + math.Sqrt(desc)) / (2 * a1)
	x2 := (-a2 - math.Sqrt(desc)) / (2 * a1)
	if (x1 > points[0].X && x1 > points[1].X && x1 > points[2].X) || (x1 < points[0].X && x1 < points[1].X && x1 < points[2].X) {
		x1 = -1
	}
	if (x2 > points[0].X && x2 > points[1].X && x2 > points[2].X) || (x2 < points[0].X && x2 < points[1].X && x2 < points[2].X) {
		x2 = -1
	}

	if x1 == -1 {
		return x2
	}
	if x2 == -1 {
		return x1
	}

	return -1
}

func (s *lab3cService) getCoefficientsOfQuadraticEquation(ctx context.Context, points []model.Point) (model.LineParameters, model.QuadraticParameters) {
	areaTriangle1 := (points[1].X - points[0].X) * points[1].Y / 2
	areaTriangle2 := (points[2].X - points[1].X) * points[2].Y / 2

	var k, b, a1, a2, a3 float64 = 0, 0, 0, 0, 0

	if areaTriangle1 > areaTriangle2 {
		k = math.Round(100*((points[2].Y-points[1].Y)/(points[2].X-points[1].X))) / 100
		b = math.Round(100*(points[1].Y-points[1].X*(points[2].Y-points[1].Y)/(points[2].X-points[1].X))) / 100
		a1 = math.Round(100*k) / 100
		a2 = math.Round(100*(b-k*points[2].X)) / 100
		a3 = math.Round(100*((points[2].X-points[0].X)/2-b*points[2].X)) / 100
	} else if areaTriangle1 < areaTriangle2 {
		k = math.Round(100*((points[1].Y-points[0].Y)/(points[1].X-points[0].X))) / 100
		b = math.Round(100*(points[1].Y-points[0].X*(points[1].Y-points[0].Y)/(points[1].X-points[0].X))) / 100
		a1 = math.Round(100*k) / 100
		a2 = math.Round(100*(b-k*points[1].X)) / 100
		a3 = math.Round(100*((points[2].X-points[0].X)/2-b*points[1].X)) / 100
	}

	line := model.LineParameters{
		K: model.DataResponse{Data: k, Flag: false},
		B: model.DataResponse{Data: b, Flag: false},
	}

	quadratic := model.QuadraticParameters{
		A1: model.DataResponse{Data: a1, Flag: false},
		A2: model.DataResponse{Data: a2, Flag: false},
		A3: model.DataResponse{Data: a3, Flag: false},
	}

	return line, quadratic
}

func (s *lab3cService) createMatrixByCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	result, err := s.getResultOfCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *lab3cService) getResultOfCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
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

func (s *lab3cService) transposeMatrix(ctx context.Context, matrix [][]float64) [][]float64 {
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
