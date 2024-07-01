package handler

import (
	"crm_api_gateway/genproto/lesson_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/lesson/getall [GET]
// @Summary Get all lessones
// @Description API for getting all lessones
// @Tags lesson
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllLesson(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		lesson := &lesson_service.GetListLessonRequest{}

		search := c.Query("search")

		page, err := ParsePageQueryParam(c)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
			return
		}
		limit, err := ParseLimitQueryParam(c)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
			return
		}

		lesson.Search = search
		lesson.Offset = int64(page)
		lesson.Limit = int64(limit)

		resp, err := h.grpcClient.LessonService().GetList(c.Request.Context(), lesson)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating lesson")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get lessons")
	}
}

// @Security ApiKeyAuth
// @Router /v1/lesson/create [POST]
// @Summary Create lesson
// @Description API for creating lessones
// @Tags lesson
// @Accept  json
// @Produce  json
// @Param		lesson body  lesson_service.CreateLesson true "lesson"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateLesson(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		lesson := &lesson_service.CreateLesson{}
		if err := c.ShouldBindJSON(&lesson); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.LessonService().Create(c.Request.Context(), lesson)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating lesson")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get lessons")
	}
}

// @Security ApiKeyAuth
// @Router /v1/lesson/update [PUT]
// @Summary Update lesson
// @Description API for Updating lessones
// @Tags lesson
// @Accept  json
// @Produce  json
// @Param		lesson body  lesson_service.UpdateLesson true "lesson"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateLesson(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		lesson := &lesson_service.UpdateLesson{}
		if err := c.ShouldBindJSON(&lesson); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.LessonService().Update(c.Request.Context(), lesson)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating lesson")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get lessons")
	}
}

// @Security ApiKeyAuth
// @Router /v1/lesson/get/{id} [GET]
// @Summary Get lesson
// @Description API for getting lesson
// @Tags lesson
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetLessonById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")
	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		lesson := &lesson_service.LessonPrimaryKey{Id: id}

		resp, err := h.grpcClient.LessonService().GetByID(c.Request.Context(), lesson)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting lesson")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get lessons")
	}
}

// @Security ApiKeyAuth
// @Router /v1/lesson/delete/{id} [DELETE]
// @Summary Delete lesson
// @Description API for deleting lesson
// @Tags lesson
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteLesson(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")
	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		lesson := &lesson_service.LessonPrimaryKey{Id: id}

		resp, err := h.grpcClient.LessonService().Delete(c.Request.Context(), lesson)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting lesson")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get lessons")
	}
}
