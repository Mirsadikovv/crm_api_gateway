package handler

import (
	"crm_api_gateway/genproto/schedule_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/schedule/getall [GET]
// @Summary Get all schedulees
// @Description API for getting all schedulees
// @Tags schedule
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllSchedule(c *gin.Context) {
	schedule := &schedule_service.GetListScheduleRequest{}

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

	schedule.Search = search
	schedule.Offset = int64(page)
	schedule.Limit = int64(limit)

	resp, err := h.grpcClient.ScheduleService().GetList(c.Request.Context(), schedule)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating schedule")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/schedule/create [POST]
// @Summary Create schedule
// @Description API for creating schedulees
// @Tags schedule
// @Accept  json
// @Produce  json
// @Param		schedule body  schedule_service.CreateSchedule true "schedule"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSchedule(c *gin.Context) {
	schedule := &schedule_service.CreateSchedule{}
	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Create(c.Request.Context(), schedule)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating schedule")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/schedule/update [PUT]
// @Summary Update schedule
// @Description API for Updating schedulees
// @Tags schedule
// @Accept  json
// @Produce  json
// @Param		schedule body  schedule_service.UpdateSchedule true "schedule"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSchedule(c *gin.Context) {
	schedule := &schedule_service.UpdateSchedule{}
	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Update(c.Request.Context(), schedule)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating schedule")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/schedule/get/{id} [GET]
// @Summary Get schedule
// @Description API for getting schedule
// @Tags schedule
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetScheduleById(c *gin.Context) {
	id := c.Param("id")
	schedule := &schedule_service.SchedulePrimaryKey{Id: id}

	resp, err := h.grpcClient.ScheduleService().GetByID(c.Request.Context(), schedule)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting schedule")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/schedule/delete/{id} [DELETE]
// @Summary Delete schedule
// @Description API for deleting schedule
// @Tags schedule
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	schedule := &schedule_service.SchedulePrimaryKey{Id: id}

	resp, err := h.grpcClient.ScheduleService().Delete(c.Request.Context(), schedule)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting schedule")
		return
	}
	c.JSON(http.StatusOK, resp)
}
