package service

import (
	"context"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"

	"backend/pkg/model"
	"backend/pkg/repository"
)

const Lab3AId = 4

type lab3aService struct {
	repo *repository.Repo
	commonEventService
}

func NewLab3aService(repo *repository.Repo) *lab3aService {
	return &lab3aService{
		repo:               repo,
		commonEventService: NewCommonEventService(),
	}
}

func (s *lab3aService) ValidateLab3AResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error) {
	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	var maxValue []float64
	for i := range matrix {
		var maxCount float64 = 0
		for j := range matrix[i] {
			if matrix[i][j] >= maxCount {
				maxCount = matrix[i][j]
			}
			if matrix[i][j] < 0 || matrix[i][j] > 1 {
				return nil, fmt.Errorf("matrix coefficients are not valid")
			}
		}
		maxValue = append(maxValue, maxCount)
	}

	for j := range matrix[0] {
		flag := true
		for i := range matrix {
			if matrix[i][j] != maxValue[j] {
				flag = false
			}
		}
		if flag {
			return nil, fmt.Errorf("there is obvious answer")
		}
	}

	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	if err := s.validateMatrix(diffMat); err != nil {
		return nil, err
	}

	inter := s.GetAlternativesMatricesIntersection(ctx, diffMat)

	var sets [][]float64
	sets = append(sets, s.GetSetNonDominatedAlternatives(ctx, inter))
	mat := s.GetAlternativesMatricesWithCoefficients(ctx, diffMat, coefficients)
	sets = append(sets, s.GetSetNonDominatedAlternatives(ctx, mat))

	res := s.GetSetNonDominatedAlternativesOnSet(ctx, sets)
	return res, nil
}

