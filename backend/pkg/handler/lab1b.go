package handler

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"backend/pkg/handler/errorResponse"
	"backend/pkg/model"
	"backend/pkg/service"
)

func (h *Handler) OpenFirstBLab(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	minutesDuration, err := strconv.Atoi(os.Getenv("FIRST_LAB_B_DURATION_MINUTES"))
	if err != nil {
		err = fmt.Errorf("ошибка получения продолжительности работы")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userDone, err := h.Service.GetLab1BVariance(ctx, userId, service.Lab1BId)
	if err != nil {
		err = fmt.Errorf("ошибка получения варианта")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := h.Service.GetUserInfo(userId, service.Lab1BId)
	if err != nil {
		err = fmt.Errorf("ошибка получения информации о лабораторной работе")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
		"variant": userDone,
	})

	go func() {
		logrus.Println(fmt.Sprintf("START user:%d lab:%d", userId, service.Lab1BId))

		time.Sleep(time.Duration(minutesDuration+1) * time.Minute)

		if h.Service.IsEmptyToken(userId, service.Lab1BId) {
			return
		}

		if userInfo.IsDone {
			return
		}

		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab1BId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab1BId)
			return
		}

		if err := h.Service.SendLabMark(ctx, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR Lab1A send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab1BId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab1BId)
		}

		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab1BId, userMark))
	}()
}

func (h *Handler) OpenFirstBLabForStudent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	user := c.Query("user_id")
	isOpen := c.Query("is_open")
	externalLab := c.Query("lab_id")

	userId, err := strconv.Atoi(user)
	if err != nil {
		err = fmt.Errorf("ошибка получения студента")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	externalLabId, err := strconv.Atoi(externalLab)
	if err != nil {
		err = fmt.Errorf("ошибка получения лабораторной работы")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	isOpenBool, err := strconv.ParseBool(isOpen)
	if err != nil {
		err = fmt.Errorf("ошибка получения изменения проведения лабораторной работы")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if isOpenBool {
		if _, err := h.Service.OpenLab1ForStudent(ctx, userId, service.Lab1BId, externalLabId); err != nil {
			err = fmt.Errorf("ошибка открытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		variance, err := h.Service.GenerateLab1BTask(ctx, userId)
		if err != nil {
			err = fmt.Errorf("ошибка формирования варианта лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := h.Service.UpdateUserVarianceLab1B(ctx, userId, service.Lab1BId, variance); err != nil {
			err = fmt.Errorf("ошибка сохранения варианта лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := h.Service.CloseLabForStudent(ctx, userId, service.Lab1BId); err != nil {
			err = fmt.Errorf("ошибка закрытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

type alternativesInput struct {
	Alternatives []string `json:"alternatives" binding:"required"`
}

func (h *Handler) AddLab1BAlternatives(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes alternativesInput
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 0 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.AddAlternativesLab1B(ctx, userId, userRes.Alternatives); err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type criteriaInput struct {
	Criterias []model.MainCriteria `json:"criterias" binding:"required"`
}

func (h *Handler) AddLab1BCriterias(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes criteriaInput
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 0 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.AddCriteriasLab1B(ctx, userId, userRes.Criterias); err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 1); err != nil {
		logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab1BMainCriteriaStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab1BCommonMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 1 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1BFirstStep(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 2); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) CheckMatrixIsCorrect(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var userRes model.Lab1BMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	for i := range userRes.Matrix {
		for j := range userRes.Matrix[i] {
			if len(userRes.Matrix) != len(userRes.Matrix[j]) {
				c.JSON(http.StatusOK, false)
				return
			}

			if userRes.Matrix[i][j] <= 0 {
				c.JSON(http.StatusOK, false)
				return
			}
			if i == j {
				if userRes.Matrix[i][j] != 1 {
					c.JSON(http.StatusOK, false)
					return
				}
			}

			maxElem := userRes.Matrix[j][i]
			minElem := userRes.Matrix[i][j]
			if userRes.Matrix[j][i] <= userRes.Matrix[i][j] {
				maxElem = userRes.Matrix[i][j]
				minElem = userRes.Matrix[j][i]
			}
			if math.Round(100/maxElem)/100 != math.Round(100*minElem)/100 {
				c.JSON(http.StatusOK, false)
				return
			}
		}
	}

	logrus.Infof("check is matrix correct %d", userId)

	data := h.Service.CheckMatrixIsCorrect(ctx, userRes.Matrix)

	c.JSON(http.StatusOK, data)
}

func (h *Handler) SendLab1BCriteriaStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	stepInt, err := strconv.Atoi(c.Query("step"))
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var userRes model.AnswerLab1BCommonMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 2 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1BSecondStep(ctx, userId, stepInt-1, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) UpdateSecondStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 3); err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab1BMarkAligning(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.Lab1BMarkAligning
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 3 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1BAlignigIeracrhie(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 4); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab1BCountCriteria(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.Lab1BCountCriteria
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 4 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1BCountCriteria(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 5); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) UpdateFivthStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 6); err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab1BQualityCriteria(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	stepInt, err := strconv.Atoi(c.Query("step"))
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var userRes model.AnswerLab1BCommonMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 5 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1QualityCriteria(ctx, userId, stepInt-1, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab1BResult(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab1A5Step
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 6 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1BResult(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1BId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1BId, 7); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1BId, err)
			return
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		userInfo, err := h.Service.GetUserInfo(userId, service.Lab1BId)
		if err != nil {
			logrus.Errorf("ERROR get user info:%d lab:%d", userId, service.Lab1BId)
			return
		}
		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab1BId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab1BId)
			return
		}

		if err := h.Service.SendLabMark(ctx, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR Lab1A send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab1BId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab1BId)
		}

		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab1BId, userMark))
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res.Set,
		"index":      res.Index,
		"max_mark":   maxMark,
	})
}
