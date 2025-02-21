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

func (h *Handler) OpenFirstALab(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	minutesDuration, err := strconv.Atoi(os.Getenv("FIRST_LAB_DURATION_MINUTES"))
	if err != nil {
		err = fmt.Errorf("ошибка получения продолжительности работы")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userDone, err := h.Service.GetLab1AVariance(ctx, userId, service.Lab1AId)
	if err != nil {
		err = fmt.Errorf("ошибка получения варианта")
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := h.Service.GetUserInfo(userId, service.Lab1AId)
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
		logrus.Println(fmt.Sprintf("START user:%d lab:%d", userId, service.Lab1AId))

		time.Sleep(time.Duration(minutesDuration+1) * time.Minute)

		if h.Service.IsEmptyToken(userId, service.Lab1AId) {
			return
		}

		if userInfo.IsDone {
			return
		}

		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab1AId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab1AId)
			return
		}

		if err := h.Service.SendLabMark(ctx, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR Lab1A send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab1AId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab1AId)
		}

		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab1AId, userMark))
	}()
}

func (h *Handler) OpenFirstALabForStudent(c *gin.Context) {
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
		if _, err := h.Service.OpenLab1ForStudent(ctx, userId, service.Lab1AId, externalLabId); err != nil {
			err = fmt.Errorf("ошибка открытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		variance, err := h.Service.GenerateLab1ATask(ctx, userId)
		if err != nil {
			err = fmt.Errorf("ошибка формирования варианта лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := h.Service.UpdateUserVarianceLab1A(ctx, userId, service.Lab1AId, variance); err != nil {
			err = fmt.Errorf("ошибка сохранения варианта лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := h.Service.CloseLabForStudent(ctx, userId, service.Lab1AId); err != nil {
			err = fmt.Errorf("ошибка закрытия лабораторной работы")
			errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SendLab1AStep(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, handlerTimeout)
	defer cancel()

	step := c.Param("step")
	stepInt, err := strconv.Atoi(step)
	if err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var userRes model.AnswerLab1ACommonMatrix
	if err := c.BindJSON(&userRes); err != nil {
		err = fmt.Errorf("ошибка отправки ответа")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1AId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != stepInt-1 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1AStep(ctx, userId, stepInt-1, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1AId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1AId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1AId, stepInt); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1AId, err)
			return
		}
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res,
		"max_mark":   maxMark,
	})
}

func (h *Handler) SendLab1AResult(c *gin.Context) {
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

	if step, err := h.Service.GetLabCurrentStep(ctx, userId, service.Lab1AId); err != nil {
		err = fmt.Errorf("необходимо открыть лабораторную работу")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if step != 5 {
		err = fmt.Errorf("необходимо проходить работу пошагово")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mark, maxMark, res, err := h.Service.CheckLab1A5Step(ctx, userId, userRes)
	if err != nil {
		err = fmt.Errorf("ошибка со стороны сервера")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.Service.IncrementPercentageDone(context.Background(), userId, service.Lab1AId, mark); err != nil {
			logrus.Errorf("can't change percentage done user_id:%d labId:%d: %v", userId, service.Lab1AId, err)
			return
		}
		if err := h.Service.UpdateLabStep(ctx, userId, service.Lab1AId, 7); err != nil {
			logrus.Errorf("can't change lab step user_id:%d labId:%d: %v", userId, service.Lab1AId, err)
			return
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		userInfo, err := h.Service.GetUserInfo(userId, service.Lab1AId)
		if err != nil {
			logrus.Errorf("ERROR get user info:%d lab:%d", userId, service.Lab1AId)
			return
		}
		userMark, err := h.Service.GetLabResult(ctx, userId, service.Lab1AId)
		if err != nil {
			logrus.Errorf("ERROR get result user:%d lab:%d", userId, service.Lab1AId)
			return
		}

		if err := h.Service.SendLabMark(ctx, userId, userInfo.ExternalLabId, userMark); err != nil {
			logrus.Errorf("ERROR Lab1A send result user:%d lab:%d", userId, userInfo.ExternalLabId)
		}

		if err := h.Service.ClearToken(userId, service.Lab1AId); err != nil {
			logrus.Errorf("ERROR clear token user:%d lab:%d", userId, service.Lab1AId)
		}

		logrus.Println(fmt.Sprintf("SEND user:%d lab:%d percentage:%d", userId, service.Lab1AId, userMark))
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentage": mark,
		"result":     res.Set,
		"index":      res.Index,
		"max_mark":   maxMark,
	})
}
