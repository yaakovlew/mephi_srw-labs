package service

import (
	"context"
	"fmt"
	"math"

	"backend/pkg/model"
	"backend/pkg/repository"
)

const Lab1AId = 1
const Lab1BId = 2

const osLimit = 10

var matrixSize map[int]float64

func init() {
	matrixSize = make(map[int]float64)
	matrixSize[1] = math.Round(100*0) / 100
	matrixSize[2] = math.Round(100*0) / 100
	matrixSize[3] = math.Round(100*0.58) / 100
	matrixSize[4] = math.Round(100*0.9) / 100
	matrixSize[5] = math.Round(100*1.12) / 100
	matrixSize[6] = math.Round(100*1.24) / 100
	matrixSize[7] = math.Round(100*1.32) / 100
	matrixSize[8] = math.Round(100*1.41) / 100
	matrixSize[9] = math.Round(100*1.45) / 100
	matrixSize[10] = math.Round(100*1.49) / 100
}

type lab1ABService struct {
	repo *repository.Repo
	commonEventService
}

func Newlab1ABService(repo *repository.Repo) *lab1ABService {
	return &lab1ABService{
		repo:               repo,
		commonEventService: NewCommonEventService(),
	}
}

func (l *lab1ABService) AddAlternativesLab1B(ctx context.Context, userId int, alternatives []string) error {
	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return err
	}

	variance.Variance.Variance.Alternatives = alternatives

	if err := l.repo.UpdateLab1BVariance(userId, Lab1BId, variance.Variance); err != nil {
		return err
	}

	return nil
}

func (l *lab1ABService) AddCriteriasLab1B(ctx context.Context, userId int, criterias []model.MainCriteria) error {
	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return err
	}

	variance.Variance.Variance.MainCriteria = criterias

	if err := l.repo.UpdateLab1BVariance(userId, Lab1BId, variance.Variance); err != nil {
		return err
	}

	return nil
}

func (l *lab1ABService) CheckLab1BFirstStep(ctx context.Context, userId int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error) {
	maxMark := 10

	data := answer.Matrix

	var totalCount, count = 0, 0

	x := l.getXValue(data)
	w := l.getWValue(x)
	mw := l.getMWValue(data, w)
	lambdaW := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(data, lambdaW)
	is := l.getISValue(lambdaMax, data)
	os := l.getOSValue(is, data)

	if os >= 10 {
		for i := range answer.Matrix {
			for j := range answer.Matrix[i] {
				answer.Matrix[i][j] = 1
			}
		}
		data = answer.Matrix
		x = l.getXValue(data)
		w = l.getWValue(x)
		mw = l.getMWValue(data, w)
		lambdaW = l.getLambdaWValue(mw, w)
		lambdaMax = l.getLambdaMaxValue(data, lambdaW)
		is = l.getISValue(lambdaMax, data)
		os = l.getOSValue(is, data)
	}

	if err := l.repo.Save1BVarianceMainCriteria(userId, Lab1BId, answer.Matrix); err != nil {
		return 0, maxMark, model.AnswerLab1BCommonMatrixIsRight{}, fmt.Errorf("failed save variance")
	}

	totalCount += len(x) + len(w) + len(mw) + len(lambdaW) + 3

	var resX []model.IsRight
	for i := range x {
		if x[i] == math.Round(100*answer.X[i])/100 {
			count++
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: true,
			})
		} else {
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: false,
			})
		}
	}

	var resultW []model.IsRight
	for i := range w {
		if w[i] == math.Round(100*answer.W[i])/100 {
			count++
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: true,
			})
		} else {
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: false,
			})
		}
	}

	var resultMW []model.IsRight
	for i := range mw {
		if mw[i] == math.Round(100*answer.MW[i])/100 {
			count++
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: true,
			})
		} else {
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: false,
			})
		}
	}

	var resultLambdaW []model.IsRight
	for i := range lambdaW {
		if lambdaW[i] == math.Round(100*answer.LambdaW[i])/100 {
			count++
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: true,
			})
		} else {
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: false,
			})
		}
	}

	lambdaMaxIsRight := model.IsRight{
		Val: lambdaMax,
	}
	if answer.LambdaMax == lambdaMax {
		count++
		lambdaMaxIsRight.IsRight = true
	} else {
		lambdaMaxIsRight.IsRight = false
	}

	isIsRight := model.IsRight{
		Val: is,
	}
	if answer.IS == is {
		count++
		isIsRight.IsRight = true
	} else {
		isIsRight.IsRight = false
	}

	osIsRight := model.IsRight{
		Val: os,
	}
	if answer.OS == os {
		count++
		osIsRight.IsRight = true
	} else {
		osIsRight.IsRight = false
	}

	var res model.AnswerLab1BCommonMatrixIsRight = model.AnswerLab1BCommonMatrixIsRight{
		X:         resX,
		W:         resultW,
		MW:        resultMW,
		LambdaW:   resultLambdaW,
		LambdaMax: lambdaMaxIsRight,
		IS:        isIsRight,
		OS:        osIsRight,
	}

	incorrect := totalCount - count

	return l.getMark(float64(maxMark), incorrect, totalCount), maxMark, res, nil
}

