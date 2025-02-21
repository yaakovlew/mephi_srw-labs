package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/pkg/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://mephi71.ru:9081", "https://mephi71.ru:9082"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "lab-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	lab1a := router.Group("/lab1a")
	{
		lecturer := lab1a.Group("/open", h.CheckFirstAHeaderLecturer)
		{
			lecturer.POST("", h.OpenFirstALabForStudent)
		}

		student1a := lab1a.Group("/variant")
		{
			security := student1a.Group("", h.CheckFirstAHeaderStudentForStart)
			{
				security.GET("", h.OpenFirstALab)
			}

			notSecurityLab1a := student1a.Group("", h.CheckFirstAHeaderStudent)
			{
				notSecurityLab1a.POST("/:step", h.SendLab1AStep)
				notSecurityLab1a.POST("/result", h.SendLab1AResult)
			}
		}
	}

	lab1b := router.Group("/lab1b")
	{
		lecturer := lab1b.Group("/open", h.CheckFirstBHeaderLecturer)
		{
			lecturer.POST("", h.OpenFirstBLabForStudent)
		}

		student1b := lab1b.Group("/variant")
		{
			security := student1b.Group("", h.CheckFirstBHeaderStudentForStart)
			{
				security.GET("", h.OpenFirstBLab)
			}

			notSecurityLab1b := student1b.Group("", h.CheckFirstBHeaderStudent)
			{
				notSecurityLab1b.POST("/alternative", h.AddLab1BAlternatives)
				notSecurityLab1b.POST("/criteria", h.AddLab1BCriterias)
				notSecurityLab1b.POST("/check_matrix", h.CheckMatrixIsCorrect)
				notSecurityLab1b.POST("/1", h.SendLab1BMainCriteriaStep)
				notSecurityLab1b.POST("/2-4", h.SendLab1BCriteriaStep)
				notSecurityLab1b.PUT("/2", h.UpdateSecondStep)
				notSecurityLab1b.POST("/5", h.SendLab1BMarkAligning)
				notSecurityLab1b.PUT("/5", h.UpdateFivthStep)
				notSecurityLab1b.POST("/6", h.SendLab1BCountCriteria)
				notSecurityLab1b.POST("/quality", h.SendLab1BQualityCriteria)
				notSecurityLab1b.POST("/result", h.SendLab1BResult)
			}
		}
	}

	lab3a := router.Group("/lab3a")
	{
		lecturerLab3a := lab3a.Group("/open", h.CheckThirdAHeaderLecturer)
		{
			lecturerLab3a.POST("", h.OpenThirdALabForStudent)
		}

		studentLab3a := lab3a.Group("/variant")
		{
			securityLab3a := studentLab3a.Group("", h.CheckThirdAHeaderStudentForStart)
			{
				securityLab3a.GET("", h.OpenThirdALab)
			}

			notSecurityLab3a := studentLab3a.Group("", h.CheckThirdAHeaderStudent)
			{
				notSecurityLab3a.POST("/alternative-sets", h.SendLab3AAlternativeSets)
				notSecurityLab3a.POST("/increment-second-step", h.IncrementLab3ASecondStep)
				notSecurityLab3a.POST("/diff-matrices", h.SendLab3AAlternativeDiffMatrices)
				notSecurityLab3a.POST("/intersection", h.SendLab3AAlternativeMatricesIntersection)
				notSecurityLab3a.POST("/non-dominated", h.SendLab3ANonDominatedSet)
				notSecurityLab3a.POST("/coff-matrices", h.SendLab3AMatricesWithCoefficients)
				notSecurityLab3a.POST("/second-non-dominated", h.SendLab3ASecondNonDominatedSets)
				notSecurityLab3a.POST("/result", h.SendLab3AResult)
				notSecurityLab3a.GET("/info", h.GetCurrentStepLab3A)
			}
		}
	}

	lab3b := router.Group("/lab3b")
	{
		lecturerLab3b := lab3b.Group("/open", h.CheckThirdBHeaderLecturer)
		{
			lecturerLab3b.POST("", h.OpenThirdBLabForStudent)
		}

		studentLab3b := lab3b.Group("/variant")
		{
			securityLab3b := studentLab3b.Group("", h.CheckThirdBHeaderStudentForStart)
			{
				securityLab3b.GET("", h.OpenThirdBLab)
			}

			notSecurityLab3b := studentLab3b.Group("", h.CheckThirdBHeaderStudent)
			{
				notSecurityLab3b.POST("/increment-zero-step", h.IncrementLabBZeroStep)
				notSecurityLab3b.POST("/increment-second-step", h.IncrementLabBSecondStep)
				//notSecurityLab3b.POST("/increment-fourth-step", h.IncrementLabBFourthStep)
				notSecurityLab3b.POST("/rule-value", h.SendLab3BValueByRule)
				notSecurityLab3b.POST("/rule-number", h.SendLab3BRuleNumber)
				notSecurityLab3b.POST("/all-matrices", h.SendLab3BAllMatrices)
				notSecurityLab3b.POST("/intersection", h.SendLab3BMatricesIntersection)
				//notSecurityLab3b.POST("/level-set", h.SendLab3BLevelSet)
				notSecurityLab3b.POST("/result", h.SendLab3BResult)
				notSecurityLab3b.GET("/info", h.GetCurrentStepLab3B)
			}
		}
	}

	lab3c := router.Group("/lab3c")
	{
		lecturer := lab3c.Group("/open", h.CheckThirdCHeaderLecturer)
		{
			lecturer.POST("", h.OpenThirdCLabForStudent)
		}

		student3c := lab3c.Group("/variant")
		{
			security := student3c.Group("", h.CheckThirdCHeaderStudentForStart)
			{
				security.GET("", h.OpenThirdCLab)
			}

			notSecurityLab3c := student3c.Group("", h.CheckThirdCHeaderStudent)
			{
				notSecurityLab3c.POST("/increment-zero-step", h.IncrementLabCZeroStep)
				notSecurityLab3c.POST("/increment-second-step", h.IncrementLabCSecondStep)
				notSecurityLab3c.POST("/alternative-matrix", h.SendLab3CAlternativeImportanceMatrix)
				notSecurityLab3c.POST("/criteria-matrix", h.SendLab3CCriteriaImportanceMatrix)
				notSecurityLab3c.POST("/current-matrix", h.SendLab3CResultMatrix)
				notSecurityLab3c.POST("/estimation", h.SendLab3CEstimation)
				notSecurityLab3c.POST("/area", h.SendLab3CArea)
				notSecurityLab3c.POST("/line", h.SendLab3CLineParameters)
				notSecurityLab3c.POST("/quadratic", h.SendLab3CQuadraticEquationParameters)
				notSecurityLab3c.POST("/result", h.SendLab3CMiddleOfArea)
				notSecurityLab3c.GET("/info", h.GetCurrentStepLab3C)
			}
		}
	}

	return router
}
