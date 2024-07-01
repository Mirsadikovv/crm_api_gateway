package handler

import (
	"crm_api_gateway/genproto/journal_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/journal/getall [GET]
// @Summary Get all journales
// @Description API for getting all journales
// @Tags journal
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllJournal(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		journal := &journal_service.GetListJournalRequest{}

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

		journal.Search = search
		journal.Offset = int64(page)
		journal.Limit = int64(limit)

		resp, err := h.grpcClient.JournalService().GetList(c.Request.Context(), journal)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating journal")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get journal")
	}
}

// @Security ApiKeyAuth
// @Router /v1/journal/create [POST]
// @Summary Create journal
// @Description API for creating journales
// @Tags journal
// @Accept  json
// @Produce  json
// @Param		journal body  journal_service.CreateJournal true "journal"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateJournal(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		journal := &journal_service.CreateJournal{}
		if err := c.ShouldBindJSON(&journal); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.JournalService().Create(c.Request.Context(), journal)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating journal")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get journal")
	}
}

// @Security ApiKeyAuth
// @Router /v1/journal/update [PUT]
// @Summary Update journal
// @Description API for Updating journales
// @Tags journal
// @Accept  json
// @Produce  json
// @Param		journal body  journal_service.UpdateJournal true "journal"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateJournal(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		journal := &journal_service.UpdateJournal{}
		if err := c.ShouldBindJSON(&journal); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.JournalService().Update(c.Request.Context(), journal)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating journal")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get journal")
	}
}

// @Security ApiKeyAuth
// @Router /v1/journal/get/{id} [GET]
// @Summary Get journal
// @Description API for getting journal
// @Tags journal
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetJournalById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		journal := &journal_service.JournalPrimaryKey{Id: id}

		resp, err := h.grpcClient.JournalService().GetByID(c.Request.Context(), journal)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting journal")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get journal")
	}
}

// @Security ApiKeyAuth
// @Router /v1/journal/delete/{id} [DELETE]
// @Summary Delete journal
// @Description API for deleting journal
// @Tags journal
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteJournal(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "administrator" {

		id := c.Param("id")
		journal := &journal_service.JournalPrimaryKey{Id: id}

		resp, err := h.grpcClient.JournalService().Delete(c.Request.Context(), journal)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting journal")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrators and managers can get journal")
	}
}