func (l *lab1ABService) CheckMatrixIsCorrect(ctx context.Context, matrix [][]float64) bool {
	x := l.getXValue(matrix)
	w := l.getWValue(x)
	mw := l.getMWValue(matrix, w)
	lambdaW := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(matrix, lambdaW)
	is := l.getISValue(lambdaMax, matrix)
	os := l.getOSValue(is, matrix)

	if os >= 10 {
		return false
	} else {
		return true
	}
}

func (l *lab1ABService) CheckLab1BSecondStep(ctx context.Context, userId int, step int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error) {
	var maxMarkAllStep float64 = 30

	data := answer.Matrix

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, 0, model.AnswerLab1BCommonMatrixIsRight{}, err
	}

	maxMark := maxMarkAllStep / float64(len(variance.CriteriaMatrix))
	maxMarkInt := int(math.Round(maxMarkAllStep / float64(len(variance.CriteriaMatrix))))

	var totalCount, count = 0, 0

	x := l.getXValue(data)
	w := l.getWValue(x)
	mw := l.getMWValue(data, w)
	lambdaW := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(data, lambdaW)
	is := l.getISValue(lambdaMax, data)
	os := l.getOSValue(is, data)

	if os >= 10 {
		for i := range answer.Matrix {
			for j := range answer.Matrix[i] {
				answer.Matrix[i][j] = 1
			}
		}
		data = answer.Matrix
		x = l.getXValue(data)
		w = l.getWValue(x)
		mw = l.getMWValue(data, w)
		lambdaW = l.getLambdaWValue(mw, w)
		lambdaMax = l.getLambdaMaxValue(data, lambdaW)
		is = l.getISValue(lambdaMax, data)
		os = l.getOSValue(is, data)
	}

	if err := l.repo.Save1BVarianceCriteriaMatrix(userId, Lab1BId, step, answer.Matrix); err != nil {
		return 0, maxMarkInt, model.AnswerLab1BCommonMatrixIsRight{}, fmt.Errorf("failed save variance")
	}

	totalCount += len(x) + len(w) + len(mw) + len(lambdaW) + 3

	var resX []model.IsRight
	for i := range x {
		if x[i] == math.Round(100*answer.X[i])/100 {
			count++
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: true,
			})
		} else {
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: false,
			})
		}
	}

	var resultW []model.IsRight
	for i := range w {
		if w[i] == math.Round(100*answer.W[i])/100 {
			count++
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: true,
			})
		} else {
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: false,
			})
		}
	}

	var resultMW []model.IsRight
	for i := range mw {
		if mw[i] == math.Round(100*answer.MW[i])/100 {
			count++
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: true,
			})
		} else {
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: false,
			})
		}
	}

	var resultLambdaW []model.IsRight
	for i := range lambdaW {
		if lambdaW[i] == math.Round(100*answer.LambdaW[i])/100 {
			count++
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: true,
			})
		} else {
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: false,
			})
		}
	}

	lambdaMaxIsRight := model.IsRight{
		Val: lambdaMax,
	}
	if answer.LambdaMax == lambdaMax {
		count++
		lambdaMaxIsRight.IsRight = true
	} else {
		lambdaMaxIsRight.IsRight = false
	}

	isIsRight := model.IsRight{
		Val: is,
	}
	if answer.IS == is {
		count++
		isIsRight.IsRight = true
	} else {
		isIsRight.IsRight = false
	}

	osIsRight := model.IsRight{
		Val: os,
	}
	if answer.OS == os {
		count++
		osIsRight.IsRight = true
	} else {
		osIsRight.IsRight = false
	}

	var res model.AnswerLab1BCommonMatrixIsRight = model.AnswerLab1BCommonMatrixIsRight{
		X:         resX,
		W:         resultW,
		MW:        resultMW,
		LambdaW:   resultLambdaW,
		LambdaMax: lambdaMaxIsRight,
		IS:        isIsRight,
		OS:        osIsRight,
	}

	incorrect := totalCount - count

	return l.getMark(maxMark, incorrect, totalCount), maxMarkInt, res, nil
}