func (s *lab3aService) validateMatrix(matrices []*mat.Dense) error {
	for _, matrix := range matrices {
		numRows, numCols := matrix.Dims()
		for i := 0; i < numRows; i++ {
			for j := 0; j < numCols; j++ {
				if matrix.At(i, j) != 1 && matrix.At(i, j) != 0 {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("matrix does not contain 1 or 2 rows")
}

func (s *lab3aService) CheckLab3AResult(ctx context.Context, userId int, userRes []float64, chosenIndex int) (int, int, []model.DataResponse, model.DataResponse, error) {
	maxMark1 := 10
	maxMark2 := 5
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
	if err != nil {
		return 0, maxMark1 + maxMark2, nil, model.DataResponse{
			Data: 0,
			Flag: false,
		}, err
	}

	var coefficients []float64
	for _, v := range variance.Criteria {
		coefficients = append(coefficients, v.Weight)
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark1 + maxMark2, nil, model.DataResponse{Flag: false}, err
	}
	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	inter := s.GetAlternativesMatricesIntersection(ctx, diffMat)

	var sets [][]float64
	sets = append(sets, s.GetSetNonDominatedAlternatives(ctx, inter))
	mat := s.GetAlternativesMatricesWithCoefficients(ctx, diffMat, coefficients)
	sets = append(sets, s.GetSetNonDominatedAlternatives(ctx, mat))

	res := s.GetSetNonDominatedAlternativesOnSet(ctx, sets)
	var data []model.DataResponse
	maxCount := res[0]
	valIndex := 0
	for i := range res {
		data = append(data, model.DataResponse{
			Data: res[i],
			Flag: false,
		})

		if res[i] > maxCount {
			maxCount = res[i]
			valIndex = i
		}
	}

	if len(res) != len(userRes) {
		return 0, maxMark1 + maxMark2, data, model.DataResponse{Flag: false}, nil
	}

	for i := range data {
		if data[i].Data != userRes[i] {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	mark := s.getMark(float64(maxMark1), incorrect, all)
	index := model.DataResponse{
		Data: valIndex,
		Flag: false,
	}
	if valIndex == chosenIndex {
		index.Flag = true
		mark = mark + maxMark2
	}

	return mark, maxMark1 + maxMark2, data, index, nil
}

func (s *lab3aService) CheckLab3ASecondNonDominatedSets(ctx context.Context, userId int, userSet []float64) (int, int, []model.DataResponse, error) {
	maxMark := 10
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
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
	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	mat := s.GetAlternativesMatricesWithCoefficients(ctx, diffMat, coefficients)
	set := s.GetSetNonDominatedAlternatives(ctx, mat)

	var data []model.DataResponse
	for i := range set {
		data = append(data, model.DataResponse{
			Data: set[i],
			Flag: false,
		})
	}

	if len(data) != len(userSet) {
		return 0, maxMark, data, nil
	}
	for i := range data {
		if data[i].Data != userSet[i] {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3aService) CheckLab3AAlternativesMatricesWithCoefficients(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error) {
	maxMark := 25
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
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
	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	mat := s.GetAlternativesMatricesWithCoefficients(ctx, diffMat, coefficients)

	rows, cols := mat.Dims()
	data := make([][]model.DataResponse, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]model.DataResponse, cols)
		for j := 0; j < cols; j++ {
			data[i][j].Data = mat.At(i, j)
			data[i][j].Flag = false
		}
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

func (s *lab3aService) CheckLab3ANonDominatedSet(ctx context.Context, userId int, userSets []float64) (int, int, []model.DataResponse, error) {
	maxMark := 15
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}
	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	inter := s.GetAlternativesMatricesIntersection(ctx, diffMat)
	set := s.GetSetNonDominatedAlternatives(ctx, inter)

	var data []model.DataResponse
	for i := 0; i < len(set); i++ {
		data = append(data, model.DataResponse{
			Data: set[i],
			Flag: false,
		})
	}

	if len(set) != len(userSets) {
		return 0, maxMark, data, nil
	}

	for i := range userSets {
		if userSets[i] != data[i].Data {
			incorrect++
		} else {
			data[i].Flag = true
		}
		all++
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3aService) CheckLab3AMatricesIntersection(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error) {
	maxMark := 5
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
	if err != nil {
		return 0, maxMark, nil, err
	}

	matrix, err := s.createMatrixByCriteria(ctx, variance)
	if err != nil {
		return 0, maxMark, nil, err
	}
	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	inter := s.GetAlternativesMatricesIntersection(ctx, diffMat)

	rows, cols := inter.Dims()
	data := make([][]model.DataResponse, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]model.DataResponse, cols)
		for j := 0; j < cols; j++ {
			data[i][j].Data = inter.At(i, j)
			data[i][j].Flag = false
		}
	}

	if len(data) != len(userMatrix) {
		return 0, maxMark, data, nil
	}

	for i := range userMatrix {
		if len(data[i]) != len(userMatrix[i]) {
			return 0, maxMark, data, nil
		}
		for j := range userMatrix[i] {
			if userMatrix[i][j] != data[i][j].Data {
				incorrect++
			} else {
				data[i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3aService) CheckLab3AAlternativeDiffMatrices(ctx context.Context, userId int, step int, userMatrices [][]float64) (int, int, [][]model.DataResponse, error) {
	maxMark := 5
	all, incorrect := 0, 0

	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
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

	diffMat := s.GetAlternativeDiffMatrices(ctx, len(matrix[0]), matrix)
	var result [][][]float64
	for _, m := range diffMat {
		rows, cols := m.Dims()
		data := make([][]float64, rows)
		for i := 0; i < rows; i++ {
			data[i] = make([]float64, cols)
			for j := 0; j < cols; j++ {
				data[i][j] = m.At(i, j)
			}
		}
		result = append(result, data)
	}

	var data [][][]model.DataResponse
	for i := 0; i < len(result); i++ {
		var dataAdd2 [][]model.DataResponse
		for j := 0; j < len(result[i]); j++ {
			var dataAdd1 []model.DataResponse
			for k := 0; k < len(result[i][j]); k++ {
				dataAdd1 = append(dataAdd1, model.DataResponse{
					Data: result[i][j][k],
					Flag: false,
				})
			}
			dataAdd2 = append(dataAdd2, dataAdd1)
		}
		data = append(data, dataAdd2)
	}

	if step >= len(data) {
		return 0, maxMark, nil, fmt.Errorf("step:%d out od range", step)
	}

	if len(data[step]) != len(userMatrices) {
		return 0, maxMark, data[step], nil
	}

	for i := range userMatrices {
		for j := range userMatrices[i] {
			if data[step][i][j].Data != userMatrices[i][j] {
				incorrect++
			} else {
				data[step][i][j].Flag = true
			}
			all++
		}
	}

	var mark float64
	if all == 0 {
		mark = 0
	} else {
		mark = float64(maxMark) * (1 - 2*float64(incorrect)/float64(all))
	}
	if mark < 0 {
		return 0, maxMark, data[step], nil
	}
	return s.getMark(float64(maxMark), incorrect, all), maxMark, data[step], nil
}

func (s *lab3aService) CheckLab3AAlternativeSets(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error) {
	maxMark := 10
	variance, err := s.repo.GetLab3Variance(userId, Lab3AId)
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

	var data [][]model.DataResponse
	for i := 0; i < len(matrix); i++ {
		var appendData []model.DataResponse
		for j := 0; j < len(matrix[i]); j++ {
			dataToAdd := model.DataResponse{
				Data: matrix[i][j],
				Flag: false,
			}
			appendData = append(appendData, dataToAdd)
		}
		data = append(data, appendData)
	}

	if len(userMatrix) != len(matrix) {
		return 0, maxMark, data, nil
	}

	all, incorrect := 0, 0
	for i := range userMatrix {
		if len(userMatrix[i]) != len(data[i]) {
			return 0, maxMark, data, nil
		}
		for j := range userMatrix[i] {
			if userMatrix[i][j] != data[i][j].Data {
				incorrect++
			} else {
				data[i][j].Flag = true
			}
			all++
		}
	}

	return s.getMark(float64(maxMark), incorrect, all), maxMark, data, nil
}

func (s *lab3aService) createMatrixByCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
	result, err := s.getResultOfCriteria(ctx, variance)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *lab3aService) getResultOfCriteria(ctx context.Context, variance model.GeneratedLab3Variance) ([][]float64, error) {
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

	return matrix, nil
}

func (s *lab3aService) GetAlternativeDiffMatrices(ctx context.Context, alternativesNumber int, criteria [][]float64) []*mat.Dense {
	alternativesDiffMatrix := make([]*mat.Dense, len(criteria))

	for i := range criteria {
		matrixData := make([]float64, alternativesNumber*alternativesNumber)
		for j := range matrixData {
			if j%(alternativesNumber+1) == 0 {
				matrixData[j] = 1
			} else {
				currentNumber := math.Round((criteria[i][j/alternativesNumber]-criteria[i][j%alternativesNumber])*100) / 100
				if currentNumber < 0 {
					currentNumber = 0
				}
				matrixData[j] = currentNumber
			}
		}
		alternativesDiffMatrix[i] = mat.NewDense(alternativesNumber, alternativesNumber, matrixData)
	}

	return alternativesDiffMatrix
}

func (s *lab3aService) GetAlternativesMatricesIntersection(ctx context.Context, matrices []*mat.Dense) *mat.Dense {
	numRows, numCols := matrices[0].Dims()
	resultMatrix := mat.NewDense(numRows, numCols, make([]float64, numRows*numCols))

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			minCell := matrices[0].At(i, j)
			for _, matrix := range matrices {
				if matrix.At(i, j) < minCell {
					minCell = matrix.At(i, j)
				}
			}
			resultMatrix.Set(i, j, minCell)
		}
	}

	return resultMatrix
}

func (s *lab3aService) GetSetNonDominatedAlternatives(ctx context.Context, matrix *mat.Dense) []float64 {
	numRows, numCols := matrix.Dims()
	set := make([]float64, numRows)

	for i := 0; i < numRows; i++ {
		var diff []float64
		for j := 0; j < numCols; j++ {
			if i != j {
				diff = append(diff, matrix.At(j, i)-matrix.At(i, j))
			}
		}
		maxCell := diff[0]
		for _, value := range diff {
			if value > maxCell {
				maxCell = value
			}
		}

		set[i] = 1 - maxCell
	}

	for i := range set {
		set[i] = math.Round(set[i]*100) / 100
	}

	return set
}

func (s *lab3aService) GetAlternativesMatricesWithCoefficients(ctx context.Context, matrices []*mat.Dense, coefficients []float64) *mat.Dense {
	numRows, numCols := matrices[0].Dims()
	resultMatrix := mat.NewDense(numRows, numCols, make([]float64, numRows*numCols))

	for k, matrix := range matrices {
		for i := 0; i < numRows; i++ {
			for j := 0; j < numCols; j++ {
				if i == j {
					resultMatrix.Set(i, j, 1)
				} else {
					resultMatrix.Set(i, j, resultMatrix.At(i, j)+matrix.At(i, j)*coefficients[k])
				}
			}
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			resultMatrix.Set(i, j, math.Round(resultMatrix.At(i, j)*100)/100)
		}
	}

	return resultMatrix
}

func (s *lab3aService) GetSetNonDominatedAlternativesOnSet(ctx context.Context, sets [][]float64) []float64 {
	minRate := make([]float64, len(sets[0]))
	copy(minRate, sets[0])

	for i, _ := range sets {
		for j := range sets[i] {
			if minRate[j] > sets[i][j] {
				minRate[j] = math.Round(sets[i][j]*100) / 100
			}
		}
	}

	return minRate
}
