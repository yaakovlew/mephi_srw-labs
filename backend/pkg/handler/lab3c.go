package handler

import (
	"context"
	"fmt"
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

func (h *Handler) OpenThirdCLab(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	minutesDuration, err := strconv.Atoi(os.Getenv("THIRD_LAB_DURATION_MINUTES"))
	if err != nil {
		err = fmt.Errorf("ошибка получения продолжительности работы")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userInfo, err := h.Service.GetUserInfo(userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("ошибка получения информации о лабораторной работе")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userDone, matrix, err := h.Service.GetVarianceLab3(ctx, userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("ошибка получения варианта")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
		"variant": userDone,
		"matrix":  matrix,
	})

	go func() {
		logrus.Println(fmt.Sprintf("START user:%d lab:%d", userId, service.Lab3CId))

		time.Sleep(time.Duration(minutesDuration) * time.Minute)

		if h.Service.IsEmptyToken(userId, service.Lab3CId) {
			return
		}

		if userInfo.IsDone {
			return
		}

		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab3CId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab3CId)
			return
		}

		if err := h.Service.SendLabMark(c, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR LAB3C send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab3CId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab3CId)
		}
		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab3CId, userMark))
	}()
}

func (h *Handler) OpenThirdCLabForStudent(c *gin.Context) {
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
		isDone, err := h.Service.OpenLab3ForStudent(ctx, userId, service.Lab3CId, externalLabId)
		if err != nil {
			err = fmt.Errorf("ошибка открытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		if isDone {
			go func() {
				routineCtx, routineCancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer routineCancel()
				for {
					select {
					case <-routineCtx.Done():
						return
					default:
						variance, err := h.Service.GenerateLab3Task(ctx, userId)
						if err != nil {
							continue
						}
						setLabA, err := h.Service.ValidateLab3AResult(ctx, variance)
						if err != nil {
							continue
						}
						setLabB, err := h.Service.ValidateLab3BResult(ctx, variance)
						if err != nil {
							continue
						}
						setLabC, err := h.Service.ValidateLab3CResult(ctx, variance)
						if err != nil {
							continue
						}

						var firstMaxA, secondMaxA float64 = 0, 0
						for _, v := range setLabA {
							if firstMaxA <= v {
								secondMaxA = firstMaxA
								firstMaxA = v
							} else if secondMaxA <= v {
								secondMaxA = v
							}
						}

						var firstMaxB, secondMaxB float64 = 0, 0
						for _, v := range setLabB {
							if firstMaxB <= v {
								secondMaxB = firstMaxB
								firstMaxB = v
							} else if secondMaxB <= v {
								secondMaxB = v
							}
						}

						var firstMaxC, secondMaxC float64 = 0, 0
						for _, v := range setLabC {
							if firstMaxC <= v {
								secondMaxC = firstMaxC
								firstMaxC = v
							} else if secondMaxC <= v {
								secondMaxC = v
							}
						}

						if secondMaxA <= 0 || secondMaxB <= 0 || secondMaxC <= 0 {
							continue
						}

						if firstMaxA >= secondMaxA+0.03 && firstMaxB >= secondMaxB+0.03 && firstMaxC >= secondMaxC+0.03 {
							if err := h.Service.UpdateUserVarianceLab3(ctx, userId, service.Lab3CId, variance); err != nil {
								return
							}
							return
						}
					}
				}
			}()
		}
	} else {
		if err := h.Service.CloseLabForStudent(ctx, userId, service.Lab3CId); err != nil {
			err = fmt.Errorf("ошибка открытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetCurrentStepLab3C(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	mark, err := h.Service.GetCurrentMark(userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("ошибка получения текущей оценки")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":    userId,
		"step":       step,
		"percentage": mark,
	})
}

func (h *Handler) SendLab3CAlternativeImportanceMatrix(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3AlternativeImportance
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 0 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CAlternativesImportance(ctx, userId, userRes.Step, userRes.Matrix)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) IncrementLabCZeroStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 0 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.UpdateLabStep(ctx, userId, service.Lab3CId, 1); err != nil {
		logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab3CCriteriaImportanceMatrix(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CriteriaImportance
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 1 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab3CCriteriaImportance(ctx, userId, userRes.Set)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab3CId, 2); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) IncrementLabCSecondStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 2 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.UpdateLabStep(ctx, userId, service.Lab3CId, 3); err != nil {
		logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab3CResultMatrix(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CCurrentMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CCurrentMatrix(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab3CEstimation(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CEstimation
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CEstimation(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab3CArea(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CArea
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CArea(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab3CLineParameters(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CLineParametersRequest
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CLineParameters(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab3CQuadraticEquationParameters(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CQuadraticParametersRequest
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userRes.Step = userRes.Step - 1
	mark, maxMark, res, err := h.Service.CheckLab3CQuadraticParameters(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab3CMiddleOfArea(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	var userRes model.AnswerLab3CMiddleOfArea
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	mark, maxMark, index, res, err := h.Service.CheckLab3CMiddleOfArea(ctx, userId, userRes.Index, userRes.Set)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab3CId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab3CId, 6); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab3CId, err)
			return
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)

		userInfo, err := h.Service.GetUserInfo(userId, service.Lab3CId)
		if err != nil {
			logrus.Errorf("ERROR get user info:%d lab:%d", userId, service.Lab3CId)
			return
		}
		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab3CId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab3CId)
			return
		}

		if err := h.Service.SendLabMark(ctx, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR LAB3A send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab3CId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab3CId)
		}

		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab3CId, userMark))
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"index":      index,
		"max_mark":   maxMark,
	})
}