func (l *lab1ABService) CheckLab1BAlignigIeracrhie(ctx context.Context, userId int, answer model.Lab1BMarkAligning) (int, int, model.Lab1BMarkAligningIsRight, error) {
	maxMark := 5

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, maxMark, model.Lab1BMarkAligningIsRight{}, err
	}

	x := l.getXValue(variance.MainCriteriaMatrix)
	w := l.getWValue(x)
	mw := l.getMWValue(variance.MainCriteriaMatrix, w)
	lambda := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(variance.MainCriteriaMatrix, lambda)
	is := l.getISValue(lambdaMax, variance.MainCriteriaMatrix)

	var ises []float64
	var wis []float64
	var cc []float64
	var wCC []float64
	var m float64 = is
	var tildaM float64 = math.Round(100*matrixSize[len(variance.MainCriteriaMatrix)]) / 100
	for i := range variance.CriteriaMatrix {
		xCriteria := l.getXValue(variance.CriteriaMatrix[i])
		wCriteria := l.getWValue(xCriteria)
		mwCriteria := l.getMWValue(variance.CriteriaMatrix[i], wCriteria)
		lambdaCriteria := l.getLambdaWValue(mwCriteria, wCriteria)
		lambdaMaxCriteria := l.getLambdaMaxValue(variance.CriteriaMatrix[i], lambdaCriteria)
		isCriteria := l.getISValue(lambdaMaxCriteria, variance.CriteriaMatrix[i])
		ises = append(ises, isCriteria)
		wisCount := math.Round(100*isCriteria*w[i]) / 100
		wis = append(wis, wisCount)
		m += wisCount
		os := math.Round(100*matrixSize[len(variance.CriteriaMatrix[i])]) / 100
		cc = append(cc, os)
		wccCount := math.Round(100*w[i]*matrixSize[len(variance.CriteriaMatrix[i])]) / 100
		wCC = append(wCC, wccCount)
		tildaM += wccCount
	}

	totalCount := len(wis) + len(wCC) + 3
	count := 0
	osi := math.Round(10000*m/tildaM) / 100

	var wisIsRigth []model.IsRight
	for i := range wis {
		data := model.IsRight{
			Val: wis[i],
		}
		if wis[i] == answer.WIS[i] {
			count++
			data.IsRight = true
		} else {
			data.IsRight = false
		}
		wisIsRigth = append(wisIsRigth, data)
	}

	var wCCIsRigth []model.IsRight
	for i := range wCC {
		data := model.IsRight{
			Val: wCC[i],
		}
		if wCC[i] == answer.WCC[i] {
			count++
			data.IsRight = true
		} else {
			data.IsRight = false
		}
		wCCIsRigth = append(wCCIsRigth, data)
	}

	var mISRight model.IsRight = model.IsRight{
		Val:     m,
		IsRight: false,
	}
	if mISRight.Val == answer.M {
		count++
		mISRight.IsRight = true
	}

	var mTildaISRight model.IsRight = model.IsRight{
		Val:     tildaM,
		IsRight: false,
	}
	if mTildaISRight.Val == answer.TildaM {
		count++
		mTildaISRight.IsRight = true
	}

	var osiIsRight model.IsRight = model.IsRight{
		Val:     osi,
		IsRight: false,
	}
	if osiIsRight.Val == answer.OSI {
		count++
		osiIsRight.IsRight = true
	}

	return l.getMark(float64(maxMark), totalCount-count, totalCount), maxMark, model.Lab1BMarkAligningIsRight{
		WIS:    wisIsRigth,
		WCC:    wCCIsRigth,
		M:      mISRight,
		TildaM: mTildaISRight,
		OSI:    osiIsRight,
	}, nil
}

func (l *lab1ABService) CheckLab1BWeights(ctx context.Context, userId int, answer model.Lab1BWeights) (int, int, model.Lab1BWeightsIsRight, error) {
	maxMark := 10

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, maxMark, model.Lab1BWeightsIsRight{}, err
	}

	x := l.getXValue(variance.MainCriteriaMatrix)
	w := l.getWValue(x)

	total := 0
	count := 0
	var weights []model.IsRight
	for i := range variance.CriteriaMatrix {
		xCriteria := l.getXValue(variance.CriteriaMatrix[i])
		wCriteria := l.getWValue(xCriteria)
		for j := range wCriteria {
			total++
			weights = append(weights, model.IsRight{
				Val:     math.Round(100*w[i]*wCriteria[j]) / 100,
				IsRight: false,
			})
		}
	}

	for i := range weights {
		if weights[i].Val == answer.Weights[i] {
			weights[i].IsRight = true
			count++
		}
	}

	return l.getMark(float64(maxMark), total-count, total), maxMark, model.Lab1BWeightsIsRight{Weights: weights}, nil
}

