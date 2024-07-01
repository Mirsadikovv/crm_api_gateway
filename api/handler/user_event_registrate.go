package handler

import (
	"crm_api_gateway/genproto/event_registrate_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/event_registrate/create [POST]
// @Summary Create event_registrate
// @Description API for creating event_registratees
// @Tags event_registrate
// @Accept  json
// @Produce  json
// @Param		event_registrate body  event_registrate_service.CreateEventRegistrate true "event_registrate"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateEventRegistrate(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" || authInfo.UserRole == "student" {

		event_registrate := &event_registrate_service.CreateEventRegistrate{}

		if err := c.ShouldBindJSON(&event_registrate); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.EventRegistrateService().Create(c.Request.Context(), event_registrate)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating event_registrate")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators, students and managers can change event_registrate")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event_registrate/update [PUT]
// @Summary Update event_registrate
// @Description API for Updating event_registratees
// @Tags event_registrate
// @Accept  json
// @Produce  json
// @Param		event_registrate body  event_registrate_service.UpdateEventRegistrate true "event_registrate"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateEventRegistrate(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" || authInfo.UserRole == "student" {

		event_registrate := &event_registrate_service.UpdateEventRegistrate{}
		if err := c.ShouldBindJSON(&event_registrate); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.EventRegistrateService().Update(c.Request.Context(), event_registrate)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating event_registrate")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators, students and managers can change event_registrate")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event_registrate/get/{id} [GET]
// @Summary Get event_registrate
// @Description API for getting event_registrate
// @Tags event_registrate
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetEventRegistrateById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" || authInfo.UserRole == "student" {

		id := c.Param("id")
		event_registrate := &event_registrate_service.EventRegistratePrimaryKey{Id: id}

		resp, err := h.grpcClient.EventRegistrateService().GetByID(c.Request.Context(), event_registrate)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting event_registrate")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators, students and managers can change event_registrate")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event_registrate/delete/{id} [DELETE]
// @Summary Delete event_registrate
// @Description API for deleting event_registrate
// @Tags event_registrate
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteEventRegistrate(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" || authInfo.UserRole == "student" {

		id := c.Param("id")
		event_registrate := &event_registrate_service.EventRegistratePrimaryKey{Id: id}

		resp, err := h.grpcClient.EventRegistrateService().Delete(c.Request.Context(), event_registrate)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting event_registrate")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators, students and managers can change event_registrate")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event_registrate/getall [GET]
// @Summary Get all eventes
// @Description API for getting all eventes
// @Tags WEB
// @Accept  json
// @Produce  json
// @Param		student_id query string true "student_id"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetRegistratedEvent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" || authInfo.UserRole == "student" {

		event := &event_registrate_service.GetListEventRegistrateRequest{}

		search := c.Query("student_id")

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

		event.Search = search
		event.Offset = int64(page)
		event.Limit = int64(limit)

		resp, err := h.grpcClient.EventRegistrateService().GetStudentEvent(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, student, administrators and managers can change event")
	}
}
