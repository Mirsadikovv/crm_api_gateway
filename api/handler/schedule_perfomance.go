package handler

import (
	"crm_api_gateway/genproto/perfomance_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/perfomance/getall [GET]
// @Summary Get all perfomancees
// @Description API for getting all perfomancees
// @Tags perfomance
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllPerfomance(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		perfomance := &perfomance_service.GetListPerfomanceRequest{}

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

		perfomance.Search = search
		perfomance.Offset = int64(page)
		perfomance.Limit = int64(limit)

		resp, err := h.grpcClient.PerfomanceService().GetList(c.Request.Context(), perfomance)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating perfomance")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get student_performance")
	}
}

// @Security ApiKeyAuth
// @Router /v1/perfomance/create [POST]
// @Summary Create perfomance
// @Description API for creating perfomancees
// @Tags perfomance
// @Accept  json
// @Produce  json
// @Param		perfomance body  perfomance_service.CreatePerfomance true "perfomance"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreatePerfomance(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "support_teacher" {

		perfomance := &perfomance_service.CreatePerfomance{}
		if err := c.ShouldBindJSON(&perfomance); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.PerfomanceService().Create(c.Request.Context(), perfomance)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating perfomance")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, teachers and support_teachers can change student_performance")
	}
}

// @Security ApiKeyAuth
// @Router /v1/perfomance/update [PUT]
// @Summary Update perfomance
// @Description API for Updating perfomancees
// @Tags perfomance
// @Accept  json
// @Produce  json
// @Param		perfomance body  perfomance_service.UpdatePerfomance true "perfomance"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdatePerfomance(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "support_teacher" {

		perfomance := &perfomance_service.UpdatePerfomance{}
		if err := c.ShouldBindJSON(&perfomance); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.PerfomanceService().Update(c.Request.Context(), perfomance)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating perfomance")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, teachers and support_teachers can change student_performance")
	}
}

// @Security ApiKeyAuth
// @Router /v1/perfomance/get/{id} [GET]
// @Summary Get perfomance
// @Description API for getting perfomance
// @Tags perfomance
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetPerfomanceById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "support_teacher" {

		id := c.Param("id")
		perfomance := &perfomance_service.PerfomancePrimaryKey{Id: id}

		resp, err := h.grpcClient.PerfomanceService().GetByID(c.Request.Context(), perfomance)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting perfomance")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, teachers and support_teachers can change student_performance")
	}
}

// @Security ApiKeyAuth
// @Router /v1/perfomance/delete/{id} [DELETE]
// @Summary Delete perfomance
// @Description API for deleting perfomance
// @Tags perfomance
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeletePerfomance(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		perfomance := &perfomance_service.PerfomancePrimaryKey{Id: id}

		resp, err := h.grpcClient.PerfomanceService().Delete(c.Request.Context(), perfomance)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting perfomance")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can delete student_performance")
	}
}