func (l *lab1ABService) CheckLab1BCountCriteria(ctx context.Context, userId int, answer model.Lab1BCountCriteria) (int, int, model.Lab1BCountCriteriaISRight, error) {
	maxMarkStep := 15

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, 0, model.Lab1BCountCriteriaISRight{}, err
	}

	total := 0
	count := 0
	countCountCriteria := 0

	var marks [][]model.IsRight
	for i := range variance.Variance.Variance.MainCriteria {
		for j := range variance.Variance.Variance.MainCriteria[i].Extra {
			if variance.Variance.Variance.MainCriteria[i].Extra[j].IsCount {
				countCountCriteria++
				total += len(variance.Variance.Variance.MainCriteria[i].Extra[j].Value)
				var sum float64 = 0
				for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
					sum += v
				}
				sum = math.Round(100*sum) / 100

				var mark []model.IsRight
				if variance.Variance.Variance.MainCriteria[i].Extra[j].IsReverse {
					var data []float64
					var sumReverse float64 = 0
					for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
						val := math.Round(100/(v/sum)) / 100
						sumReverse += val
						data = append(data, val)
					}
					for _, v := range data {
						mark = append(mark, model.IsRight{
							Val:     math.Round(100*v/sumReverse) / 100,
							IsRight: false,
						})
					}
				} else {
					for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
						mark = append(mark, model.IsRight{
							Val:     math.Round(100*v/sum) / 100,
							IsRight: false,
						})
					}
				}
				marks = append(marks, mark)
			} else {
				continue
			}
		}
	}

	var maxMark float64 = float64(maxMarkStep) / float64(countCountCriteria)
	var maxMarkInt int = int(math.Round(float64(maxMarkStep) / float64(countCountCriteria)))

	for i := range marks {
		for j := range marks[i] {
			if marks[i][j].Val == answer.Marks[i][j] {
				marks[i][j].IsRight = true
				count++
			}
		}
	}

	return l.getMark(float64(maxMark), total-count, total), maxMarkInt, model.Lab1BCountCriteriaISRight{
		Marks: marks,
	}, nil
}

func (l *lab1ABService) CheckLab1QualityCriteria(ctx context.Context, userId int, index int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error) {
	maxMarkStep := 15

	data := answer.Matrix

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, 0, model.AnswerLab1BCommonMatrixIsRight{}, err
	}

	var totalCount, count = 0, 0
	qualityCriteriaCount := 0

	for i := range variance.Variance.Variance.MainCriteria {
		for j := range variance.Variance.Variance.MainCriteria[i].Extra {
			if !variance.Variance.Variance.MainCriteria[i].Extra[j].IsCount {
				qualityCriteriaCount++
			}
		}
	}

	if index >= len(variance.QualityCriteriaMatrix) {
		return 0, 0, model.AnswerLab1BCommonMatrixIsRight{}, fmt.Errorf("invalid index")
	}

	x := l.getXValue(data)
	w := l.getWValue(x)
	mw := l.getMWValue(data, w)
	lambdaW := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(data, lambdaW)
	is := l.getISValue(lambdaMax, data)
	os := l.getOSValue(is, data)

	if os >= 10 {
		for i := range answer.Matrix {
			for j := range answer.Matrix[i] {
				answer.Matrix[i][j] = 1
			}
		}
		data = answer.Matrix
		x = l.getXValue(data)
		w = l.getWValue(x)
		mw = l.getMWValue(data, w)
		lambdaW = l.getLambdaWValue(mw, w)
		lambdaMax = l.getLambdaMaxValue(data, lambdaW)
		is = l.getISValue(lambdaMax, data)
		os = l.getOSValue(is, data)
	}

	var maxMark float64 = float64(maxMarkStep) / float64(qualityCriteriaCount)
	var maxMarkInt int = int(math.Round(float64(maxMarkStep) / float64(qualityCriteriaCount)))

	if err := l.repo.Save1BVarianceQualityMatrix(userId, Lab1BId, index, answer.Matrix); err != nil {
		return 0, maxMarkInt, model.AnswerLab1BCommonMatrixIsRight{}, fmt.Errorf("failed save variance")
	}

	totalCount += len(x) + len(w) + len(mw) + len(lambdaW) + 3

	var resX []model.IsRight
	for i := range x {
		if x[i] == math.Round(100*answer.X[i])/100 {
			count++
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: true,
			})
		} else {
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: false,
			})
		}
	}

	var resultW []model.IsRight
	for i := range w {
		if w[i] == math.Round(100*answer.W[i])/100 {
			count++
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: true,
			})
		} else {
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: false,
			})
		}
	}

	var resultMW []model.IsRight
	for i := range mw {
		if mw[i] == math.Round(100*answer.MW[i])/100 {
			count++
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: true,
			})
		} else {
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: false,
			})
		}
	}

	var resultLambdaW []model.IsRight
	for i := range lambdaW {
		if lambdaW[i] == math.Round(100*answer.LambdaW[i])/100 {
			count++
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: true,
			})
		} else {
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: false,
			})
		}
	}

	lambdaMaxIsRight := model.IsRight{
		Val: lambdaMax,
	}
	if answer.LambdaMax == lambdaMax {
		count++
		lambdaMaxIsRight.IsRight = true
	} else {
		lambdaMaxIsRight.IsRight = false
	}

	isIsRight := model.IsRight{
		Val: is,
	}
	if answer.IS == is {
		count++
		isIsRight.IsRight = true
	} else {
		isIsRight.IsRight = false
	}

	osIsRight := model.IsRight{
		Val: os,
	}
	if answer.OS == os {
		count++
		osIsRight.IsRight = true
	} else {
		osIsRight.IsRight = false
	}

	var res model.AnswerLab1BCommonMatrixIsRight = model.AnswerLab1BCommonMatrixIsRight{
		X:         resX,
		W:         resultW,
		MW:        resultMW,
		LambdaW:   resultLambdaW,
		LambdaMax: lambdaMaxIsRight,
		IS:        isIsRight,
		OS:        osIsRight,
	}

	incorrect := totalCount - count

	return l.getMark(float64(maxMark), incorrect, totalCount), maxMarkInt, res, nil
}

