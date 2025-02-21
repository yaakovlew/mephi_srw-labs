package service

import (
	"context"

	"gonum.org/v1/gonum/mat"

	"backend/pkg/model"
	"backend/pkg/repository"
)

type external interface {
	SendLabMark(ctx context.Context, userId, labId, percentage int) error
	GetUserId(ctx context.Context, token string) (int, error)
}

type commonLab interface {
	GetLabResult(ctx context.Context, userId, labId int) (int, error)
	IncrementPercentageDone(ctx context.Context, userId, labId, mark int) error
	UpdateLabStep(ctx context.Context, userId, labId, step int) error
	GetUserIdByToken(labId int, token string) (int, error)
	SaveUserToken(userId, labId int, userHeader string) error
	GetUserInfo(userId, labId int) (model.UserRepo, error)
	OpenLab1ForStudent(ctx context.Context, userId, labId, externalLabId int) (bool, error)
	OpenLab3ForStudent(ctx context.Context, userId, labId, externalLabId int) (bool, error)
	CloseLabForStudent(ctx context.Context, userId, labId int) error
	ClearToken(userId, labId int) error
	GetLabCurrentStep(ctx context.Context, userId, labId int) (int, error)
	IsEmptyToken(userId, labId int) bool
	GetCurrentMark(userId, labId int) (int, error)
}

type lab3a interface {
	ValidateLab3AResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error)
	GetAlternativeDiffMatrices(ctx context.Context, alternativesNumber int, criteria [][]float64) []*mat.Dense
	GetAlternativesMatricesIntersection(ctx context.Context, matrices []*mat.Dense) *mat.Dense
	GetSetNonDominatedAlternatives(ctx context.Context, matrix *mat.Dense) []float64
	GetAlternativesMatricesWithCoefficients(ctx context.Context, matrices []*mat.Dense, coefficients []float64) *mat.Dense
	GetSetNonDominatedAlternativesOnSet(ctx context.Context, sets [][]float64) []float64

	CheckLab3AResult(ctx context.Context, userId int, userRes []float64, chosenIndex int) (int, int, []model.DataResponse, model.DataResponse, error)
	CheckLab3ASecondNonDominatedSets(ctx context.Context, userId int, userSet []float64) (int, int, []model.DataResponse, error)
	CheckLab3AAlternativesMatricesWithCoefficients(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error)
	CheckLab3ANonDominatedSet(ctx context.Context, userId int, userSets []float64) (int, int, []model.DataResponse, error)
	CheckLab3AMatricesIntersection(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error)
	CheckLab3AAlternativeDiffMatrices(ctx context.Context, userId int, step int, userMatrices [][]float64) (int, int, [][]model.DataResponse, error)
	CheckLab3AAlternativeSets(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error)
}

type lab3b interface {
	GetPointEstimateByAlternativeMatrix(ctx context.Context, matrix [][]model.Point) []float64
	GetPointsByLukasiewiczImplication(ctx context.Context, matrix [][]float64, ruleResult []string) ([][][]model.Point, [][]model.Point)
	ValidateLab3BResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error)
	NextMatrix(ctx context.Context, userId int) ([][]string, error)

	CheckLab3BRulesValue(ctx context.Context, userId int, step int, userMatrix map[string][]float64) (int, int, map[string][]model.DataResponse, error)
	CheckLab3BCriteriaValue(ctx context.Context, userId int, userMatrix [][]float64) (int, int, [][]model.DataResponse, error)
	CheckLab3BAllMatrices(ctx context.Context, userId int, step int, userPoints [][]model.Point) (int, int, [][]model.PointCheck, error)
	CheckLab3BMatricesIntersection(ctx context.Context, userId int, userPoints [][]model.Point) (int, int, [][]model.PointCheck, error)
	CheckLab3BAnswerLab3bLevelSet(ctx context.Context, userId int, step int, answerLevelSet []model.AnswerLevelSet) (int, int, model.AnswerLab3bLevelSetResponse, error)
	CheckLab3BResult(ctx context.Context, userId int, userIndex int, userSet []float64) (int, int, model.DataResponse, []model.DataResponse, error)
}

