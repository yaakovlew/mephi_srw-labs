package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"backend/pkg/handler/errorResponse"
	"backend/pkg/service"
)

const (
	labToken       = "lab-token"
	lecturerHeader = "lecturer-token"
)

func (h *Handler) CheckThirdAHeaderStudentForStart(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3AId, userHeader)
	if err != nil {
		userId, err = h.Service.GetUserId(context.Background(), userHeader)
		if err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка получения user_id").Error())
			return
		}
		if err := h.Service.SaveUserToken(userId, service.Lab3AId, userHeader); err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка сохранения токена").Error())
			return
		}
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3AId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckThirdBHeaderStudentForStart(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("SECOND_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3BId, userHeader)
	if err != nil {
		userId, err = h.Service.GetUserId(context.Background(), userHeader)
		if err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка получения user_id").Error())
			return
		}
		if err := h.Service.SaveUserToken(userId, service.Lab3BId, userHeader); err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка сохранения токена").Error())
			return
		}
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3BId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !user.IsDone {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3BId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckThirdCHeaderStudentForStart(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("THIRD_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3CId, userHeader)
	if err != nil {
		userId, err = h.Service.GetUserId(context.Background(), userHeader)
		if err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка получения user_id").Error())
			return
		}
		if err := h.Service.SaveUserToken(userId, service.Lab3CId, userHeader); err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка сохранения токена").Error())
			return
		}
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !user.IsDone {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckThirdAHeaderLecturer(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	lecHeader := c.GetHeader(lecturerHeader)
	if lecHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty lecturer auth token").Error())
		return
	}

	lecHeaderEnv := os.Getenv("LECTURER_HEADER")
	if lecHeader != lecHeaderEnv {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid lecturer token").Error())
		return
	}
}

func (h *Handler) CheckThirdBHeaderLecturer(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("SECOND_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	lecHeader := c.GetHeader(lecturerHeader)
	if lecHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty lecturer auth token").Error())
		return
	}

	lecHeaderEnv := os.Getenv("LECTURER_HEADER")
	if lecHeader != lecHeaderEnv {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid lecturer token").Error())
		return
	}
}

func (h *Handler) CheckThirdCHeaderLecturer(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("THIRD_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	lecHeader := c.GetHeader(lecturerHeader)
	if lecHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty lecturer auth token").Error())
		return
	}

	lecHeaderEnv := os.Getenv("LECTURER_HEADER")
	if lecHeader != lecHeaderEnv {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid lecturer token").Error())
		return
	}
}

func (h *Handler) CheckThirdAHeaderStudent(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3AId, userHeader)
	if err != nil {

		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("user id not found").Error())
		return
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3AId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckThirdBHeaderStudent(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("SECOND_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3BId, userHeader)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3BId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !user.IsDone {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3BId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckThirdCHeaderStudent(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("THIRD_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab3CId, userHeader)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab3CId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.Service.GetUserInfo(userId, service.Lab3AId)
	if err != nil {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !user.IsDone {
		err = fmt.Errorf("необходимо пройти лабораторную работу 3а")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, _, err := h.Service.GetVarianceLab3(c, userId, service.Lab3CId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckFirstAHeaderStudentForStart(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_A_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab1AId, userHeader)
	if err != nil {
		userId, err = h.Service.GetUserId(context.Background(), userHeader)
		if err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка получения user_id").Error())
			return
		}
		if err := h.Service.SaveUserToken(userId, service.Lab1AId, userHeader); err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка сохранения токена").Error())
			return
		}
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab1AId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.Service.GetLab1AVariance(c, userId, service.Lab1AId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckFirstAHeaderLecturer(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_A_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	lecHeader := c.GetHeader(lecturerHeader)
	if lecHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty lecturer auth token").Error())
		return
	}

	lecHeaderEnv := os.Getenv("LECTURER_HEADER")
	if lecHeader != lecHeaderEnv {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid lecturer token").Error())
		return
	}
}

func (h *Handler) CheckFirstAHeaderStudent(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_A_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab1AId, userHeader)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("user id not found").Error())
		return
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab1AId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := h.Service.GetLab1AVariance(c, userId, service.Lab1AId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckFirstBHeaderLecturer(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_B_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	lecHeader := c.GetHeader(lecturerHeader)
	if lecHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty lecturer auth token").Error())
		return
	}

	lecHeaderEnv := os.Getenv("LECTURER_HEADER")
	if lecHeader != lecHeaderEnv {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid lecturer token").Error())
		return
	}
}

func (h *Handler) CheckFirstBHeaderStudentForStart(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_B_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab1BId, userHeader)
	if err != nil {
		userId, err = h.Service.GetUserId(context.Background(), userHeader)
		if err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка получения user_id").Error())
			return
		}
		if err := h.Service.SaveUserToken(userId, service.Lab1BId, userHeader); err != nil {
			errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("ошибка сохранения токена").Error())
			return
		}
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab1BId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := h.Service.GetLab1BVariance(c, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) CheckFirstBHeaderStudent(c *gin.Context) {
	header := c.GetHeader(labToken)
	if header == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("empty auth token").Error())
		return
	}

	token := os.Getenv("FIRST_B_LAB_TOKEN")
	if header != token {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not valid token").Error())
		return
	}

	userHeader := c.GetHeader(authorizationHeader)
	if userHeader == "" {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.Service.GetUserIdByToken(service.Lab1BId, userHeader)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("user id not found").Error())
		return
	}
	c.Set(userCTX, userId)

	user, err := h.Service.GetUserInfo(userId, service.Lab1BId)
	if err != nil {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.IsDone {
		err = fmt.Errorf("лабораторная работа закрыта для прохождения")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := h.Service.GetLab1BVariance(c, userId, service.Lab1BId); err != nil {
		err = fmt.Errorf("ошибка формирования варианта, обратитесь к администратору")
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