func (l *lab1ABService) CheckLab1BResult(ctx context.Context, userId int, answer model.AnswerLab1A5Step) (int, int, model.AnswerLab1A5StepIsRight, error) {
	maxMark1 := 20
	maxMark2 := 5

	variance, err := l.repo.GetLab1BVariance(userId, Lab1BId)
	if err != nil {
		return 0, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{}, err
	}

	var mat1 [][]float64
	qualityMatrix := 0
	countMatrix := 0
	for i := range variance.Variance.Variance.MainCriteria {
		for j := range variance.Variance.Variance.MainCriteria[i].Extra {
			if variance.Variance.Variance.MainCriteria[i].Extra[j].IsCount {
				var sum float64 = 0
				for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
					sum += v
				}
				sum = math.Round(100*sum) / 100

				var mark []float64
				if variance.Variance.Variance.MainCriteria[i].Extra[j].IsReverse {
					var data []float64
					var sumReverse float64 = 0
					for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
						val := math.Round(100/(v/sum)) / 100
						sumReverse += val
						data = append(data, val)
					}
					for _, v := range data {
						mark = append(mark, math.Round(100*v/sumReverse)/100)
					}
				} else {
					for _, v := range variance.Variance.Variance.MainCriteria[i].Extra[j].Value {
						mark = append(mark, math.Round(100*v/sum)/100)
					}
				}

				mat1 = append(mat1, mark)
				countMatrix++
			} else {
				data := variance.QualityCriteriaMatrix[qualityMatrix]
				x := l.getXValue(data)
				w := l.getWValue(x)

				mat1 = append(mat1, w)

				qualityMatrix++
			}
		}
	}

	x := l.getXValue(variance.MainCriteriaMatrix)
	w := l.getWValue(x)
	var weights []float64
	for i := range variance.CriteriaMatrix {
		xCriteria := l.getXValue(variance.CriteriaMatrix[i])
		wCriteria := l.getWValue(xCriteria)
		for j := range wCriteria {
			weights = append(weights, math.Round(100*w[i]*wCriteria[j])/100)
		}
	}

	mat2 := make([][]float64, 1)
	mat2[0] = weights

	res, err := l.multiplyMatrices(l.transposeMatrix(mat1), l.transposeMatrix(mat2))
	if err != nil {
		return 0, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{}, err
	}

	set := l.transposeMatrix(res)[0]
	maxIndex := 0
	maxCount := set[0]
	for i := range set {
		if set[i] > maxCount {
			maxIndex = i
			maxCount = set[i]
		}
	}

	maxIndex++

	totalCount1 := len(mat1)
	count1 := 0

	var resSet []model.IsRight
	for i := range set {
		if answer.Set[i] == set[i] {
			resSet = append(resSet, model.IsRight{
				Val:     set[i],
				IsRight: true,
			})
			count1++
		} else {
			resSet = append(resSet, model.IsRight{
				Val:     set[i],
				IsRight: false,
			})
		}
	}

	mark := l.getMark(float64(maxMark1), totalCount1-count1, totalCount1)

	resIndex := model.IsRightInt{
		Val:     maxIndex,
		IsRight: false,
	}
	if maxIndex == answer.Index {
		resIndex.IsRight = true
		mark += maxMark2
	}

	return mark, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{
		Set:   resSet,
		Index: resIndex,
	}, nil
}

