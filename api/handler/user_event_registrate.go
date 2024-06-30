package handler

import (
	"crm_api_gateway/genproto/event_registrate_service"
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
	event_registrate := &event_registrate_service.CreateEventRegistrate{}

	// if !validator.CheckDeadline() {
	// 	handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
	// 	return
	// }

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
	id := c.Param("id")
	event_registrate := &event_registrate_service.EventRegistratePrimaryKey{Id: id}

	resp, err := h.grpcClient.EventRegistrateService().GetByID(c.Request.Context(), event_registrate)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting event_registrate")
		return
	}
	c.JSON(http.StatusOK, resp)
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
	id := c.Param("id")
	event_registrate := &event_registrate_service.EventRegistratePrimaryKey{Id: id}

	resp, err := h.grpcClient.EventRegistrateService().Delete(c.Request.Context(), event_registrate)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting event_registrate")
		return
	}
	c.JSON(http.StatusOK, resp)
}
