package handler

import (
	"crm_api_gateway/genproto/event_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/event/getall [GET]
// @Summary Get all eventes
// @Description API for getting all eventes
// @Tags event
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllEvent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		event := &event_service.GetListEventRequest{}

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

		event.Search = search
		event.Offset = int64(page)
		event.Limit = int64(limit)

		resp, err := h.grpcClient.EventService().GetList(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change event")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event/create [POST]
// @Summary Create event
// @Description API for creating eventes
// @Tags event
// @Accept  json
// @Produce  json
// @Param		event body  event_service.CreateEvent true "event"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateEvent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		event := &event_service.CreateEvent{}
		if err := c.ShouldBindJSON(&event); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if ok, err := validator.IsSunday(event.StartTime); err != nil && ok {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong day of week day must be sunday"), "error while validating event day")
			return
		}

		resp, err := h.grpcClient.EventService().Create(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change event")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event/update [PUT]
// @Summary Update event
// @Description API for Updating eventes
// @Tags event
// @Accept  json
// @Produce  json
// @Param		event body  event_service.UpdateEvent true "event"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateEvent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		event := &event_service.UpdateEvent{}
		if err := c.ShouldBindJSON(&event); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if ok, err := validator.IsSunday(event.StartTime); err != nil && ok {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong day of week day must be sunday"), "error while validating event day")
			return
		}

		resp, err := h.grpcClient.EventService().Update(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change event")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event/get/{id} [GET]
// @Summary Get event
// @Description API for getting event
// @Tags event
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetEventById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		event := &event_service.EventPrimaryKey{Id: id}

		resp, err := h.grpcClient.EventService().GetByID(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change event")
	}
}

// @Security ApiKeyAuth
// @Router /v1/event/delete/{id} [DELETE]
// @Summary Delete event
// @Description API for deleting event
// @Tags event
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteEvent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		event := &event_service.EventPrimaryKey{Id: id}

		resp, err := h.grpcClient.EventService().Delete(c.Request.Context(), event)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting event")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change event")
	}
}