func (l *lab1ABService) GetVarianceLab1A(ctx context.Context, userId, labId int) (model.GeneratedLab1AVariance, error) {
	variance, err := l.repo.GetLab1AVariance(userId, labId)
	if err != nil {
		return model.GeneratedLab1AVariance{}, err
	}

	return variance, nil
}

func (l *lab1ABService) CheckLab1AStep(ctx context.Context, userId int, step int, answer model.AnswerLab1ACommonMatrix) (int, int, model.AnswerLab1ACommonMatrixIsRight, error) {
	maxMark := 15

	if step > 4 {
		return 0, maxMark, model.AnswerLab1ACommonMatrixIsRight{}, fmt.Errorf("the big step count")
	}

	variance, err := l.repo.GetLab1AVariance(userId, Lab1AId)
	if err != nil {
		return 0, maxMark, model.AnswerLab1ACommonMatrixIsRight{}, err
	}

	data := variance.Variance.Matrices[step]

	var totalCount, count = 0, 0

	x := l.getXValue(data)
	w := l.getWValue(x)
	mw := l.getMWValue(data, w)
	lambdaW := l.getLambdaWValue(mw, w)
	lambdaMax := l.getLambdaMaxValue(data, lambdaW)
	is := l.getISValue(lambdaMax, data)
	os := l.getOSValue(is, data)

	totalCount += len(x) + len(w) + len(mw) + len(lambdaW) + 3

	resW := w
	resMatrix := data

	for os >= 10 {
		data = l.matrixCorrect2(resMatrix, resW)

		x = l.getXValue(data)

		w = l.getWValue(x)

		mw = l.getMWValue(data, w)

		lambdaW = l.getLambdaWValue(mw, w)
		lambdaMax = l.getLambdaMaxValue(data, lambdaW)
		is = l.getISValue(lambdaMax, data)
		os = l.getOSValue(is, data)

		resMatrix = data
	}

	var resX []model.IsRight
	for i := range x {
		if x[i] == math.Round(100*answer.X[i])/100 {
			count++
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: true,
			})
		} else {
			resX = append(resX, model.IsRight{
				Val:     x[i],
				IsRight: false,
			})
		}
	}

	var resultW []model.IsRight
	for i := range w {
		if w[i] == math.Round(100*answer.W[i])/100 {
			count++
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: true,
			})
		} else {
			resultW = append(resultW, model.IsRight{
				Val:     w[i],
				IsRight: false,
			})
		}
	}

	var resultMW []model.IsRight
	for i := range mw {
		if mw[i] == math.Round(100*answer.MW[i])/100 {
			count++
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: true,
			})
		} else {
			resultMW = append(resultMW, model.IsRight{
				Val:     mw[i],
				IsRight: false,
			})
		}
	}

	var resultLambdaW []model.IsRight
	for i := range lambdaW {
		if lambdaW[i] == math.Round(100*answer.LambdaW[i])/100 {
			count++
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: true,
			})
		} else {
			resultLambdaW = append(resultLambdaW, model.IsRight{
				Val:     lambdaW[i],
				IsRight: false,
			})
		}
	}

	lambdaMaxIsRight := model.IsRight{
		Val: lambdaMax,
	}
	if answer.LambdaMax == lambdaMax {
		count++
		lambdaMaxIsRight.IsRight = true
	} else {
		lambdaMaxIsRight.IsRight = false
	}

	isIsRight := model.IsRight{
		Val: is,
	}
	if answer.IS == is {
		count++
		isIsRight.IsRight = true
	} else {
		isIsRight.IsRight = false
	}

	osIsRight := model.IsRight{
		Val: os,
	}
	if answer.OS == os {
		count++
		osIsRight.IsRight = true
	} else {
		osIsRight.IsRight = false
	}

	var res model.AnswerLab1ACommonMatrixIsRight = model.AnswerLab1ACommonMatrixIsRight{
		X:         resX,
		W:         resultW,
		MW:        resultMW,
		LambdaW:   resultLambdaW,
		LambdaMax: lambdaMaxIsRight,
		IS:        isIsRight,
		OS:        osIsRight,
	}

	incorrect := totalCount - count

	return l.getMark(float64(maxMark), incorrect, totalCount), maxMark, res, nil
}