type lab3c interface {
	GetImportanceData(ctx context.Context, criteriaData, alternativeData []model.ImportancePoints, criteria [][]float64, criteriaWeight []float64) ([][]string, []string)
	GetResultAlternativesData(ctx context.Context, criteriaData, alternativeData []model.ImportancePoints, alternativeImportance [][]string, criteriaImportance []string) ([][][]model.Point, [][]model.Point, [][]model.Point)
	GetMiddleOfArea(ctx context.Context, points [][]model.Point) []float64
	ValidateLab3CResult(ctx context.Context, variance model.GeneratedLab3Variance) ([]float64, error)

	CheckLab3CMiddleOfArea(ctx context.Context, userId int, userIndex int, userMid []float64) (int, int, model.DataResponse, []model.DataResponse, error)
	CheckLab3CQuadraticParameters(ctx context.Context, userId int, userData model.AnswerLab3CQuadraticParametersRequest) (int, int, model.AnswerLab3CQuadraticParameters, error)
	CheckLab3CLineParameters(ctx context.Context, userId int, userData model.AnswerLab3CLineParametersRequest) (int, int, model.AnswerLab3CLineParameters, error)
	CheckLab3CArea(ctx context.Context, userId int, userArea model.AnswerLab3CArea) (int, int, model.DataResponse, error)
	CheckLab3CCriteriaMatrix(ctx context.Context, userId int, userMatrix [][]model.Point) (int, int, [][]model.PointCheck, error)
	CheckLab3CAlternativeMatrices(ctx context.Context, userId int, step int, userMatrices [][]model.Point) (int, int, [][]model.PointCheck, error)
	CheckLab3CCurrentMatrix(ctx context.Context, userId int, userEstimation model.AnswerLab3CCurrentMatrix) (int, int, [][]model.PointCheck, error)
	CheckLab3CEstimation(ctx context.Context, userId int, userEstimation model.AnswerLab3CEstimation) (int, int, []model.PointCheck, error)
	CheckLab3CCriteriaImportance(ctx context.Context, userId int, userCriteriaMatrix []string) (int, int, []model.DataResponse, error)
	CheckLab3CAlternativesImportance(ctx context.Context, userId int, step int, userAlterMatrix []string) (int, int, []model.DataResponse, error)
}

type lab1ab interface {
	CheckLab1AStep(ctx context.Context, userId int, step int, answer model.AnswerLab1ACommonMatrix) (int, int, model.AnswerLab1ACommonMatrixIsRight, error)
	CheckLab1A5Step(ctx context.Context, userId int, answer model.AnswerLab1A5Step) (int, int, model.AnswerLab1A5StepIsRight, error)

	AddAlternativesLab1B(ctx context.Context, userId int, alternatives []string) error
	AddCriteriasLab1B(ctx context.Context, userId int, criterias []model.MainCriteria) error
	CheckLab1BFirstStep(ctx context.Context, userId int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error)
	CheckLab1BSecondStep(ctx context.Context, userId int, step int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error)
	CheckLab1BAlignigIeracrhie(ctx context.Context, userId int, answer model.Lab1BMarkAligning) (int, int, model.Lab1BMarkAligningIsRight, error)
	CheckLab1BWeights(ctx context.Context, userId int, answer model.Lab1BWeights) (int, int, model.Lab1BWeightsIsRight, error)
	CheckLab1BCountCriteria(ctx context.Context, userId int, answer model.Lab1BCountCriteria) (int, int, model.Lab1BCountCriteriaISRight, error)
	CheckLab1QualityCriteria(ctx context.Context, userId int, index int, answer model.AnswerLab1BCommonMatrix) (int, int, model.AnswerLab1BCommonMatrixIsRight, error)
	CheckLab1BResult(ctx context.Context, userId int, answer model.AnswerLab1A5Step) (int, int, model.AnswerLab1A5StepIsRight, error)
	CheckMatrixIsCorrect(ctx context.Context, matrix [][]float64) bool
}

type taskLab3 interface {
	GenerateLab3Task(ctx context.Context, userId int) (model.GeneratedLab3Variance, error)
	UpdateUserVarianceLab3(ctx context.Context, userId, labId int, variance model.GeneratedLab3Variance) error
	GetVarianceLab3(ctx context.Context, userId, labId int) (model.UserLab3Task, [][]float64, error)
	CheckLab3Variance(ctx context.Context, userId, labId int) error
}

type taskLab1A interface {
	UpdateUserVarianceLab1A(ctx context.Context, userId int, labId int, variance model.GeneratedLab1AVariance) error
	CheckLab1AVariance(ctx context.Context, userId, labId int) error
	GetLab1AVariance(ctx context.Context, userId, labId int) (model.GeneratedLab1AVariance, error)
	GenerateLab1ATask(ctx context.Context, userId int) (model.GeneratedLab1AVariance, error)
}

type taskLab1B interface {
	UpdateUserVarianceLab1B(ctx context.Context, userId int, labId int, variance model.GeneratedVarianceLab1B) error
	CheckLab1BVariance(ctx context.Context, userId, labId int) error
	GetLab1BVariance(ctx context.Context, userId, labId int) (model.UserVarianceLab1B, error)
	GenerateLab1BTask(ctx context.Context, userId int) (model.GeneratedVarianceLab1B, error)
}

type Service struct {
	external
	taskLab1A
	taskLab1B
	taskLab3
	commonLab
	lab1ab
	lab3a
	lab3b
	lab3c
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		external:  NewExternalService(),
		taskLab1A: NewTaskLab1a(repo),
		taskLab1B: NewTaskLab1b(repo),
		taskLab3:  NewTaskLab3(repo),
		commonLab: NewCommonLabService(repo),
		lab1ab:    Newlab1ABService(repo),
		lab3a:     NewLab3aService(repo),
		lab3b:     NewLab3bService(repo),
		lab3c:     NewLab3cService(repo),
	}
}
