package handler

import (
	"crm_api_gateway/genproto/schedule_service"
	"errors"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change schedule")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change schedule")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change schedule")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		schedule := &schedule_service.SchedulePrimaryKey{Id: id}

		resp, err := h.grpcClient.ScheduleService().GetByID(c.Request.Context(), schedule)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting schedule")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change schedule")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		schedule := &schedule_service.SchedulePrimaryKey{Id: id}

		resp, err := h.grpcClient.ScheduleService().Delete(c.Request.Context(), schedule)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting schedule")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can change schedule")
	}
}

// @Security ApiKeyAuth
// @Router /v1/students/schedule/get/{id} [GET]
// @Summary Get schedule
// @Description API for getting schedule
// @Tags WEB
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetStudentSchedule(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "student" {

		id := c.Param("id")
		schedule := &schedule_service.SchedulePrimaryKey{Id: id}

		resp, err := h.grpcClient.ScheduleService().GetStudentSchedule(c.Request.Context(), schedule)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting schedule of student")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, students can get schedule")
	}
}

// @Security ApiKeyAuth
// @Router /v1/schedule/from_to/get [GET]
// @Summary Get all schedules between date
// @Description API for getting all schedules between date
// @Tags Report
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		from_date query string true "from_date"
// @Param		to_date query string true "to_date"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetWeekSchedule(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		schedule := &schedule_service.GetWeekScheduleRequest{}

		search := c.Query("search")

		from_date := c.Query("from_date")

		to_date := c.Query("to_date")
		schedule.Search = search
		schedule.FromDate = string(from_date)
		schedule.ToDate = string(to_date)

		resp, err := h.grpcClient.ScheduleService().GetForWeek(c.Request.Context(), schedule)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting schedule by date")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get schedule")
	}
}