func (l *lab1ABService) CheckLab1A5Step(ctx context.Context, userId int, answer model.AnswerLab1A5Step) (int, int, model.AnswerLab1A5StepIsRight, error) {
	maxMark1 := 15
	maxMark2 := 10

	variance, err := l.repo.GetLab1AVariance(userId, Lab1AId)
	if err != nil {
		return 0, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{}, err
	}

	mat1 := make([][]float64, len(variance.Variance.Matrices)-1)
	for i := 1; i < len(variance.Variance.Matrices); i++ {
		x := l.getXValue(variance.Variance.Matrices[i])
		mat1[i-1] = l.getWValue(x)
	}

	mat2 := make([][]float64, 1)
	x := l.getXValue(variance.Variance.Matrices[0])
	mat2[0] = l.getWValue(x)

	res, err := l.multiplyMatrices(l.transposeMatrix(mat1), l.transposeMatrix(mat2))
	if err != nil {
		return 0, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{}, err
	}

	set := l.transposeMatrix(res)[0]
	maxIndex := 0
	maxCount := set[0]
	for i := range set {
		if set[i] > maxCount {
			maxIndex = i
			maxCount = set[i]
		}
	}

	maxIndex++

	totalCount1 := len(set)
	count1 := 0

	var resSet []model.IsRight
	for i := range set {
		if answer.Set[i] == set[i] {
			resSet = append(resSet, model.IsRight{
				Val:     set[i],
				IsRight: true,
			})
			count1++
		} else {
			resSet = append(resSet, model.IsRight{
				Val:     set[i],
				IsRight: false,
			})
		}
	}

	mark := l.getMark(float64(maxMark1), totalCount1-count1, totalCount1)

	resIndex := model.IsRightInt{
		Val:     maxIndex,
		IsRight: false,
	}
	if maxIndex == answer.Index {
		resIndex.IsRight = true
		mark += maxMark2
	}

	return mark, maxMark1 + maxMark2, model.AnswerLab1A5StepIsRight{
		Set:   resSet,
		Index: resIndex,
	}, nil
}

func (l *lab1ABService) getXValue(matrix [][]float64) []float64 {
	var x []float64 = make([]float64, len(matrix))

	for i := range matrix {
		var count float64 = 1
		for j := range matrix[i] {
			count *= matrix[i][j]
		}

		x[i] = math.Round(math.Pow(count, 1.0/float64(len(matrix[i])))*100) / 100
	}

	return x
}

func (l *lab1ABService) getWValue(data []float64) []float64 {
	var sum float64 = 0

	for i := range data {
		sum += data[i]
	}

	var res []float64
	for i := range data {
		res = append(res, math.Round(100*data[i]/sum)/100)
	}

	return res
}

func (l *lab1ABService) getMWValue(matrix [][]float64, w []float64) []float64 {
	var x []float64 = make([]float64, len(matrix))

	for i := range matrix {
		var count float64 = 0
		for j := range matrix[i] {
			count += matrix[i][j] * w[j]
		}

		x[i] = math.Round(count*100) / 100
	}

	return x
}

func (l *lab1ABService) getLambdaWValue(mw, w []float64) []float64 {
	var res []float64
	for i := range mw {
		res = append(res, math.Round(100*mw[i]/w[i])/100)
	}

	return res
}

func (l *lab1ABService) getLambdaMaxValue(matrix [][]float64, lambdaWValue []float64) float64 {
	var sum float64 = 0

	for i := range lambdaWValue {
		sum += lambdaWValue[i]
	}

	return math.Round(100*(sum/float64(len(matrix)))) / 100
}

func (l *lab1ABService) getISValue(labmdaMax float64, matrix [][]float64) float64 {
	first := math.Round(100*(labmdaMax-float64(len(matrix)))) / 100

	second := float64(len(matrix)) - 1
	return math.Round(100*first/second) / 100
}

func (l *lab1ABService) getOSValue(isValue float64, matrix [][]float64) float64 {
	if v, ok := matrixSize[len(matrix)]; !ok {
		return 0
	} else {
		return math.Round(10000*isValue/v) / 100
	}
}

func (l *lab1ABService) matrixCorrect1(matrix [][]float64, w []float64) [][]float64 {
	var correct [][]float64 = make([][]float64, len(matrix))
	for i := range correct {
		correct[i] = make([]float64, len(matrix[i]))
	}

	for i := range correct {
		for j := range correct[i] {
			correct[i][j] = math.Abs(math.Round(100 * (matrix[i][j] - (w[i] / w[j])) / 100))
		}
	}

	return correct
}

func (l *lab1ABService) matrixCorrect1Sum(matrix [][]float64) []float64 {
	var res []float64
	for i := range matrix {
		var count float64 = 0
		for j := range matrix[i] {
			count += matrix[i][j]
		}

		res = append(res, math.Round(100*count)/100)
	}

	return res
}

func (l *lab1ABService) matrixCorrect2(matrix [][]float64, w []float64) [][]float64 {
	var correct [][]float64 = make([][]float64, len(matrix))
	for i := range correct {
		correct[i] = make([]float64, len(matrix[i]))
	}

	for i := range correct {
		for j := range correct[i] {
			correct[i][j] = math.Round(100*w[i]/w[j]) / 100
		}
	}

	return correct
}

func (l *lab1ABService) getRes(w []float64, data [][]float64) ([]float64, error) {
	wMatrix := make([][]float64, 1)
	wMatrix[0] = w

	res, err := l.multiplyMatrices(data, wMatrix)
	if err != nil {
		return nil, err
	}

	if res != nil {
		return res[0], nil
	} else {
		return nil, fmt.Errorf("failed multiply matrices")
	}
}

func (l *lab1ABService) getAlternativeIndex(data []float64) int {
	var maxData float64 = data[0]
	var index = 0

	for i := range data {
		if data[i] >= maxData {
			maxData = data[i]
			index = i
		}
	}

	return index + 1
}

func (l *lab1ABService) transposeMatrix(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	transposed := make([][]float64, cols)
	for i := range transposed {
		transposed[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func (l *lab1ABService) multiplyMatrices(a [][]float64, b [][]float64) ([][]float64, error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("incorrect matrix size")
	}

	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("incorrect matrix sizes a:%d b:%d", len(a[0]), len(b))
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}

	for i := range a {
		for j := range b[0] {
			for k := range b {
				result[i][j] += a[i][k] * b[k][j]
			}
			result[i][j] = math.Round(100*result[i][j]) / 100
		}
	}

	return result, nil
}

func (l *lab1ABService) getMarkWISConsistency(w []float64, is []float64) []float64 {
	var res []float64

	for i := range w {
		res = append(res, w[i]*is[i])
	}

	return res
}

func (l *lab1ABService) getCCConsistency(matrix [][]float64) (float64, error) {
	if v, ok := matrixSize[len(matrix)]; !ok {
		return 0, fmt.Errorf("not found size")
	} else {
		return v, nil
	}
}

func (l *lab1ABService) getMarkWCCConsistency(w []float64, cc []float64) []float64 {
	var res []float64

	for i := range w {
		res = append(res, math.Round(100*w[i]*cc[i])/100)
	}

	return res
}

func (l *lab1ABService) getM(w []float64, is float64) float64 {
	var sum float64 = is

	for i := range w {
		sum += w[i]
	}

	sum = math.Round(100*sum) / 100

	return sum
}

func (l *lab1ABService) getTildaM(wCC []float64, size float64) float64 {
	var sum float64 = size

	for i := range wCC {
		sum += wCC[i]
	}

	sum = math.Round(100*sum) / 100

	return sum
}

func (l *lab1ABService) getOsi(m, tildaM float64) float64 {
	return math.Round(100*m/tildaM) / 100
}

func (l *lab1ABService) getBeforeFirstLevel(wGroup, wCriteria float64) float64 {
	return math.Round(100*wGroup*wCriteria) / 100
}

func (l *lab1ABService) getCountCriteria(isReverse bool, data []float64) [][]float64 {
	if !isReverse {
		var result [][]float64 = make([][]float64, 1)
		var res []float64

		var sum float64 = 0
		for i := range data {
			sum += data[i]
		}

		for i := range data {
			res = append(res, math.Round(100*data[i]/sum)/100)
		}

		result[0] = res

		return result
	} else {
		var result [][]float64 = make([][]float64, 2)
		var res []float64

		var sum float64 = 0
		for i := range data {
			sum += data[i]
		}

		for i := range data {
			res = append(res, math.Round(100/(data[i]/sum))/100)
		}

		result[0] = res

		var secSum float64 = 0
		for i := range res {
			secSum += res[i]
		}

		var secRes []float64
		for i := range res {
			secRes = append(secRes, math.Round(100*res[i]/secSum)/100)
		}

		result[1] = secRes

		return result
	}
}
